package openstack

import (
	"fmt"
	"log"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// OpenStackClient holds the authenticated client
type OpenStackClient struct {
	Provider *gophercloud.ProviderClient
}

// NewOpenStackClient initializes and authenticates with OpenStack
func NewOpenStackClient() (*OpenStackClient, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		Username:         os.Getenv("OS_USERNAME"),
		Password:         os.Getenv("OS_PASSWORD"),
		TenantID:         os.Getenv("OS_TENANT_ID"),
		DomainName:       os.Getenv("OS_DOMAIN_NAME"),
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		log.Printf("Failed to authenticate with OpenStack: %v", err)
		return nil, fmt.Errorf("authentication error: %w", err)
	}

	return &OpenStackClient{Provider: provider}, nil
}
