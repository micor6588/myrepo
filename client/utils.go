package main

import (
	"ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

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
