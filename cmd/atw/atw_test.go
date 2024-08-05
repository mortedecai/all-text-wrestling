package main

import (
	"io"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Atw", func() {
	Describe("App", func() {
		It("should be able to create with the default Stdout", func() {
			app := New(os.Stdout)
			Expect(app).ToNot(BeNil())
			Expect(app.out).To(Equal(os.Stdout))
		})
		It("should be possible to write to the file", func() {
			outputFile, err := os.CreateTemp("./", "writeTest.output")
			defer os.Remove(outputFile.Name())
			Expect(err).ToNot(HaveOccurred())
			Expect(outputFile).ToNot(BeNil())

			app := New(outputFile)
			app.Write("Hello\n")
			outputFile.Seek(0, 0)

			data, err := io.ReadAll(outputFile)
			Expect(err).ToNot(HaveOccurred())
			Expect(data).ToNot(BeEmpty())

			msg := string(data)
			Expect(msg).To(Equal("Hello\n"))
		})
	})
})
