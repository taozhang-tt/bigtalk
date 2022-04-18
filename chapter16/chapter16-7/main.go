package main

import "fmt"

// 相当于Context
type Work struct {
	state    State
	Hour     int
	Finished bool
}

// 相当于Request()
func (w *Work) WriteProgram() {
	w.state.WriteProgram(*w)
}

func (w *Work) SetState(s State) {
	w.state = s
}

func NewWork() *Work {
	return &Work{
		state: new(ForenoonState),
	}
}

type State interface {
	// 相当于Handle()
	WriteProgram(Work)
}

// 上午工作类
type ForenoonState struct{}

func (fs *ForenoonState) WriteProgram(w Work) {
	if w.Hour < 12 {
		fmt.Printf("当前时间：%v 点，上午工作，精神百倍。\n", w.Hour)
	} else {
		w.SetState(new(NoonState))
		w.WriteProgram()
	}
}

// 中午工作类
type NoonState struct{}

func (ns *NoonState) WriteProgram(w Work) {
	if w.Hour < 13 {
		fmt.Printf("当前时间：%v 点，饿了，午饭，犯困，午休。\n", w.Hour)
	} else {
		w.SetState(new(AfternoonState))
		w.WriteProgram()
	}

}

// 下午工作类
type AfternoonState struct{}

func (as *AfternoonState) WriteProgram(w Work) {
	if w.Hour < 17 {
		fmt.Printf("当前时间：%v 点，下午状态还不错，继续努力。\n", w.Hour)
	} else {
		w.SetState(new(EveningState))
		w.WriteProgram()
	}
}

// 晚上工作类
type EveningState struct{}

func (es *EveningState) WriteProgram(w Work) {
	if w.Finished {
		w.SetState(new(RestState))
		w.WriteProgram()
	} else {
		if w.Hour < 21 {
			fmt.Printf("当前时间：%v 点，加班哦，疲累至极。\n", w.Hour)
		} else {
			w.SetState(new(SleepingState))
			//w.WriteProgram()
		}
	}
}

// 睡眠状态类
type SleepingState struct{}

func (ss *SleepingState) WriteProgram(w Work) {
	fmt.Printf("当前时间：%v 点，不行了，睡着了。\n", w.Hour)
}

// 休息状态
type RestState struct{}

func (rs *RestState) WriteProgram(w Work) {
	fmt.Printf("当前时间：%v 点，下班回家了。\n", w.Hour)
}

func main() {
	work := NewWork()
	work.Hour = 9
	work.WriteProgram()
	work.Hour = 10
	work.WriteProgram()
	work.Hour = 11
	work.WriteProgram()
	work.Hour = 12
	work.WriteProgram()
	work.Hour = 13
	work.WriteProgram()
	work.Hour = 14
	work.WriteProgram()
	work.Hour = 17

	work.Finished = true

	work.WriteProgram()
	work.Hour = 19
	work.WriteProgram()
	work.Hour = 22
	work.WriteProgram()
}
