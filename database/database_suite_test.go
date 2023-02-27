package database_test

import (
	"testing"

	"github.com/Jackson-soft/venus/database"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var db *database.Database

func TestMysql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mysql Suite")
}

var _ = BeforeSuite(func() {
	dsn := "root:ruisi_112@tcp(127.0.0.1:3306)/ruisi"
	var err error
	db, err = database.OpenDB("mysql", dsn)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(db).ShouldNot(BeNil())
})

var _ = AfterSuite(func() {
	err := db.Close()
	Expect(err).ShouldNot(HaveOccurred())
})
