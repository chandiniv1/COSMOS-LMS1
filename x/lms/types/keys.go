package types

const (
	ModuleName = "lms"

	StoreKey = ModuleName

	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	adminKey   = []byte{0x01}
	StudentKey = []byte{0x02}
	//LeavesKey   = []byte{0x03}
	CounterKey        = []byte{0x03}
	SequenceKey       = []byte{0x04}
	AcceptedLeavesKey = []byte{0x05}
	AppliedLeavesKey  = []byte{0x06}
)

func AdminStoreKey(address string) []byte {
	key := make([]byte, len(adminKey)+len(address))
	copy(key, adminKey)
	copy(key[len(adminKey):], []byte(address))
	return key
}

func StudentStoreKey(studentID string) []byte {
	key := make([]byte, len(StudentKey)+len(studentID))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], []byte(studentID))
	return key
}

// func LeavesStoreKey(leaveID string) []byte {
// 	key := make([]byte, len(LeavesKey)+len(leaveID))
// 	copy(key, LeavesKey)
// 	copy(key[len(LeavesKey):], leaveID)
// 	return key
// }

func AcceptedLeavesStoreKey(admin string, leaveID string) []byte {
	key := make([]byte, len(AcceptedLeavesKey)+len(admin)+len(SequenceKey)+len(leaveID))
	copy(key, AcceptedLeavesKey)
	copy(key[len(AcceptedLeavesKey):], []byte(admin))
	copy(key[len(AcceptedLeavesKey)+len(admin):], SequenceKey)
	// copy(key, SequenceKey)
	copy(key[len(AcceptedLeavesKey)+len(admin)+len(SequenceKey):], []byte(leaveID))
	// copy(key[len(SequenceKey):], leaveID)
	return key
}

func AppliedLeavesStoreKey(leaveID string, leavesCount string) []byte {
	key := make([]byte, len(AppliedLeavesKey)+len(leaveID)+len(leavesCount))
	copy(key, AppliedLeavesKey)
	copy(key[len(AppliedLeavesKey):], []byte(leaveID))
	copy(key[len(AppliedLeavesKey)+len(leaveID):], []byte(leavesCount))
	return key
}

func LeavesCounterKey(id string) []byte {
	key := make([]byte, len(CounterKey)+len(id))
	copy(key, CounterKey)
	copy(key[len(CounterKey):], []byte(id))
	return key
}
