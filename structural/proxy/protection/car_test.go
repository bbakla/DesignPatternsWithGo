package protection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfUserHasAccessToUseTheCar(t *testing.T) {
	t.Run("Dont have role to drive", func(t *testing.T) {
		proxyDriver := &ProxyDriver{
			Roles: []string{"passenger"},
		}

		user := &User{
			Roles:  []string{"passenger"},
			Driver: proxyDriver,
		}

		started := user.Driver.StartEngine()
		assert.False(t, started, "role is not enough")

		user.Driver.Drive()
	})

	t.Run("Dont have role to drive", func(t *testing.T) {
		proxyDriver := &ProxyDriver{
			Roles: []string{"driver"},
		}
		user := &User{
			Roles:  []string{"driver"},
			Driver: proxyDriver,
		}

		started := user.Driver.StartEngine()
		assert.True(t, started, "role is enough")

		user.Driver.Drive()
	})

}
