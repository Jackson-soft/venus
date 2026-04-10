package zlog

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InciseFileBackend", func() {
	Context("NewInciseFile", func() {
		It("should create with custom settings", func() {
			dir := GinkgoT().TempDir()
			b, err := NewInciseFile(dir, "test.log", "myapp", 10)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b).ShouldNot(BeNil())
			defer b.Close()

			Expect(b.maxFileSize).Should(Equal(int64(10 * 1024 * 1024)))
			Expect(b.namePrefix).Should(Equal("myapp"))
		})

		It("should use defaults for empty prefix and zero maxSize", func() {
			dir := GinkgoT().TempDir()
			b, err := NewInciseFile(dir, "", "", 0)
			Expect(err).ShouldNot(HaveOccurred())
			defer b.Close()

			Expect(b.namePrefix).Should(Equal(defaultPrefix))
			Expect(b.maxFileSize).Should(Equal(defaultMaxSize))
		})

		It("should error for empty file path", func() {
			_, err := NewInciseFile("", "", "", 0)
			Expect(err).Should(HaveOccurred())
		})

		It("should create nested directories", func() {
			dir := filepath.Join(GinkgoT().TempDir(), "sub", "dir")
			b, err := NewInciseFile(dir, "", "test", 0)
			Expect(err).ShouldNot(HaveOccurred())
			defer b.Close()

			info, err := os.Stat(dir)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(info.IsDir()).Should(BeTrue())
		})
	})

	Context("Write and Sync", func() {
		It("should write data and sync", func() {
			dir := GinkgoT().TempDir()
			b, err := NewInciseFile(dir, "", "test", 0)
			Expect(err).ShouldNot(HaveOccurred())
			defer b.Close()

			msg := []byte("hello world\n")
			n, err := b.Write(msg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(n).Should(Equal(len(msg)))
			Expect(b.Sync()).Should(Succeed())
		})
	})

	Context("Symlink", func() {
		It("should create symlink when configured", func() {
			dir := GinkgoT().TempDir()
			b, err := NewInciseFile(dir, "latest.log", "test", 0)
			Expect(err).ShouldNot(HaveOccurred())
			defer b.Close()

			linkPath := filepath.Join(dir, "latest.log")
			_, err = os.Lstat(linkPath)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("File rotation", func() {
		It("should create log files in the directory", func() {
			dir := GinkgoT().TempDir()
			b, err := NewInciseFile(dir, "", "rotate", 0)
			Expect(err).ShouldNot(HaveOccurred())
			defer b.Close()

			b.Write([]byte("test data\n"))
			b.Sync()

			entries, err := os.ReadDir(dir)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(entries).ShouldNot(BeEmpty())
		})
	})
})
