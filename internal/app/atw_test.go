package app_test

import (
	"context"
	"github.com/mortedecai/all-text-wrestling/internal/atwerrors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/all-text-wrestling/internal/app"
)

var _ = Describe("ATW App", func() {
	Describe("console creation", func() {
		It("should return not yet implemented", func() {
			app, err := app.New(app.ATW_CONSOLE, context.TODO(), nil)
			Expect(err).To(MatchError(atwerrors.NotYetImplemented))
			Expect(app).To(BeNil())
		})
	})
})
