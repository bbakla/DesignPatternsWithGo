package protection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfUserHasAccessToUseTheCar(t *testing.T) {
	t.Run("Dont have role to drive", func(t *testing.T) {
		user := &User{
			Driver: &ProxyDriver{
				Roles: []string{"passenger"},
			},
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
			Driver: proxyDriver,
		}

		started := user.Driver.StartEngine()
		assert.True(t, started, "role is enough")

		user.Driver.Drive()
	})

}
