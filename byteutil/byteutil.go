package byteutil

import (
	"encoding/binary"
)

// IntToBytes 将int转换为[]byte
// 注意：此函数始终返回8字节的切片，因为内部将int转换为int64处理
func IntToBytes(n int) []byte {
	// 预分配固定大小的字节切片（int64需要8字节）
	b := make([]byte, 8)
	// 直接使用binary.BigEndian.PutUint64写入数据
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}

// BytesToInt 将[]byte转换为int
// 如果输入的字节切片长度不足8字节，将返回0
func BytesToInt(bys []byte) int {
	// 检查输入字节切片长度是否为8字节
	// 如果不是8字节，返回0以避免panic
	if len(bys) != 8 {
		return 0
	}
	// 直接使用binary.BigEndian.Uint64读取字节切片
	return int(binary.BigEndian.Uint64(bys))
}

const (
	// 32的十六进制表示，用于大小写转换
	caseBit = 0x20
	// 大写转换掩码，等于^caseBit但适用于byte类型
	caseMask = 0xDF
	// ASCII字符的最大值
	maxASCII = ''
)

// ByteToLower 将字节转换为小写
// 使用位运算实现大写到小写的转换，比加减法更高效
func ByteToLower(b byte) byte {
	// 只处理ASCII字符
	if b <= maxASCII {
		// 使用位运算将大写字母转换为小写字母
		// ASCII中大写字母与小写字母相差32(0x20)
		if 'A' <= b && b <= 'Z' {
			b |= caseBit
		}
	}
	return b
}

// ByteToUpper 将字节转换为大写
// 使用位运算实现小写到大写的转换，比加减法更高效
func ByteToUpper(b byte) byte {
	// 只处理ASCII字符
	if b <= maxASCII {
		// 使用位运算将小写字母转换为大写字母
		// 清除第6位(0x20)来实现大写转换
		if 'a' <= b && b <= 'z' {
			b &= caseMask // 等同于b &= ^caseBit，但避免了整数溢出问题
		}
	}
	return b
}
