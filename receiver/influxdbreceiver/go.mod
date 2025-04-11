module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver

go 1.23.0

require (
	github.com/influxdata/influxdb-client-go/v2 v2.14.0
	github.com/influxdata/influxdb-observability/common v0.5.12
	github.com/influxdata/influxdb-observability/influx2otel v0.5.12
	github.com/influxdata/influxdb1-client v0.0.0-20220302092344-a9ab5670611c
	github.com/influxdata/line-protocol/v2 v2.2.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.123.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/component/componentstatus v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/component/componenttest v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/config/confighttp v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/confmap v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/consumer v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/consumer/consumererror v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/consumer/consumertest v0.123.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/pdata v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/receiver v1.29.1-0.20250411074447-4fb7c24ebecc
	go.opentelemetry.io/collector/receiver/receiverhelper v0.123.1-0.20250411083450-568c3ce6fd43
	go.opentelemetry.io/collector/receiver/receivertest v0.123.1-0.20250411083450-568c3ce6fd43
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/frankban/quicktest v1.14.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oapi-codegen/runtime v1.0.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.123.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/client v1.29.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/config/configauth v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/config/configcompression v1.29.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/config/configopaque v1.29.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/config/configtls v1.29.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/extension/extensionauth v1.29.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/featuregate v1.29.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/pipeline v0.123.1-0.20250411074447-4fb7c24ebecc // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.123.1-0.20250411083450-568c3ce6fd43 // indirect
	go.opentelemetry.io/collector/semconv v0.123.1-0.20250411083450-568c3ce6fd43 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.10.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.60.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/log v0.11.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.71.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil
