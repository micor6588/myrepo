package main

import (
	"ChatRoom/common/message"
	process2 "ChatRoom/server/process"
	"ChatRoom/server/utils"
	"fmt"
	"io"
	"net"
)

//先创建一个Processor 的结构体
type Processor struct {
	Conn net.Conn
}

// ServerProcessMessage 编写一个ServerProcessMessage函数
// ServerProcessMessage 功能：根据客户端发送信息的种类不同，决定调用哪个函数来处理
func (pro *Processor) ServerProcessMessage(mes *message.Message) (err error) {
	switch mes.MessageType {
	case message.LoginMessageType:
		//处理登录逻辑
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessLogin(mes) // type : data
	case message.RegisterMesssageType:
		//处理注册的相关逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) process2() (err error) {

	// 循环的向客户端发送信息
	for {
		//这里我们将读取数据包，直接将其封装成ReadPackage()，返回Message,err
		//创建一个Transfer实例,完成任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPackage()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return err
			} else {
				fmt.Println("read package err")
				return err
			}

		}
		// fmt.Println("mes=", mes)
		err = this.ServerProcessMessage(&mes)
		if err != nil {
			fmt.Println("消息预处理错误，err=", err)
			return err
		}

	}
}
