package gateway

import (
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"log"
	"net/http"
	"net/http/httputil"
	"pkg/discover"
	"pkg/loadbalance"
	"strconv"
	"strings"
)

const (
	GatewayURL = "http://localhost:10000"
)

type ServiceInfo struct {

}

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
		log.Println(addr, port)
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", addr, port)
		req.URL.Path = "/" + destPath

		priority := req.Header.Get("Length")
		if priority == "" {
			req.Header.Set("Length", "0")
		} else {
			num, _ := strconv.Atoi(priority)
			req.Header.Set("Length", fmt.Sprintf("%d", num))
		}
	}
	return &httputil.ReverseProxy{
		Director: director,
	}, nil
}
