package system

type ErrorLog struct {
	Id         uint64
	Type       uint `gorm:"comment:type 1:返回错误 2:代码错误"`
	ErrContent string
	ErrData    []byte `gorm:"type:mediumblob;serializer:gob"`
}
