package mysql_test

import (
	"testing"

	"github.com/Jackson-soft/venus/mysql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var db *mysql.MySQL

func TestMysql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mysql Suite")
}

var _ = BeforeSuite(func() {
	dsn := "root:ruisi_112@tcp(127.0.0.1:3306)/ruisi"
	db = mysql.NewMySQL()

	err := db.Open(dsn)
	Expect(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	err := db.Close()
	Expect(err).ShouldNot(HaveOccurred())
})
