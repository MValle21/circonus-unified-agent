package swap

import (
	"fmt"

	"github.com/circonus-labs/circonus-unified-agent/cua"
	"github.com/circonus-labs/circonus-unified-agent/plugins/inputs"
	"github.com/circonus-labs/circonus-unified-agent/plugins/inputs/system"
)

type SwapStats struct {
	ps system.PS
}

func (_ *SwapStats) Description() string {
	return "Read metrics about swap memory usage"
}

func (_ *SwapStats) SampleConfig() string { return "" }

func (s *SwapStats) Gather(acc cua.Accumulator) error {
	swap, err := s.ps.SwapStat()
	if err != nil {
		return fmt.Errorf("error getting swap memory info: %s", err)
	}

	fieldsG := map[string]interface{}{
		"total":        swap.Total,
		"used":         swap.Used,
		"free":         swap.Free,
		"used_percent": swap.UsedPercent,
	}
	fieldsC := map[string]interface{}{
		"in":  swap.Sin,
		"out": swap.Sout,
	}
	acc.AddGauge("swap", fieldsG, nil)
	acc.AddCounter("swap", fieldsC, nil)

	return nil
}

func init() {
	ps := system.NewSystemPS()
	inputs.Add("swap", func() cua.Input {
		return &SwapStats{ps: ps}
	})
}