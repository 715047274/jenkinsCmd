package registry

import (
	"errors"
	"reflect"
)

type ServiceRegistry struct {
	services map[string]interface{}
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string]interface{}),
	}
}

func (r *ServiceRegistry) RegisterService(name string, service interface{}) {
	r.services[name] = service
}

func (r *ServiceRegistry) GetService(name string, service interface{}) (interface{}, error) {
	service, exists := r.services[name]
	if !exists {
		return nil, errors.New("service not found")
	}
	return service, nil
}

func (r *ServiceRegistry) GetServiceAs(name string, targetType reflect.Type) (interface{}, error) {
	service, exists := r.services[name]
	if !exists {
		return nil, errors.New("service not found")
	}
	return service, nil
}
