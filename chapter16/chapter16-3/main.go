package main

// 分类版

import "fmt"

func main() {
	w := new(Work)
	w.Hour = 9
	w.WriteProgram()
	w.Hour = 10
	w.WriteProgram()
	w.Hour = 12
	w.WriteProgram()
	w.Hour = 13
	w.WriteProgram()
	w.Hour = 14
	w.WriteProgram()
	w.Hour = 17
	w.WriteProgram()

	w.WorkFinished = true
	//w.WorkFinished = false

	w.Hour = 19
	w.WriteProgram()
	w.Hour = 22
	w.WriteProgram()
}

type Work struct {
	Hour         int
	WorkFinished bool
}

func (w *Work) WriteProgram() {
	if w.Hour < 12 {
		fmt.Printf("当前时间：%v 上午工作，精神百倍\n", w.Hour)
	} else if w.Hour < 13 {
		fmt.Printf("当前时间：%v 饿了，午饭，犯困，午休\n", w.Hour)
	} else if w.Hour < 17 {
		fmt.Printf("当前时间：%v 下午状态不错，继续努力\n", w.Hour)
	} else {
		if w.WorkFinished {
			fmt.Printf("当前时间：%v 下班回家了\n", w.Hour)
		} else {
			if w.Hour < 21 {
				fmt.Printf("当前时间：%v 加班哦，疲劳至极\n", w.Hour)
			} else {
				fmt.Printf("当前时间：%v 不行了，睡着了\n", w.Hour)
			}
		}
	}
}
