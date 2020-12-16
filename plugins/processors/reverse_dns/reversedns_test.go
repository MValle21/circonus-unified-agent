package reverse_dns

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/circonus-labs/circonus-unified-agent/config"
	"github.com/circonus-labs/circonus-unified-agent/metric"
	"github.com/circonus-labs/circonus-unified-agent/testutil"
)

func TestSimpleReverseLookup(t *testing.T) {
	now := time.Now()
	m, _ := metric.New("name", map[string]string{
		"dest_ip": "8.8.8.8",
	}, map[string]interface{}{
		"source_ip": "127.0.0.1",
	}, now)

	dns := newReverseDNS()
	dns.Log = &testutil.Logger{}
	dns.Lookups = []lookupEntry{
		{
			Field: "source_ip",
			Dest:  "source_name",
		},
		{
			Tag:  "dest_ip",
			Dest: "dest_name",
		},
	}
	acc := &testutil.Accumulator{}
	_ = dns.Start(acc)
	_ = dns.Add(m, acc)
	_ = dns.Stop()
	// should be processed now.

	require.Len(t, acc.GetCUAMetrics(), 1)
	processedMetric := acc.GetCUAMetrics()[0]
	f, ok := processedMetric.GetField("source_name")
	require.True(t, ok)
	if runtime.GOOS != "windows" {
		// lookupAddr on Windows works differently than on Linux so `source_name` won't be "localhost" on every environment
		require.EqualValues(t, "localhost", f)
	}

	tag, ok := processedMetric.GetTag("dest_name")
	require.True(t, ok)
	require.EqualValues(t, "dns.google.", tag)
}

func TestLoadingConfig(t *testing.T) {
	c := config.NewConfig()
	err := c.LoadConfigData([]byte("[[processors.reverse_dns]]\n" + sampleConfig))
	require.NoError(t, err)

	require.Len(t, c.Processors, 1)
}
