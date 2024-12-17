package registry

import (
	"fmt"
	"reflect"
	"sync"
)

var currentRegistry *Registry

type Registry struct {
	mu       sync.RWMutex
	registry map[string][]interface{}
}

func newRegistry() *Registry {
	return &Registry{
		registry: make(map[string][]interface{}),
	}
}

func init() {
	println("Init registry")
	currentRegistry = newRegistry()
}

func GetRegistry() *Registry {
	if currentRegistry == nil {
		println("There's no cached registry. Returning a new one.")
		return newRegistry()
	}

	return currentRegistry
}

func (r *Registry) Register(iface interface{}, impl interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	ifaceType := reflect.TypeOf(iface).String()

	fmt.Printf("Registering %T interface.\n", ifaceType)

	if reflect.TypeOf(impl).Implements(reflect.TypeOf(iface).Elem()) {
		r.registry[ifaceType] = append(r.registry[ifaceType], impl)
		return nil
	}

	return fmt.Errorf("implementation %T does not implement %v", impl, ifaceType)
}

func (r *Registry) GetAll(iface interface{}) ([]interface{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ifaceType := reflect.TypeOf(iface).String()
	impls, exists := r.registry[ifaceType]

	if !exists {
		return nil, fmt.Errorf("no implementations registered for interface %v", ifaceType)
	}

	return impls, nil
}
