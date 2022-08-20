package main

import "fmt"

type Student struct {
	name string
}

func judge() {

}

func (s Student) calAvg(data []float64) (avgResult float64) {
	sum := 0.0

	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func main() {
	a001 := Student{"Kojima"}
	fmt.Println(a001.calAvg([]float64{70, 80, 90}))
}
