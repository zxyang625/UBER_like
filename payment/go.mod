module payment

go 1.16

require (
	github.com/go-kit/kit v0.12.0
	github.com/oklog/oklog v0.3.2
	github.com/oklog/run v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin/zipkin-go v0.3.0
	github.com/prometheus/client_golang v1.11.0
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf
	google.golang.org/grpc v1.43.0
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
	pkg v0.0.0
)

replace pkg => ../pkg
