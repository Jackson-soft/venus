package mysql_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("mysql text test", func() {
	Context("select", func() {
		It("ss", func() {
			query := "select * from demo;"
			rows, err := db.QueryForMapSlice(query)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rows).ShouldNot(BeNil())
		})
	})
})

var _ = Describe("mysql bin test", func() {
	Context("select", func() {
		It("map slice", func() {
			query := "select * from demo where id > ?;"
			rows, err := db.QueryForMapSlice(query, 1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rows).ShouldNot(BeNil())
		})
	})
})
