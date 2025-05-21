package fsnotify

/*
   配置文件监视器，用于配置热加载

   目前的实现是：
   linux 基于 https://man7.org/linux/man-pages/man7/fanotify.7.html
   macOs 暂未实现
*/

// Event represents a single file system notification.
type (
	Event struct {
		Name_ string // Relative path to the file or directory.
		Op_   Op     // File operation that triggered the event.
	}

	// Op describes a set of file operations.
	Op uint

	// action
	Action uint
)

// These are the generalized file operations that can trigger a notification.
const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

const (
	MarkAdd Action = 1 << iota
	MarkRemove
	MarkFlush
)
