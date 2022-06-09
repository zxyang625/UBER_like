package discover

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
	Err "pkg/error"
	"strconv"
)

type DiscoverClientImpl struct {
	Host    string
	Port    int
	UseGRPC bool
	client  consul.Client
	config  *api.Config
}

func NewDiscoverClient(consulHost string, consulPort int, useGRPC bool) (DiscoveryClient, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, Err.Errorf(Err.DiscoverNewClientFail, "NewDiscoverClient failed, err: %v", err)
	}
	client := consul.NewClient(apiClient)
	return &DiscoverClientImpl{
		Host:    consulHost,
		Port:    consulPort,
		UseGRPC: useGRPC,
		client:  client,
		config:  consulConfig,
	}, nil
}

func (d *DiscoverClientImpl) Register(serviceName, healthCheckUrl, instanceHost, instancePort string, meta map[string]string, logger kitlog.Logger) (string, bool) {
	instanceID := serviceName + "-" + uuid.New().String()
	port, _ := strconv.Atoi(instancePort)
	serviceRegistration := &api.AgentServiceRegistration{
		ID:      instanceID,
		Name:    serviceName,
		Address: instanceHost,
		Port:    port + 1,
		Meta:    meta,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "30s",
			Interval:                       "10s",
			Timeout:                        "1s",
			Notes:                          "Consul check service health status.",
		},
	}
	//if instanceHost == "" {
	//	serviceRegistration.Address = "127.0.0.1"
	//}
	//if port == 0 {
	//	if d.UseGRPC {
	//		port = 8082
	//		serviceRegistration.Port = 8083
	//	} else {
	//		port = 8081
	//	}
	//}
	if healthCheckUrl == "" {
		if d.UseGRPC {
			serviceRegistration.Check.GRPC = instanceHost + ":" + strconv.Itoa(port) + "/" + serviceName
		} else {
			serviceRegistration.Check.HTTP = "http://" + instanceHost + ":" + strconv.Itoa(port) + "/health-check"
		}
	}
	err := d.client.Register(serviceRegistration)
	if err != nil {
		logger.Log(Err.New(Err.DiscoverRegisterFail, err))
		return "", false
	}
	log.Println("Register OriginService Success")
	return instanceID, true
}

func (d *DiscoverClientImpl) DeRegister(instanceId string, logger kitlog.Logger) bool {
	serviceRegistration := &api.AgentServiceRegistration{
		ID: instanceId,
	}
	err := d.client.Deregister(serviceRegistration)
	if err != nil {
		logger.Log(Err.New(Err.DiscoverDeregisterFail, err))
		return false
	}
	log.Println("Deregister OriginService Success")
	return true
}

func (d *DiscoverClientImpl) DiscoverServices(serviceName string, tag string, passingOnly bool) ([]*api.ServiceEntry, error) {
	entries, _, err := d.client.Service(serviceName, tag, passingOnly, nil)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
