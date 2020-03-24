package system

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRawSocketAvailable(t *testing.T) {
	// Can't test this one better.
	assert.NotPanics(t, func() { IsRawSocketAvailable() })
}

func TestListLocalhostIpv4(t *testing.T) {
	ips, err := ListLocalhostIpv4()
	assert.NoError(t, err)
	assert.Contains(t, ips, "127.0.0.1")
}

func TestListLocalhostIpv4_WithError(t *testing.T) {
	mockInterfaces := func() ([]net.Interface, error) {
		return nil, errors.New("boom")
	}
	_, err := listLocalhostIpv4(mockInterfaces)
	assert.Error(t, err)
}
