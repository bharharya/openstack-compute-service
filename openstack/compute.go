package openstack

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

// ComputeService interacts with OpenStack compute API
type ComputeService struct {
	Client *OpenStackClient
}

// NewComputeService initializes the compute service
func NewComputeService(client *OpenStackClient) *ComputeService {
	return &ComputeService{Client: client}
}

// CreateInstance launches a new virtual machine
func (cs *ComputeService) CreateInstance(name, imageID, flavorID, networkID string) (*servers.Server, error) {
    endpointOpts := gophercloud.EndpointOpts{
		Region: "your-region", // Replace with your actual OpenStack region
	}

	client, err := openstack.NewComputeV2(cs.Client.Provider, endpointOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to create compute client: %w", err)
	}

	serverCreateOpts := servers.CreateOpts{
		Name:      name,
		ImageRef:  imageID,
		FlavorRef: flavorID,
		Networks:  []servers.Network{{UUID: networkID}},
	}

	server, err := servers.Create(client, serverCreateOpts).Extract()
	if err != nil {
		log.Printf("Instance creation failed: %v", err)
		return nil, fmt.Errorf("instance creation error: %w", err)
	}

	return server, nil
}

// DeleteInstance deletes a virtual machine
func (cs *ComputeService) DeleteInstance(instanceID string) error {
	// Define endpoint options
	endpointOpts := gophercloud.EndpointOpts{
		Region: "your-region", // Replace with your actual OpenStack region
	}

	// Create Compute V2 client
	client, err := openstack.NewComputeV2(cs.Client.Provider, endpointOpts)
	if err != nil {
		return fmt.Errorf("failed to create compute client: %w", err)
	}

	// Attempt to delete the instance
	err = servers.Delete(client, instanceID).ExtractErr()
	if err != nil {
		log.Printf("Failed to delete instance %s: %v", instanceID, err)
		return fmt.Errorf("instance deletion error: %w", err)
	}

	return nil
}

