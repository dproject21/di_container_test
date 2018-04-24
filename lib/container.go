package lib

import (
	"github.com/dproject21/di_container_test/sampleinterface"
	"github.com/fgrosse/goldi"
)

var container *goldi.Container

func CreateContainer() {
	registry := goldi.NewTypeRegistry()

	RegisterTypes(registry)

	// create a new container when your application loads
	config := map[string]interface{}{}
	container = goldi.NewContainer(registry, config)
}

func GetPrinter() sampleinterface.SamplePrinter {
	p := container.MustGet("printer").(sampleinterface.SamplePrinter)
	return p
}
