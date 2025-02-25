package main

import (
    "github.com/bharharya/openstack-compute-service/routes"
)

func main() {
    r := routes.SetupRouter()
    r.Run(":8080")
}
