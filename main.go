package main

import (
	"github.com/dronestock/cos/internal"
	"github.com/dronestock/drone"
)

func main() {
	drone.New(internal.New).Boot()
}
