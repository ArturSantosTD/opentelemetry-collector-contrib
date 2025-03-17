module github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension

go 1.23.0

require (
	bitbucket.org/atlassian/go-asap/v2 v2.9.0
	github.com/SermoDigital/jose v0.9.2-0.20180104203859-803625baeddc
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v1.27.1-0.20250316222041-da4512d625a4
	go.opentelemetry.io/collector/component/componenttest v0.121.1-0.20250313100724-0885401136ff
	go.opentelemetry.io/collector/config/configopaque v1.27.1-0.20250316222041-da4512d625a4
	go.opentelemetry.io/collector/confmap v1.27.1-0.20250316222041-da4512d625a4
	go.opentelemetry.io/collector/confmap/xconfmap v0.121.1-0.20250313100724-0885401136ff
	go.opentelemetry.io/collector/extension v1.27.1-0.20250316222041-da4512d625a4
	go.opentelemetry.io/collector/extension/extensionauth v0.121.1-0.20250313100724-0885401136ff
	go.opentelemetry.io/collector/extension/extensiontest v0.121.1-0.20250313100724-0885401136ff
	go.uber.org/multierr v1.11.0
	google.golang.org/grpc v1.71.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/vincent-petithory/dataurl v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.27.1-0.20250316222041-da4512d625a4 // indirect
	go.opentelemetry.io/collector/pdata v1.27.1-0.20250316222041-da4512d625a4 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
