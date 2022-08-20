package main

import "fmt"

type Student struct {
	name string
}

func (s Student) calAvg(data []float64) (avgResult float64) {
	sum := 0.0

	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func (s Student) judge(avg float64) (judgeResult string) {
	if avg >= 60 {
		judgeResult = "合格"
	} else if avg <= 30 {
		judgeResult = "赤点"
	} else {
		judgeResult = "不合格"
	}
	return
}

func main() {
	a001 := Student{"Kojima"}
	data := []float64{90, 90, 90, 30, 20, 40}
	var avg float64 = a001.calAvg(data)
	result := a001.judge(avg)
	fmt.Println(a001.name, "の平均点は、", avg, "点です")
	fmt.Println(a001.name + "は、" + result)
}
