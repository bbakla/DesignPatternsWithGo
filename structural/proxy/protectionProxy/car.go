package protectionProxy

import (
	"fmt"
	"strings"
)

type Client struct {
	Driver Driveable
}

type Driveable interface {
	StartEngine() bool
	Drive() bool
}

type Car struct {
}

func (c *Car) StartEngine() bool {
	fmt.Println("Starting the car")

	return true
}

func (c *Car) Drive() bool {
	fmt.Println("Driving the car")
	return true
}

type ProxyAuth struct {
	Driveable
	Roles []string
}

func (proxy *ProxyAuth) StartEngine() bool {
	if contains(proxy.Roles, "driver") {
		return true
	} else {
		fmt.Println("Not enough right to start the engine")

		return false
	}

	return proxy.Driveable.StartEngine()
}

func (proxy *ProxyAuth) Drive() bool {
	if contains(proxy.Roles, "driver") {
		return true
	} else {
		fmt.Println("Not enough right to drive the engine")

		return false
	}

	return proxy.Driveable.Drive()
}

func contains(roles []string, roleName string) bool {
	for _, name := range roles {
		if strings.Contains(name, roleName) {
			return true
		}
	}

	return false
}
