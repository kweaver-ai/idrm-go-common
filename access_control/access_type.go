package access_control

type AccessType int32

const (
	NONE       AccessType = 0
	GET_ACCESS AccessType = 1 << (iota - 1)
	POST_ACCESS
	PUT_ACCESS
	DELETE_ACCESS
)

var AllAccessType = map[AccessType]struct{}{
	GET_ACCESS:    {},
	POST_ACCESS:   {},
	PUT_ACCESS:    {},
	DELETE_ACCESS: {},
}

func (a AccessType) ToInt32() int32 {
	return int32(a)
}

//Exist return  b exist a AccessType
func (a AccessType) Exist(b int32) bool {
	return (a.ToInt32() & b) == a.ToInt32() //>0
}

//ExistBatch return multiple b all exist a AccessType
func (a AccessType) ExistBatch(args ...int32) bool {
	for _, e := range args { //<=0
		if (a.ToInt32() & e) != a.ToInt32() { //false a.ToInt32() & e == 0
			return false
		}
	}
	return true
}

//Add return  b add a AccessType
func (a AccessType) Add(b int32) int32 {
	return a.ToInt32() | b
}

//Reduce return b reduce a AccessType, if b has not exist a AccessType return self,include  multiple a AccessType
func (a AccessType) Reduce(b int32) int32 {
	if a.ToInt32()&b == a.ToInt32() {
		return b ^ a.ToInt32()
	}
	return b
}

func Verify(a int32) bool {
	if a == NONE.ToInt32() {
		return true
	}
	var tmp = NONE
	for accessType, _ := range AllAccessType {
		tmp = AccessType(tmp.Add(accessType.ToInt32()))
	}
	if a > tmp.ToInt32() {
		return false
	}
	for accessType, _ := range AllAccessType {
		if (accessType.ToInt32() & a) == accessType.ToInt32() {
			return true
		}
	}

	return false
}
