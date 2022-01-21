package loadbalance

import (
	"github.com/hashicorp/consul/api"
	"math/rand"
	Err "pkg/error"
)

type LoadBalancer interface {
	Select(instances interface{}) (addr string, port int, err error)
}

type LoadBalancerImpl struct {

}

var defaultLoadBalancer = LoadBalancerImpl{}

func NewLoadBalancer() LoadBalancer {
	return &defaultLoadBalancer
}

func (d *LoadBalancerImpl) Select(instances interface{}) (addr string, port int, err error) {
	instanceList, ok := instances.([]*api.ServiceEntry)
	if !ok {
		return "", 0, Err.New(Err.LoadBalancerSelectFail, "unsupported instances type")
	}
	instance := instanceList[rand.Intn(len(instanceList))]
	return instance.Service.Address, instance.Service.Port, nil
}




