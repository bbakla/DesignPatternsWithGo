package protection

import (
	"fmt"
	"strings"
)

type User struct {
	Roles  []string
	Driver Driveable
}

type Driveable interface {
	StartEngine() bool
	Drive() bool
}

type Driver struct {
}

func (driver *Driver) StartEngine() bool {
	return true
}

func (driver *Driver) Drive() bool {
	return true
}

type ProxyDriver struct {
	Driveable
	Roles []string
}

func (proxy *ProxyDriver) StartEngine() bool {
	if contains(proxy.Roles, "driver") {
		return true
	} else {
		fmt.Println("Not enough right to start the engine")
	}

	return false
}

func (proxy *ProxyDriver) Drive() bool {
	if contains(proxy.Roles, "driver") {
		return true
	} else {
		fmt.Println("Not enough right to drive the engine")
	}

	return false
}

func contains(files []string, fileName string) bool {
	for _, name := range files {
		if strings.Contains(name, fileName) {
			return true
		}
	}

	return false
}
