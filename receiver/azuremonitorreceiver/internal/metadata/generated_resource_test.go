// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, tt := range []string{"default", "all_set", "none_set"} {
		t.Run(tt, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt)
			rb := NewResourceBuilder(cfg)
			rb.SetAzuremonitorSubscriptionID("azuremonitor.subscription_id-val")
			rb.SetAzuremonitorTenantID("azuremonitor.tenant_id-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch tt {
			case "default":
				assert.Equal(t, 0, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 2, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", tt)
			}

			val, ok := res.Attributes().Get("azuremonitor.subscription_id")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.Equal(t, "azuremonitor.subscription_id-val", val.Str())
			}
			val, ok = res.Attributes().Get("azuremonitor.tenant_id")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.Equal(t, "azuremonitor.tenant_id-val", val.Str())
			}
		})
	}
}
