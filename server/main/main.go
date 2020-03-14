package main

import (
	"fmt"
	"net"
)


//处理服务器和客户端之间的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//这里调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误，err=", err)
		return
	}

}

func main() {
	//提示信息
	fmt.Println("服务器在8886端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8886")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err= ", err)
		return
	}
	//一旦监听成功,就等待客户端来连接服务器
	for {
		fmt.Println("----------等待客户端连接服务器-----------")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功，就启动一个协程和客户端保持通信
		go process(conn)
	}

}



/*
// ReadPackage 读取数据包中的内容(将读取包的任务交给readpackage这个函数)
func ReadPackage(conn net.Conn) (mes message.Message, err error) {
	//这里我们读取数据包，直接封装成一个函数readPkg()
	buf := make([]byte, 8096)
	//conn.Read在conn没有关闭的情况下，才会阻塞
	fmt.Println("读取客户端的相关信息")
	n, err := conn.Read((buf[:4]))
	if n != 4 || err != nil {
		err = errors.New("conn.Read  header faild ")
		return
	}
	//依据buf[:4]转成uint32类型
	var packageLength uint32
	packageLength = binary.BigEndian.Uint32(buf[0:4])
	//依据包的长度packageLength,读取消息内容
	n, err = conn.Read(buf[:packageLength])
	if n != int(packageLength) || err != nil {
		err = errors.New("conn.Read  body faild ")
		return
	}

	//把packageLength反序列化——>message.Meaaage
	err = json.Unmarshal(buf[:packageLength], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return

}

// WritePackage 发送数据包中的内容(将读取包的任务交给WritePackage这个函数）
func WritePackage(conn net.Conn, data []byte) (err error) {
	//先发送一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	num, err := conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("length send err=", err)
		return
	}

	if num != 4 || err != nil {
		fmt.Println("conn write(bytes) err=", err)
		return
	}
	fmt.Println("客户端，发送消息长度成功")
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}

// ServerProcessLogin 编写一个函数，专门处理登录请求
func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码
	//1.先从message中取出message.Data,并直接反序列化成LoginMessage
	var loginMeaasge message.LoginMessage
	err = json.Unmarshal([]byte(mes.MessageData), loginMeaasge)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//先声明一个返回信息
	var responseMessage message.Message
	responseMessage.MessageType = message.LoginResponceMessageType

	//再声明一个loginResMessage,并完成赋值
	var loginResponsegeMessage message.LoginResponMessage

	//再声明一个LoginResponMess
	//如果用户的id=100,密码=123456，认为是合法的
	if loginMeaasge.UserID == 100 && loginMeaasge.UserPwd == "123456" {
		//合法
		loginResponsegeMessage.Code = 200
	} else {
		//不合法
		loginResponsegeMessage.Code = 500
		loginResponsegeMessage.Error = "服务器内部错误..."
	}
	//3.将loginResponseMessage序列化
	data, err := json.Marshal(loginResponsegeMessage)
	if err != nil {
		fmt.Println("json.Mashal err=", err)
		return
	}

	//4.将data赋值给responMessage
	responseMessage.MessageData = string(data)

	//5. 将data赋值给responseMessage,并准备发送
	data, err = json.Marshal(responseMessage)
	if err != nil {
		fmt.Println("json.Mashal err=", err)
		return
	}

	//6.发送data,我们将其封装到WritePackage函数当中
	err = WritePackage(conn, data)
	return err

}

// ServerProcessMessage 编写一个ServerProcessMessage函数
// ServerProcessMessage 功能：根据客户端发送信息的种类不同，决定调用哪个函数来处理
func ServerProcessMessage(conn net.Conn, mes *message.Message) (err error) {
	switch mes.MessageType {
	case message.LoginMessageType:
		//处理登录逻辑
	case message.RegisterMesssageType:
		//处理注册的相关逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
*/
