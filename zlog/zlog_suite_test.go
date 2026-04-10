package zlog

import (
	"sync"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestZLogSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ZLog Suite")
}

var _ = BeforeSuite(func() {
	std.SetBackend(&mockBackend{})
})

// mockBackend implements Backend for testing
type mockBackend struct {
	mu       sync.Mutex
	messages [][]byte
	synced   bool
	closed   bool
}

func (m *mockBackend) Write(buf []byte) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	cp := make([]byte, len(buf))
	copy(cp, buf)
	m.messages = append(m.messages, cp)
	return len(buf), nil
}

func (m *mockBackend) Sync() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.synced = true
	return nil
}

func (m *mockBackend) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.closed = true
	return nil
}

func (m *mockBackend) messageCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.messages)
}

func newTestZLog(level Level) (*ZLog, *mockBackend) {
	z := NewZLog(level)
	mb := &mockBackend{}
	z.SetBackend(mb)
	return z, mb
}
