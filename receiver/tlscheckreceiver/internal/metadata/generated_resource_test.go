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
			rb.SetTlscheckTarget("tlscheck.target-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch tt {
			case "default":
				assert.Equal(t, 1, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 1, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", tt)
			}

			val, ok := res.Attributes().Get("tlscheck.target")
			assert.True(t, ok)
			if ok {
				assert.Equal(t, "tlscheck.target-val", val.Str())
			}
		})
	}
}
