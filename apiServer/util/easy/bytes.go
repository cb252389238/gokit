package easy

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 二进制加密
func BinaryEncode(message string) ([]byte, error) {
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

// 二进制 解码消息
func BinaryDecode(message []byte) (string, error) {
	// 读取消息的长度
	reader := bytes.NewReader(message)
	r := bufio.NewReader(reader)
	lengthByte, _ := r.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(r.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = r.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
