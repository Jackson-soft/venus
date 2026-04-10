package httpkit_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Jackson-soft/venus/httpkit"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("httpkit", func() {
	Context("WebDo", func() {
		It("should succeed with 200 response", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"status":"ok"}`))
			}))
			defer ts.Close()

			body, err := httpkit.WebDo(&httpkit.WebBase{
				Url:    ts.URL,
				Method: http.MethodGet,
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).Should(Equal(`{"status":"ok"}`))
		})

		It("should return error for non-200 status", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
			}))
			defer ts.Close()

			_, err := httpkit.WebDo(&httpkit.WebBase{
				Url:    ts.URL,
				Method: http.MethodGet,
			})
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("404"))
		})

		It("should send custom headers", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get(httpkit.HeaderType) != httpkit.HeaderJson {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if r.Header.Get(httpkit.HeaderAuth) != "Bearer token123" {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("ok"))
			}))
			defer ts.Close()

			body, err := httpkit.WebDo(&httpkit.WebBase{
				Url:    ts.URL,
				Method: http.MethodGet,
				Header: map[string]string{
					httpkit.HeaderType: httpkit.HeaderJson,
					httpkit.HeaderAuth: "Bearer token123",
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).Should(Equal("ok"))
		})

		It("should send request body", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				data, _ := io.ReadAll(r.Body)
				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}))
			defer ts.Close()

			body, err := httpkit.WebDo(&httpkit.WebBase{
				Url:    ts.URL,
				Method: http.MethodPost,
				Header: map[string]string{httpkit.HeaderType: httpkit.HeaderJson},
				Body:   strings.NewReader(`{"key":"value"}`),
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).Should(Equal(`{"key":"value"}`))
		})

		It("should error for invalid URL", func() {
			_, err := httpkit.WebDo(&httpkit.WebBase{
				Url:    "http://127.0.0.1:0/bad",
				Method: http.MethodGet,
			})
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("HttpDo", func() {
		It("should unmarshal JSON response", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"name":"alice","age":30}`))
			}))
			defer ts.Close()

			var result struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}
			err := httpkit.HttpDo(&httpkit.WebBase{
				Url:    ts.URL,
				Method: http.MethodGet,
			}, &result)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Name).Should(Equal("alice"))
			Expect(result.Age).Should(Equal(30))
		})

		It("should error for invalid JSON", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("not json"))
			}))
			defer ts.Close()

			var result map[string]any
			err := httpkit.HttpDo(&httpkit.WebBase{
				Url:    ts.URL,
				Method: http.MethodGet,
			}, &result)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Header constants", func() {
		It("should have correct values", func() {
			Expect(httpkit.HeaderAuth).Should(Equal("Authorization"))
			Expect(httpkit.HeaderType).Should(Equal("Content-Type"))
			Expect(httpkit.HeaderJson).Should(Equal("application/json"))
			Expect(httpkit.HeaderUrl).Should(Equal("application/x-www-form-urlencoded"))
		})
	})
})
