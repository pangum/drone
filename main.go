package main

import (
	"github.com/dronestock/drone"
	"github.com/pangum/drone/internal"
)

func main() {
	drone.New(internal.New).Boot()
}
