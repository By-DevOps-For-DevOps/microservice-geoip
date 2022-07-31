package handlers

import (
	"github.com/oschwald/geoip2-golang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestCountry(t *testing.T) {
	reader, err := geoip2.Open("../db/GeoLite2-Country.mmdb")
	require.NoError(t, err)

	defer reader.Close()
	require.NoError(t, err)

	record, err := reader.Country(net.ParseIP("81.2.69.160"))
	require.NoError(t, err)

	assert.False(t, record.Country.IsInEuropeanUnion)
	assert.False(t, record.RegisteredCountry.IsInEuropeanUnion)
	assert.False(t, record.RepresentedCountry.IsInEuropeanUnion)
}
