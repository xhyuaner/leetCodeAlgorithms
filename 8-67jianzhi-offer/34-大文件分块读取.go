package jianzhi

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 带缓冲的读入
func main() {
	file, err := os.Open("largeFilepath/file.txt")
	if err != nil {
		fmt.Println("文件打开失败" + err.Error())
		return
	}

	defer func() {
		file.Close()
		fmt.Println("文件关闭")
	}()

	// 创建带缓冲的读取器
	reader := bufio.NewReader(file)
	// 使用bufio.NewReader().Read()方法进行分块读取文件，并且创建一个指定大小的byte切片的缓冲区来存放每一块文件
	bytes := make([]byte, 6)
	for {
		// Read不会，单独返回EOF
		c, err := reader.Read(bytes)
		if err == nil {
			// 对每块文件执行操作，可以启动多个goroutine进行处理
			fmt.Print(string(bytes[:c]))
		} else if err == io.EOF {
			break
		} else {
			fmt.Println("出错！！")
			return
		}
	}

	fmt.Println("读取完毕")

	//for {
	//	str, err := reader.ReadString('\n')
	//	if err == nil {
	//		fmt.Print(str)
	//	} else if err == io.EOF {
	//		// ReadString最后会同EOF和最后的数据一起返回
	//		//fmt.Println(str)
	//		break
	//	} else {
	//		fmt.Println("出错！！")
	//		return
	//	}
	//}
}
