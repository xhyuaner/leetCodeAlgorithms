package jianzhi

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadBigFile 带缓冲的读入
func ReadBigFile(filePath string, chunkSize int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建带缓冲的读取器
	reader := bufio.NewReader(file)
	// 使用bufio.NewReader().Read()方法进行分块读取文件，并且创建一个指定大小的byte切片的缓冲区来存放每一块文件
	buffer := make([]byte, chunkSize)
	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		// 对每块文件执行操作，可以启动多个goroutine进行处理
		processChunk(buffer[:bytesRead])
	}
	return nil
}

func processChunk(chunk []byte) {
	// 处理每一块数据的逻辑
	fmt.Printf("Read %d bytes：\n", len(chunk))
	// 可以在这里对块数据进行进一步处理
	fmt.Println(string(chunk))
}
