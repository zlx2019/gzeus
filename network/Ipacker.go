/**
  @author: Zero
  @date: 2023/3/4 16:05:13
  @desc:

**/

package network

import "io"

// IPacker 数据解包与封包接口
type IPacker interface {
	Pack(message *Message) ([]byte, error)
	UnPack(reader io.Reader) (*Message, error)
}
