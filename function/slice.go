package function

// 切片工具包
//
//

// CheckInNumberSlice 交易一个int类型值 是否存在于一个切片中
func CheckInNumberSlice[T uint64 | int32](v T, s []T) bool {
	for _, item := range s {
		if v == item {
			return true
		}
	}
	return false
}

// DelElsInSlice 从一个切片中删除指定的值
func DelElsInSlice[T uint64 | int32](v T, old []T) (new []T) {
	for i, item := range old {
		if v == item {
			new = append(old[:i], old[i+1:]...)
			return
		}
	}
	return old
}
