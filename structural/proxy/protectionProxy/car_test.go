package protectionProxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfUserHasAccessToUseTheCar(t *testing.T) {
	t.Run("Dont have role to drive", func(t *testing.T) {
		user := &Client{
			Driver: &ProxyAuth{
				Roles:     []string{"passenger"},
				Driveable: &Car{},
			},
		}

		started := user.Driver.StartEngine()
		assert.False(t, started, "role is not enough")

		user.Driver.Drive()
	})

	t.Run("Dont have role to drive", func(t *testing.T) {
		proxyDriver := &ProxyAuth{
			Roles:     []string{"driver"},
			Driveable: &Car{},
		}
		user := &Client{
			Driver: proxyDriver,
		}

		started := user.Driver.StartEngine()
		assert.True(t, started, "role is enough")

		user.Driver.Drive()
	})

}
