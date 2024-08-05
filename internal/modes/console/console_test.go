package console_test

import (
	"context"
	"github.com/mortedecai/all-text-wrestling/internal/atwerrors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/all-text-wrestling/internal/modes/console"
)

var _ = Describe("Console", func() {
	It("should be possible to create a new console mode with no options", func() {
		ctx, cancel := context.WithCancel(context.Background())
		Expect(cancel).ToNot(BeNil())

		app, err := console.New(ctx, nil)
		Expect(app).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
	})
	It("should return a not yet implemented error when starting", func() {
		ctx, _ := context.WithCancel(context.Background())
		app, err := console.New(ctx, nil)
		Expect(err).ToNot(HaveOccurred())
		err = app.Start()
		Expect(err).To(MatchError(atwerrors.NotYetImplemented))
		Expect(err.Error()).To(HaveSuffix("Start() error"))
	})
	It("should return a not yet implemented error when stopping", func() {
		ctx, _ := context.WithCancel(context.Background())
		app, err := console.New(ctx, nil)
		Expect(err).ToNot(HaveOccurred())
		err = app.Stop()
		Expect(err).To(MatchError(atwerrors.NotYetImplemented))
		Expect(err.Error()).To(HaveSuffix("Stop() error"))
	})
})
