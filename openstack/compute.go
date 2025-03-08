package openstack

import (
	"context"
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

// ComputeClient struct holds the OpenStack client
type ComputeClient struct {
	client *gophercloud.ServiceClient
}

// NewComputeClient initializes and returns a ComputeClient
func NewComputeClient(provider *gophercloud.ProviderClient) (*ComputeClient, error) {
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to create compute client: %v", err)
	}
	return &ComputeClient{client: client}, nil
}

// CreateInstance creates a new instance in OpenStack
func (c *ComputeClient) CreateInstance(name, imageID, flavorID, networkID string) (string, error) {
	serverOpts := servers.CreateOpts{
		Name:      name,
		ImageRef:  imageID,
		FlavorRef: flavorID,
		Networks: []servers.Network{
			{UUID: networkID},
		},
	}

	server, err := servers.Create(c.client, serverOpts).Extract()
	if err != nil {
		return "", fmt.Errorf("failed to create instance: %v", err)
	}

	return server.ID, nil
}

// StopInstance stops a running instance
func (c *ComputeClient) StopInstance(instanceID string) error {
	err := servers.Stop(c.client, instanceID).ExtractErr()
	if err != nil {
		return fmt.Errorf("failed to stop instance: %v", err)
	}
	return nil
}

// DeleteInstance deletes an instance
func (c *ComputeClient) DeleteInstance(instanceID string) error {
	err := servers.Delete(c.client, instanceID).ExtractErr()
	if err != nil {
		return fmt.Errorf("failed to delete instance: %v", err)
	}
	return nil
}

// ListUserInstances retrieves a list of instances for a given user
func (c *ComputeClient) ListUserInstances(userID string) ([]servers.Server, error) {
	var instances []servers.Server

	err := servers.List(c.client, servers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		for _, srv := range serverList {
			if srv.UserID == userID {
				instances = append(instances, srv)
			}
		}
		return true, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %v", err)
	}

	return instances, nil
}

// GetInstanceStatus retrieves the current status of an instance
func (c *ComputeClient) GetInstanceStatus(instanceID string) (string, error) {
	server, err := servers.Get(c.client, instanceID).Extract()
	if err != nil {
		return "", fmt.Errorf("failed to get instance status: %v", err)
	}

	return server.Status, nil
}

// MonitorInstanceUsage tracks the instance runtime for credit deduction
func (c *ComputeClient) MonitorInstanceUsage(instanceID string, creditDeduction func(string, float64)) {
	for {
		status, err := c.GetInstanceStatus(instanceID)
		if err != nil {
			fmt.Printf("Error fetching instance status: %v\n", err)
			break
		}

		if status == "ACTIVE" {
			creditDeduction(instanceID, 0.1) // Deduct 0.1 credit per minute
		}

		time.Sleep(60 * time.Second) // Check every minute
	}
}
