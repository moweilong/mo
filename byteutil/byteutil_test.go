package byteutil

import (
	"fmt"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	fmt.Println(IntToBytes(1))
	fmt.Println(BytesToInt(IntToBytes(1)))
}

func TestBytesToIntWithInvalidLength(t *testing.T) {
	// 测试空切片
	emptyBytes := []byte{}
	result1 := BytesToInt(emptyBytes)
	fmt.Printf("Empty bytes result: %d\n", result1)
	
	// 测试长度小于8的切片
	shortBytes := []byte{1, 2, 3}
	result2 := BytesToInt(shortBytes)
	fmt.Printf("Short bytes (len=3) result: %d\n", result2)
	
	// 测试长度大于8的切片
	longBytes := make([]byte, 10)
	for i := range longBytes {
		longBytes[i] = byte(i)
	}
	result3 := BytesToInt(longBytes)
	fmt.Printf("Long bytes (len=10) result: %d\n", result3)
}

func TestByteToLower(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{{
		input:    'A',
		expected: 'a',
	}, {
		input:    'a',
		expected: 'a',
	}, {
		input:    'Z',
		expected: 'z',
	}, {
		input:    '1',
		expected: '1',
	}, {
		input:    '!',
		expected: '!',
	}, {
		input:    byte(128), // 非ASCII字符
		expected: byte(128),
	}}
	
	for _, tt := range tests {
		result := ByteToLower(tt.input)
		if result != tt.expected {
			t.Errorf("ByteToLower(%c) = %c; want %c", tt.input, result, tt.expected)
		}
		fmt.Printf("ByteToLower(%c) = %c\n", tt.input, result)
	}
}

func TestByteToUpper(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{{
		input:    'a',
		expected: 'A',
	}, {
		input:    'A',
		expected: 'A',
	}, {
		input:    'z',
		expected: 'Z',
	}, {
		input:    '1',
		expected: '1',
	}, {
		input:    '!',
		expected: '!',
	}, {
		input:    byte(128), // 非ASCII字符
		expected: byte(128),
	}}
	
	for _, tt := range tests {
		result := ByteToUpper(tt.input)
		if result != tt.expected {
			t.Errorf("ByteToUpper(%c) = %c; want %c", tt.input, result, tt.expected)
		}
		fmt.Printf("ByteToUpper(%c) = %c\n", tt.input, result)
	}
}
