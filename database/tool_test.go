package database_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Jackson-soft/venus/database"
)

var _ = Describe("Tool", func() {
	Context("rebind", func() {
		It("mysql to postgres", func() {
			Expect(database.Rebind("select * from table where id = ?")).To(Equal("select * from table where id = $1"))

			Expect(database.Rebind("select * from table where id = ? and name = ?")).To(Equal("select * from table where id = $1 and name = $2"))
		})
	})
})
