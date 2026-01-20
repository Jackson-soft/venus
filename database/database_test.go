package database_test

import (
	"context"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("mysql text test", func() {
	ginkgo.Context("select", func() {
		ginkgo.It("ss", func() {
			query := "select * from demo;"
			rows, err := db.QueryMapSliceContext(context.Background(), query)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			gomega.Expect(rows).ShouldNot(gomega.BeNil())
		})
	})
})

var _ = ginkgo.Describe("mysql bin test", func() {
	ginkgo.Context("select", func() {
		ginkgo.It("map slice", func() {
			query := "select * from demo where id > ?;"
			rows, err := db.QueryMapSliceContext(context.Background(), query, 1)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			gomega.Expect(rows).ShouldNot(gomega.BeNil())
		})
	})
})
