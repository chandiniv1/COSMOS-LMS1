package types

const (
	ModuleName = "lms"

	StoreKey = ModuleName

	// RouterKey = ModuleName
	// QuerierRoute = ModuleName
)

var (
	AdminKey    = []byte{0x01}
	StudentKey  = []byte{0x02}
	LeavesKey   = []byte{0x03}
	sequenceKey = []byte{0x04}
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

func LeavesStoreKey(admin string,leaveID string) []byte {
	key := make([]byte, len(LeavesKey)+len(admin)+len(sequenceKey)+len(leaveID))
	copy(key, LeavesKey)
	copy(key[len(LeavesKey):], admin)
	copy(key,sequenceKey)
	copy(key[len(sequenceKey):],leaveID)
	return key
}
