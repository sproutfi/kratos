package template_test

import (
	"context"
	"testing"

	"github.com/ory/kratos/courier"
	"github.com/ory/kratos/courier/template/testhelpers"

	"github.com/ory/kratos/courier/template"
	"github.com/ory/kratos/internal"
)

func TestRecoverInvalid(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	t.Run("test=with courier templates directory", func(t *testing.T) {
		_, reg := internal.NewFastRegistryWithMocks(t)
		tpl := template.NewRecoveryInvalid(reg, &template.RecoveryInvalidModel{})

		testhelpers.TestRendered(t, ctx, tpl)
	})

	t.Run("case=test remote resources", func(t *testing.T) {
		testhelpers.TestRemoteTemplates(t, "courier/builtin/templates/recovery/invalid", courier.TypeRecoveryInvalid)
	})
}
