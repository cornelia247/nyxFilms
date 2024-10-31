package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Registry defines a service registry.
type Registry interface {
	// Register creates a service instance record in the registry.
	Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error

	// Deregister removes a service instance record from the registry.
	Deregister(ctx context.Context, instanceID string, serviceName string) error

	// ServiceAddresses returns the list of addresses or active instances of the given service.
	ServiceAddresses(ctx context.Context, serviceName string) ([]string, error)

	// ReportHealthyState is a push mechanism for reporting healthy state to the registry.
	ReportHealthyState(instanceID string, serviceName string)error

}

// ErrNotFound is returned when no service addresses are found.
var ErrNotFound = errors.New("no service addresses found")

// GenerateInstanceID generates a pseudi-random service instance identifier, using service name suffixed by dash and a random number
func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}