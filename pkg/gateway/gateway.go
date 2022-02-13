package gateway

import (
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"log"
	"net/http"
	"net/http/httputil"
	"pkg/discover"
	"pkg/loadbalance"
	"strings"
)

func NewReverseProxy(consulHost string, consulPort int, logger kitlog.Logger) (*httputil.ReverseProxy, error) {
	client, err := discover.NewDiscoverClient(consulHost, consulPort, true)
	if err != nil {
		logger.Log("NewDiscoverClient", "fail", "err", err)
		return nil, err
	}
	director := func(req *http.Request) {
		reqPath := req.URL.Path
		if reqPath == "" {
			logger.Log("method", "NewReverseProxy", "err", "empty url req path")
			return
		}
		pathArray := strings.Split(reqPath, "/")
		serviceName := pathArray[1]
		log.Println(pathArray[0], pathArray[1], pathArray[2])
		instances, err := client.DiscoverServices(serviceName, "", true)
		if err != nil {
			logger.Log("service name", serviceName, "msg", "query instances failed", "err", err)
			return
		}

		if len(instances) == 0 {
			logger.Log("service name", serviceName, "err", "no such service instance")
			return
		}

		destPath := strings.Join(pathArray[1:], "/")
		addr, port, err := loadbalance.NewLoadBalancer().Select(instances)
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", addr, port)
		req.URL.Path = "/" + destPath
	}
	return &httputil.ReverseProxy{
		Director: director,
	}, nil
}