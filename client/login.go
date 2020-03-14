//构建用户登陆的功能

package main

import (
	"ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//写一个函数，完成登录功能
func login(userID int, userPwd string) (err error) {
	//下一个就要开始定协议
	fmt.Printf(" useID=%d  userPwd=%s", userID, userPwd)

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8886")
	if err != nil {
		fmt.Println("net .Dial err=", err)
		return
	}

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.MessageType = message.LoginMessageType
	//3.创建一个loginMessage结构体
	var loginMessage message.LoginMessage
	loginMessage.UserID = userID
	loginMessage.UserPwd = userPwd

	//4.将loginMessage序列化
	data, err := json.Marshal(loginMessage)
	if err != nil {
		fmt.Println("json.Marshar err=", err)
		return
	}
	//5.将data赋值给mes.MessageData字段
	mes.MessageData = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshar err=", err)
		return
	}

	//7.这个时候data就是我们要发送的消息
	//7.1先把data的长度发送给服务器
	//先获得data的长度-->转成一个表示长度的byte切片
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
	fmt.Println("客户端发送消息长度欧克")

	//这里还需要处理服务器返回的消息
	//休眠20秒
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20秒")
	mes, err = ReadPackage(conn)
	if err != nil {
		fmt.Println("readPackage(conn) err=", err)
		return
	}

	//将message的data部分反序列化成LoginResponseMessage
	var loginResponseMessage message.LoginResponMessage
	err = json.Unmarshal([]byte(mes.MessageData), &loginResponseMessage)
	if loginResponseMessage.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResponseMessage.Code == 500 {
		fmt.Println(loginResponseMessage.Error)
	}

	return
}
