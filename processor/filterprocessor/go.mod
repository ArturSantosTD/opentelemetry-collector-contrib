module github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor

go 1.23.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.123.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.123.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.123.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/component/componenttest v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/confmap v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/confmap/xconfmap v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/consumer v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/consumer/consumertest v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/pdata v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/pipeline v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/processor v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/processor/processorhelper v0.123.1-0.20250411083450-568c3ce6fd43
	go.opentelemetry.io/collector/processor/processortest v0.123.1-0.20250411083450-568c3ce6fd43
	go.opentelemetry.io/otel v1.35.0
	go.opentelemetry.io/otel/metric v1.35.0
	go.opentelemetry.io/otel/sdk/metric v1.35.0
	go.opentelemetry.io/otel/trace v1.35.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/alecthomas/participle/v2 v2.1.4 // indirect
	github.com/antchfx/xmlquery v1.4.4 // indirect
	github.com/antchfx/xpath v1.3.4 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elastic/go-grok v0.3.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/expr-lang/expr v1.17.2 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.123.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twmb/murmur3 v1.1.8 // indirect
	github.com/ua-parser/uap-go v0.0.0-20240611065828-3a4781585db6 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/featuregate v1.29.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/processor/xprocessor v0.123.1-0.20250411083450-568c3ce6fd43 // indirect
	go.opentelemetry.io/collector/semconv v0.123.1-0.20250411083450-568c3ce6fd43 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.10.0 // indirect
	go.opentelemetry.io/otel/log v0.11.0 // indirect
	go.opentelemetry.io/otel/sdk v1.35.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.71.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter => ../../internal/filter

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl => ../../pkg/ottl

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
