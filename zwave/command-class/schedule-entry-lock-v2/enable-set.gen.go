// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv2

// <no value>

type ScheduleEntryLockEnableSet struct {
	UserIdentifier byte

	Enabled byte
}

func ParseScheduleEntryLockEnableSet(payload []byte) ScheduleEntryLockEnableSet {
	val := ScheduleEntryLockEnableSet{}

	i := 2

	val.UserIdentifier = payload[i]
	i++

	val.Enabled = payload[i]
	i++

	return val
}