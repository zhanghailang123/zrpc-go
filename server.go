package zrpc_go

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"zrpc-go/codec"
)

const MagicNumber = 0x3befc

// Option 定义 协商的编码方式
type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error", err)
			return
		}
		//TODO
	}
}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

func ServerConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()

	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options server: ", err)
		return
	}

	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic server %x", opt.MagicNumber)
		return
	}

	//f := gob.NewDecoder(conn)
	f := codec.NewCodecFuncMap[opt.CodecType]

	if f == nil {
		log.Printf("rpc server ")
	}
}
