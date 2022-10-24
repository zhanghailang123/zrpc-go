package codec

import "io"

type Header struct {
	//服务名和方法名
	ServiceMethod string
	seq           uint64
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

// 定义类型 在原类型上创造新的类型 函数式接口？？？
type NewCodecFunc func(closer io.ReadWriteCloser) Codec

// 为string定义一个别名 Type
type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
