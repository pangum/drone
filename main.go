package main

import (
	"github.com/dronestock/drone"
	"github.com/pangum/drone/internal/plugin"
)

func main() {
	drone.New(plugin.New).Boot()
}
