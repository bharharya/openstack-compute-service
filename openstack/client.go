package openstack

import (
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/openstack"
    "log"
)

func GetOpenStackClient(authURL, username, password string) (*gophercloud.ProviderClient, error) {
    opts := gophercloud.AuthOptions{
        IdentityEndpoint: authURL,
        Username: username,
        Password: password,
    }

    provider, err := openstack.AuthenticatedClient(opts)
    if err != nil {
        log.Println("Failed to authenticate OpenStack:", err)
        return nil, err
    }

    return provider, nil
}

