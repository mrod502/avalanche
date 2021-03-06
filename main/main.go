package main

import (
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/mrod502/avalanche/platform"
)

func main() {
	var db database.Database
	vm, err := platform.NewVM(snow.DefaultContextTest())
}
