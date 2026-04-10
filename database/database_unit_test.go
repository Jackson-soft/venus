package database_test

import (
	"github.com/Jackson-soft/venus/database"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rebind extended", func() {
	It("should not change query without placeholders", func() {
		Expect(database.Rebind("SELECT * FROM users")).Should(Equal("SELECT * FROM users"))
	})

	It("should rebind single placeholder", func() {
		Expect(database.Rebind("SELECT * FROM users WHERE id = ?")).Should(
			Equal("SELECT * FROM users WHERE id = $1"))
	})

	It("should rebind multiple placeholders", func() {
		Expect(database.Rebind("INSERT INTO t (a, b, c) VALUES (?, ?, ?)")).Should(
			Equal("INSERT INTO t (a, b, c) VALUES ($1, $2, $3)"))
	})

	It("should rebind mixed query", func() {
		Expect(database.Rebind("SELECT * FROM t WHERE a = ? AND b > ? ORDER BY c LIMIT ?")).Should(
			Equal("SELECT * FROM t WHERE a = $1 AND b > $2 ORDER BY c LIMIT $3"))
	})
})

var _ = Describe("OpenDB", func() {
	It("should error for invalid driver", func() {
		_, err := database.OpenDB("nonexistent_driver", "dsn_string")
		Expect(err).Should(HaveOccurred())
	})
})
