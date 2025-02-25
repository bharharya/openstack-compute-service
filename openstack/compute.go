package openstack

import (
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func CreateInstance(provider *gophercloud.ProviderClient, name string, flavorID string, imageID string, networkID string) (*servers.Server, error) {
    client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
    if err != nil {
        return nil, err
    }

    createOpts := servers.CreateOpts{
        Name:      name,
        FlavorRef: flavorID,
        ImageRef:  imageID,
        Networks: []servers.Network{
            {UUID: networkID},
        },
    }

    server, err := servers.Create(client, createOpts).Extract()
    if err != nil {
        log.Println("Error creating instance:", err)
        return nil, err
    }

    return server, nil
}
