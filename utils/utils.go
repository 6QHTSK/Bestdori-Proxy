package utils

func GetListSliced(List []int, offset uint, limit uint) (SlicedList []int) {
	if offset >= uint(len(List)) {
		return SlicedList
	}
	end := offset + limit
	if end > uint(len(List)) {
		end = uint(len(List))
	}
	return List[offset:end]
}
