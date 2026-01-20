package database_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/Jackson-soft/venus/database"
)

var _ = ginkgo.Describe("Tool", func() {
	ginkgo.Context("rebind", func() {
		ginkgo.It("mysql to postgres", func() {
			gomega.Expect(database.Rebind("select * from table where id = ?")).To(gomega.Equal("select * from table where id = $1"))

			gomega.Expect(database.Rebind("select * from table where id = ? and name = ?")).To(gomega.Equal("select * from table where id = $1 and name = $2"))

			gomega.Expect(database.Rebind("insert into table (a,b,c) values (?,?,?)")).To(gomega.Equal("insert into table (a,b,c) values ($1,$2,$3)"))
		})
	})
})
