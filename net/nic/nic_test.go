package nic_test

import (
	"net"
	"testing"

	"github.com/marrbor/goutil/net/nic"
	"github.com/stretchr/testify/assert"
)

func TestGetMacAddress(t *testing.T) {
	ip, err := nic.GetIP()
	assert.NoError(t, err)

	card, err := nic.GetInterface(ip)
	assert.NoError(t, err)

	mac, err := nic.GetMacAddress(card)
	assert.NoError(t, err)
	t.Log(mac)
}

func TestValidateMacAddress(t *testing.T) {
	nics, err := net.Interfaces()
	assert.NoError(t, err)
	for _, card := range nics {
		adr := card.HardwareAddr.String()
		if 0 < len(adr) {
			t.Logf(adr)
			assert.True(t, nic.ValidateMacAddress(adr))
		}
	}
}
