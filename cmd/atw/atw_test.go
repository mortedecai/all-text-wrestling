package main

import (
	"github.com/mortedecai/all-text-wrestling/internal/modes/console"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Atw", func() {
	Describe("App", func() {
		It("should currently return nil for the app", func() {
			app, cancel := prepare()
			Expect(app).To(BeNil())
			Expect(cancel).ToNot(BeNil())
		})
		XIt("should be possible to retrieve the app and cancel function", func() {
			app, cancel := prepare()
			Expect(app).ToNot(BeNil())
			Expect(cancel).ToNot(BeNil())

			consoleApp, ok := app.(*console.ConsoleATW)
			Expect(ok).To(BeTrue())
			Expect(consoleApp).ToNot(BeNil())

			outputFile, err := os.CreateTemp("./", "writeTest.output")
			defer os.Remove(outputFile.Name())
			Expect(err).ToNot(HaveOccurred())
			Expect(outputFile).ToNot(BeNil())
			consoleApp.StringWriter = outputFile
		})
		/*
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
		*/
	})
})
