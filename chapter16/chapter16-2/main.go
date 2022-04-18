package main

// 函数版

import "fmt"

var (
	Hour         int
	WorkFinished bool
)

func main() {
	Hour = 9
	WriteProgram()
	Hour = 10
	WriteProgram()
	Hour = 12
	WriteProgram()
	Hour = 13
	WriteProgram()
	Hour = 14
	WriteProgram()
	Hour = 17
	WriteProgram()

	//WorkFinished = true
	WorkFinished = false

	Hour = 19
	WriteProgram()
	Hour = 22
	WriteProgram()
}

func WriteProgram() {
	if Hour < 12 {
		fmt.Printf("当前时间：%v 上午工作，精神百倍\n", Hour)
	} else if Hour < 13 {
		fmt.Printf("当前时间：%v 饿了，午饭，犯困，午休\n", Hour)
	} else if Hour < 17 {
		fmt.Printf("当前时间：%v 下午状态不错，继续努力\n", Hour)
	} else {
		if WorkFinished {
			fmt.Printf("当前时间：%v 下班回家了\n", Hour)
		} else {
			if Hour < 21 {
				fmt.Printf("当前时间：%v 加班哦，疲劳至极\n", Hour)
			} else {
				fmt.Printf("当前时间：%v 不行了，睡着了\n", Hour)
			}
		}
	}
}
