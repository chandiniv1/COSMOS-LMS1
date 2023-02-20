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

func AdminStoreKey(admin string,leaveid string) []byte {
	key := make([]byte, len(AdminKey)+len(admin)+len(sequenceKey)+len(leaveid))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], admin)
	copy(key,sequenceKey)
	copy(key[len(sequenceKey):],leaveid)
	return key
}

func StudentStoreKey(student string) []byte {
	key := make([]byte, len(StudentKey)+len(student))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], student)
	return key
}

func LeavesStoreKey(classID string) []byte {
	key := make([]byte, len(LeavesKey)+len(classID))
	copy(key, LeavesKey)
	copy(key[len(LeavesKey):], classID)
	return key
}
