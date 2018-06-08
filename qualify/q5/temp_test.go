package main

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

//3 5
//1 3 800 18:00 21:00
//1 2 650 13:30 14:00
//2 3 100 14:00 18:00
//2 3 300 14:30 19:00
//2 3 200 15:00 19:30
//output:950
func TestTemp(t *testing.T) {
	edge := []*train{
		new(train).init(1, 3, 800, fmtTime("18:00"), fmtTime("21:00")),
		new(train).init(1, 2, 650, fmtTime("13:30"), fmtTime("14:00")),
		new(train).init(2, 3, 100, fmtTime("14:00"), fmtTime("18:00")),
		new(train).init(2, 3, 300, fmtTime("14:30"), fmtTime("19:00")),
		new(train).init(2, 3, 200, fmtTime("15:00"), fmtTime("19:30")),
	}
	temp(3, edge)
}

//3 5
//1 2 1000 0:00 12:00
//1 2 100 0:30 14:00
//1 2 100 0:30 15:00
//2 3 300 16:00 24:00
//2 3 200 16:30 24:00
//output:1300
func TestTemp1(t *testing.T) {
	edge := []*train{
		new(train).init(1, 2, 1000, fmtTime("0:00"), fmtTime("12:00")),
		new(train).init(1, 2, 100, fmtTime("0:30"), fmtTime("14:00")),
		new(train).init(1, 2, 100, fmtTime("0:30"), fmtTime("15:00")),
		new(train).init(2, 3, 300, fmtTime("16:00"), fmtTime("24:00")),
		new(train).init(2, 3, 200, fmtTime("16:30"), fmtTime("24:00")),
	}
	temp(3, edge)
}

//3 4
//1 2 100 0:30 14:00
//1 2 200 0:30 15:00
//2 3 300 16:00 24:00
//2 3 200 16:30 24:00
//output:-1
func TestTemp2(t *testing.T) {
	edge := []*train{
		new(train).init(1, 2, 100, fmtTime("0:30"), fmtTime("14:00")),
		new(train).init(1, 2, 200, fmtTime("0:30"), fmtTime("15:00")),
		new(train).init(2, 3, 300, fmtTime("16:00"), fmtTime("24:00")),
		new(train).init(2, 3, 200, fmtTime("16:30"), fmtTime("24:00")),
	}
	temp(3, edge)
}

//3 4
//1 2 100 0:30 14:00
//1 2 200 1:00 16:00
//2 3 300 16:00 24:00
//2 3 200 16:30 24:00
//output:400
func TestTemp3(t *testing.T) {
	edge := []*train{
		new(train).init(1, 2, 100, fmtTime("0:30"), fmtTime("14:00")),
		new(train).init(1, 2, 200, fmtTime("1:00"), fmtTime("16:00")),
		new(train).init(2, 3, 300, fmtTime("16:00"), fmtTime("24:00")),
		new(train).init(2, 3, 200, fmtTime("16:30"), fmtTime("24:00")),
	}
	temp(3, edge)
}

//3 4
//1 2 100 0:30 14:00
//1 2 200 1:00 16:30
//2 3 300 16:00 24:00
//2 3 200 16:30 24:00
//output:-1
func TestTemp4(t *testing.T) {
	edge := []*train{
		new(train).init(1, 2, 100, fmtTime("0:30"), fmtTime("14:00")),
		new(train).init(1, 2, 200, fmtTime("1:00"), fmtTime("16:30")),
		new(train).init(2, 3, 300, fmtTime("16:00"), fmtTime("24:00")),
		new(train).init(2, 3, 200, fmtTime("16:30"), fmtTime("24:00")),
	}
	temp(3, edge)
}

