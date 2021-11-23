package fsnotify

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/sys/unix"
)

type Watcher struct {
	fd_  int
	rwb_ io.ReadWriter
}

// Procfs constants.
const (
	ProcFsFd     = "/proc/self/fd"
	ProcFsFdInfo = "/proc/self/fdinfo"
)

func New() (*Watcher, error) {
	fd, err := unix.FanotifyInit(unix.FAN_CLOEXEC|unix.FAN_CLASS_NOTIF|unix.FAN_NONBLOCK, unix.O_RDWR|unix.O_LARGEFILE)
	if err != nil {
		return nil, err
	}

	file := os.NewFile(uintptr(fd), "")

	return &Watcher{
		fd_:  fd,
		rwb_: bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file)),
	}, nil
}

func (w *Watcher) Mark(fileName string, act Action) error {
	switch act {
	case MarkAdd:
		return unix.FanotifyMark(w.fd_, unix.FAN_MARK_ADD, unix.FAN_CLOSE_WRITE, unix.AT_FDCWD, fileName)
	default:
		return fmt.Errorf("not support act: %v", act)
	}
}

func (w *Watcher) Wait() (Event, error) {
	var (
		metaData unix.FanotifyEventMetadata

		err   error
		event Event
	)

	if err = binary.Read(w.rwb_, binary.LittleEndian, &metaData); err != nil {
		return event, fmt.Errorf("fanotify read event error: %w", err)
	}

	if metaData.Vers != unix.FANOTIFY_METADATA_VERSION {
		if err = unix.Close(int(metaData.Fd)); err != nil {
			return event, fmt.Errorf("fanotify failed to close event fd: %w", err)
		}

		return event, fmt.Errorf("fanotify wrong metadata version")
	}

	event.Name_, err = getPath(metaData.Fd)
	if err != nil {
		return event, err
	}
	event.Op_ = newOp(metaData.Mask)
	return event, nil
}

func (w *Watcher) ResponseAllow(ev *unix.FanotifyEventMetadata) error {
	if err := binary.Write(
		w.rwb_,
		binary.LittleEndian,
		&unix.FanotifyResponse{
			Fd:       ev.Fd,
			Response: unix.FAN_ALLOW,
		},
	); err != nil {
		return fmt.Errorf("fanotify: response error, %w", err)
	}

	return nil
}

// ResponseDeny sends a deny message back to fanotify, used for permission checks.
func (w *Watcher) ResponseDeny(ev *unix.FanotifyEventMetadata) error {
	if err := binary.Write(
		w.rwb_,
		binary.LittleEndian,
		&unix.FanotifyResponse{
			Fd:       ev.Fd,
			Response: unix.FAN_DENY,
		},
	); err != nil {
		return fmt.Errorf("fanotify: response error, %w", err)
	}

	return nil
}

func (w *Watcher) Close() error {
	return unix.Close(w.fd_)
}

func newOp(mask uint64) Op {
	if mask&unix.FAN_CLOSE_WRITE == unix.FAN_CLOSE_WRITE {
		return Write
	}
	return Create
}

// getPath returns path to file for FD inside event metadata.
func getPath(fd int32) (string, error) {
	path, err := os.Readlink(
		filepath.Join(
			ProcFsFd,
			strconv.FormatUint(
				uint64(fd),
				10,
			),
		),
	)
	if err != nil {
		return "", fmt.Errorf("fanotify get path error: %w", err)
	}

	return path, nil
}
