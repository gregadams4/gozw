// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package configurationv2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(BulkGet{})
}

// <no value>
type BulkGet struct {
	ParameterOffset uint16

	NumberOfParameters byte
}

func (cmd BulkGet) CommandClassID() byte {
	return 0x70
}

func (cmd BulkGet) CommandID() byte {
	return byte(CommandBulkGet)
}

func (cmd *BulkGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ParameterOffset = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfParameters = payload[i]
	i++

	return nil
}

func (cmd *BulkGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterOffset)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.NumberOfParameters)

	return
}