//3 5
//1 3 800 18:00 21:00
//1 2 650 13:30 14:00
//1 2 350 13:30 14:00
//1 2 150 13:30 14:00
//2 3 100 14:00 18:00
//2 3 300 14:30 19:00
//2 3 200 15:00 19:30
//output:450
func TestTemp5(t *testing.T) {
	edge := []*train{
		new(train).init(1, 3, 800, fmtTime("18:00"), fmtTime("21:00")),
		new(train).init(1, 2, 650, fmtTime("13:30"), fmtTime("14:00")),
		new(train).init(1, 2, 350, fmtTime("13:30"), fmtTime("14:00")),
		new(train).init(1, 2, 150, fmtTime("13:30"), fmtTime("14:00")),
		new(train).init(2, 3, 100, fmtTime("14:00"), fmtTime("18:00")),
		new(train).init(2, 3, 300, fmtTime("14:30"), fmtTime("19:00")),
		new(train).init(2, 3, 200, fmtTime("15:00"), fmtTime("19:30")),
	}
	temp(3, edge)
}

//output:300
func TestTemp6(t *testing.T) {
	edge := []*train{
		new(train).init(1, 3, 50, fmtTime("18:30"), fmtTime("19:00")),
		new(train).init(1, 4, 600, fmtTime("18:00"), fmtTime("21:00")),
		new(train).init(1, 2, 150, fmtTime("12:00"), fmtTime("12:30")),
		new(train).init(2, 3, 100, fmtTime("13:30"), fmtTime("18:00")),
		new(train).init(3, 4, 50, fmtTime("20:00"), fmtTime("20:30")),
		new(train).init(3, 4, 50, fmtTime("21:00"), fmtTime("21:30")),
		new(train).init(2, 4, 200, fmtTime("15:00"), fmtTime("19:30")),
	}
	temp(4, edge)
}

//output: 80
func TestTemp7(t *testing.T) {
	edge := []*train{
		new(train).init(1, 3, 600, fmtTime("9:00"), fmtTime("9:30")),
		new(train).init(2, 1, 600, fmtTime("9:30"), fmtTime("10:00")),
		new(train).init(1, 3, 600, fmtTime("10:30"), fmtTime("11:00")),
		new(train).init(1, 2, 30, fmtTime("4:00"), fmtTime("5:00")),
		new(train).init(2, 4, 50, fmtTime("5:30"), fmtTime("9:00")),
		new(train).init(3, 4, 50, fmtTime("7:00"), fmtTime("7:30")),
		new(train).init(3, 4, 50, fmtTime("10:00"), fmtTime("10:30")),
		new(train).init(3, 4, 50, fmtTime("11:30"), fmtTime("12:30")),
		new(train).init(4, 2, 200, fmtTime("8:00"), fmtTime("8:30")),
	}
	temp(4, edge)
}

//output: 200
func TestTemp8(t *testing.T) {
	edge := []*train{
		new(train).init(1, 4, 200, fmtTime("3:00"), fmtTime("3:30")),
		new(train).init(1, 4, 200, fmtTime("5:00"), fmtTime("5:30")),
		new(train).init(4, 4, 200, fmtTime("4:00"), fmtTime("4:30")),
		new(train).init(4, 4, 200, fmtTime("5:00"), fmtTime("5:30")),
	}
	temp(4, edge)
}

func TestTempMax(t *testing.T) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	n := r.Intn(5) + 2
	m := r.Intn(20) + 1
	edge := make([]*train, m, m)
	for i := range edge {
		s := r.Intn(n) + 1
		ta := r.Intn(n) + 1
		c := r.Intn(1000)
		ts := r.Intn(48)
		td := r.Intn(48)
		edge[i] = new(train).init(s, ta, c, ts, td)
	}
	fmt.Println("n", n, "path", m)
	fmt.Println("seed", seed)
	temp(n, edge)
}

func BenchmarkTemp(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := 500
		m := 15000
		edge := make([]*train, m, m)
		for i := range edge {
			s := rand.Intn(n) + 1
			t := rand.Intn(n) + 1
			c := rand.Intn(1000)
			ts := rand.Intn(48)
			td := rand.Intn(48) + ts
			edge[i] = new(train).init(s, t, c, ts, td)
		}
		b.StartTimer()
		temp(n, edge)
	}
}
