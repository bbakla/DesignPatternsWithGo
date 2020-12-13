package main

import (
// "fmt"
)

type ComputerAssembling interface {
	ChooseMonitor(string) ComputerAssembling
	ChooseBrand(string) ComputerAssembling
	ChooseProcessor(string) ComputerAssembling
	ChooseStorageSize(int) ComputerAssembling
	Build() Computer
}

type Computer struct {
	Brand       string
	Monitor     string
	Processor   string
	StorageSize int
}

// keeping computer instance in the builder struct violates the immutability. It is not a good code practice here.
type ComputerBuilder struct {
	computer Computer
}

func (c *ComputerBuilder) ChooseMonitor(monitor string) ComputerAssembling {
	c.computer.Monitor = monitor

	return c
}

func (c *ComputerBuilder) ChooseBrand(brand string) ComputerAssembling {
	c.computer.Brand = brand

	return c
}
func (c *ComputerBuilder) ChooseProcessor(processor string) ComputerAssembling {
	c.computer.Processor = processor

	return c
}
func (c *ComputerBuilder) ChooseStorageSize(storageSize int) ComputerAssembling {
	c.computer.StorageSize = storageSize

	return c
}

func (c *ComputerBuilder) Build() Computer {
	return c.computer
}
