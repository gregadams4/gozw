// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package hrvcontrol

// <no value>

type HrvControlVentilationRateReport struct {
	VentilationRate byte
}

func ParseHrvControlVentilationRateReport(payload []byte) HrvControlVentilationRateReport {
	val := HrvControlVentilationRateReport{}

	i := 2

	val.VentilationRate = payload[i]
	i++

	return val
}