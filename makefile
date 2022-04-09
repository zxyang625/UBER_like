# need to change the paths below when running at linux
# protoc -I pkg/pb --grpc-gateway_out=logtostderr=true:pkg/pb/ pkg/pb/billing.proto
prometheus_path := "E:/Download/prometheus-2.32.1.windows-amd64/prometheus-2.32.1.windows-amd64"
zipkin_path := "E:/Download/Zipkin"
project_path := "C:/Program Files/GoLand/Projects/UBER"
consul:
	consul agent -dev

prometheus:
	$(prometheus_path)/prometheus.exe --config.file=$(prometheus_path)/prometheus.yml

zipkin:
	java -jar $(zipkin_path)/zipkin-server-2.12.9-exec.jar
