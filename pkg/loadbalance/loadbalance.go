package loadbalance

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"math"
	"math/rand"
	Err "pkg/error"
	"sync"
)

type LoadBalancer interface {
	RandomSelect(instances interface{}) (*api.ServiceEntry, error)
	LeastConnSelect(instances interface{}) (*api.ServiceEntry, error)
}

type QueueBalancer interface {
	CalcPriority(length int) int
	SelectQueue(size int) int
}

type LoadBalancerImpl struct {
	Services [][]*Service
	Once     sync.Once
}

type QueueBalancerImpl struct {
	WeightList []struct {
		EffectiveWeight int
		CurrentWeight   int
	}
}

var defaultQueueBalancer = QueueBalancerImpl{
	WeightList: []struct {
		EffectiveWeight int
		CurrentWeight   int
	}{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
	},
}

var defaultLoadBalancerImpl = LoadBalancerImpl{}

func NewQueueBalancer() QueueBalancer {
	return &defaultQueueBalancer
}

func NewLoadBalancer() LoadBalancer {
	return &LoadBalancerImpl{Services: [][]*Service{}}
}

func (d *LoadBalancerImpl) RandomSelect(instances interface{}) (*api.ServiceEntry, error) {
	instanceList, ok := instances.([]*api.ServiceEntry)
	if !ok {
		return nil, Err.New(Err.LoadBalancerSelectFail, "unsupported instances type")
	}
	instance := instanceList[rand.Intn(len(instanceList))]
	return instance, nil
}

func (d *LoadBalancerImpl) LeastConnSelect(instances interface{}) (*api.ServiceEntry, error) {
	instanceList, ok := instances.([]*api.ServiceEntry)
	if !ok {
		return nil, Err.New(Err.LoadBalancerSelectFail, "unsupported instances type")
	}

	serviceIndex := -1
	for k := 0; k < len(d.Services); k++ {
		if d.Services[k][0].Entry.Service.Service == instanceList[0].Service.Service {
			serviceIndex = k
		}
	}

	if serviceIndex == -1 {
		serviceIndex = len(d.Services)
		d.Services = append(d.Services, []*Service{})
		for i := 0; i < len(instanceList); i++ {
			d.Services[serviceIndex] = append(d.Services[serviceIndex], &Service{Entry: instanceList[i], Connections: 0})
		}
	}

	bestIndex := 0
	for i, v := range d.Services[serviceIndex] {
		if i == 0 || v.Connections < d.Services[serviceIndex][bestIndex].Connections {
			bestIndex = i
		}
	}
	d.Services[serviceIndex][bestIndex].Connections += 1

	fmt.Println(d.Services[serviceIndex][bestIndex].Entry.Service.Service, d.Services[serviceIndex][bestIndex].Connections)
	return d.Services[serviceIndex][bestIndex].Entry, nil
	//fmt.Println(instanceList[0].Service.Address, instanceList[0].Service.Port)
	//return instanceList[0], nil
}

func (d *QueueBalancerImpl) CalcPriority(length int) int {
	res := math.Log(float64(length))/math.Log(1.5) + 1
	return int(math.Floor(res))
}

// SelectQueue 权重平滑负载均衡
func (d *QueueBalancerImpl) SelectQueue(size int) int {
	total := 0
	bestIndex, bestWeight := 0, 0
	for i := 1; i <= size; i++ {
		d.WeightList[i].CurrentWeight += d.WeightList[i].EffectiveWeight
		w := d.WeightList[i].CurrentWeight

		total += i

		if bestIndex == 0 || w > bestWeight {
			bestIndex = i
			bestWeight = w
		}
	}

	d.WeightList[bestIndex].CurrentWeight -= total
	return bestIndex
}
