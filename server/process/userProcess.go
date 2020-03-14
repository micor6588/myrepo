package process2

import (
	"ChatRoom/common/message"
	"ChatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

// ServerProcessLogin 编写一个函数，专门处理登录请求
func (user *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	//由于使用了分层的模式（MVC），先创建了一个Transfer实例,然后读取
	tf := &utils.Transfer{}
	err = tf.WritePackage(data)
	return

}
