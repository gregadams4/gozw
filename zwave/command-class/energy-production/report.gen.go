// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package energyproduction

// <no value>

type EnergyProductionReport struct {
	ParameterNumber byte

	Size byte

	Scale byte

	Precision byte

	Value []byte
}

func ParseEnergyProductionReport(payload []byte) EnergyProductionReport {
	val := EnergyProductionReport{}

	i := 2

	val.ParameterNumber = payload[i]
	i++

	val.Size = (payload[i] & 0x07)

	val.Scale = (payload[i] & 0x18) << 3

	val.Precision = (payload[i] & 0xE0) << 5

	i += 1

	val.Value = payload[i : i+1]
	i += 1

	return val
}