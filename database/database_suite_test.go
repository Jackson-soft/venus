package database_test

import (
	"testing"

	"github.com/Jackson-soft/venus/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var db *database.Database

func TestMysql(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Mysql Suite")
}

var _ = ginkgo.BeforeSuite(func() {
	dsn := "root:ruisi_112@tcp(127.0.0.1:3306)/ruisi"
	var err error
	db, err = database.OpenDB("mysql", dsn)
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
	gomega.Expect(db).ShouldNot(gomega.BeNil())
})

var _ = ginkgo.AfterSuite(func() {
	err := db.Close()
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
})
