package znet

import "github.com/geoffqiu/zinx-tbit/ziface"

// Request 请求
type Request struct {
	conn ziface.IConnection // 已经和客户端建立好的 链接
	msg  ziface.IMessage    // 客户端请求的数据
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 获取请求的消息的ID
func (r *Request) GetMsgID() uint8 {
	return r.msg.GetMsgID()
}

// GetSn 获取请求的消息的序列号
func (r *Request) GetSn() uint8 {
	return r.msg.GetSn()
}
