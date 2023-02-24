package types

const (
	ModuleName = "lms"

	StoreKey = ModuleName

	// RouterKey = ModuleName
	// QuerierRoute = ModuleName
)

var (
	AdminKey   = []byte{0x01}
	StudentKey = []byte{0x02}
	//LeavesKey   = []byte{0x03}
	sequenceKey       = []byte{0x04}
	AcceptedLeavesKey = []byte{0x05}
	AppliedLeavesKey  = []byte{0x06}
)

func AdminStoreKey(address string) []byte {
	key := make([]byte, len(AdminKey)+len(address))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], address)
	return key
}

func StudentStoreKey(studentID string) []byte {
	key := make([]byte, len(StudentKey)+len(studentID))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], studentID)
	return key
}

// func LeavesStoreKey(leaveID string) []byte {
// 	key := make([]byte, len(LeavesKey)+len(leaveID))
// 	copy(key, LeavesKey)
// 	copy(key[len(LeavesKey):], leaveID)
// 	return key
// }

func AcceptedLeavesStoreKey(admin string, leaveID string) []byte {
	key := make([]byte, len(AcceptedLeavesKey)+len(admin)+len(sequenceKey)+len(leaveID))
	copy(key, AcceptedLeavesKey)
	copy(key[len(AcceptedLeavesKey):], admin)
	copy(key, sequenceKey)
	copy(key[len(sequenceKey):], leaveID)
	return key
}
func AppliedLeavesStoreKey(stdntAddress string, leaveID string) []byte {
	key := make([]byte, len(AppliedLeavesKey)+len(stdntAddress)+len(sequenceKey)+len(leaveID))
	copy(key, AppliedLeavesKey)
	copy(key[len(AppliedLeavesKey):], stdntAddress)
	copy(key, sequenceKey)
	copy(key[len(sequenceKey):], leaveID)
	return key
}
