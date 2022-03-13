package define2

import (
	"Amanisis/httpclient"
	"Amanisis/model/ServerModel"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

//从管道中读取数据并解码
func process(conn net.Conn) (name string, body string) {
	reader := bufio.NewReader(conn)
	name, err := Decode(reader)
	body, err = Decode(reader)
	if err != nil {
		fmt.Println("decode msg failed, err:", err)
	}
	return
}
func InitDefine2Server() {
	//服务器监听请求
	listen, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
		}
		fmt.Println("here")
		name, msg := process(conn)
		fmt.Println(name, msg)
		serverList := ServerModel.GetServerListByName(name)
		server := serverList[0]
		resp := httpclient.Post(GetUrl(server.Ip, server.Name), msg)
		//服务器返回报文

		data, _ := Encode(resp)
		conn.Write(data)
		conn.Write(data)
	}
	//通信结束
}

//"key:value" "test1.ping"
func GetUrl(ip string, name string) string {
	strings.ReplaceAll(name, ".", "/")
	str := ip + "/" + name
	return str
}
