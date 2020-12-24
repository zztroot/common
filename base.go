package common

const (
	errorSliceType = "两个切片类型不同"
	errorSliceTypeCannotDispose = "该类型切片无法处理"
)

type allTools struct {
	Slice SliceDispose
}


func GetTools() *allTools{
	return &allTools{}
}