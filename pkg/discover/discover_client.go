package discover

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
)

type DiscoveryClient interface {
	Register(serviceName, healthCheckUrl, instanceHost, instancePort string, meta map[string]string, logger kitlog.Logger) (string, bool)
	DeRegister(instanceId string, logger kitlog.Logger) bool
	DiscoverServices(serviceName string, tag string, passingOnly bool) ([]*api.ServiceEntry, error)
}
