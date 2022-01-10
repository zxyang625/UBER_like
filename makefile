# need to change the paths below when running at linux
prometheus_path := "E:/Download/prometheus-2.32.1.windows-amd64/prometheus-2.32.1.windows-amd64"
zipkin_path := "E:/Download/Zipkin"
consul:
	consul agent -dev

prometheus:
	$(prometheus_path)/prometheus.exe --config.file=$(prometheus_path)/prometheus.yml

zipkin:
	java -jar $(zipkin_path)/zipkin-server-2.12.9-exec.jar
