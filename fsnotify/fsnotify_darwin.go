package fsnotify

type Watcher struct {
	kq_ int
}

func New() (*Watcher, error) {
	return &Watcher{}, nil
}

func (w *Watcher) Mark(name string, act Action) error {
	return nil
}

func (w *Watcher) Wait() (Event, error) {
	return Event{}, nil
}

func (w *Watcher) Close() error {
	return nil
}
