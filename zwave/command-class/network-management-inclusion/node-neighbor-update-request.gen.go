// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementinclusion

// <no value>

type NodeNeighborUpdateRequest struct {
	SeqNo byte

	NodeId byte
}

func ParseNodeNeighborUpdateRequest(payload []byte) NodeNeighborUpdateRequest {
	val := NodeNeighborUpdateRequest{}

	i := 2

	val.SeqNo = payload[i]
	i++

	val.NodeId = payload[i]
	i++

	return val
}