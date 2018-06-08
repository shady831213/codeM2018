package main

import (
	"testing"
	"time"
	"math/rand"
	"bufio"
	"fmt"
	"os/exec"
	"io"
)

//3 7
//1 3 5
//2 3 2
//3 1 7
//1 2 0
//2 3 2
//3 2 0
//2 1 5
//output 0 2 5
func TestTemp(t *testing.T) {
	edges := [][]int{
		[]int{1, 3, 5},
		[]int{2, 3, 2},
		[]int{3, 1, 7},
		[]int{1, 2, 0},
		[]int{2, 3, 2},
		[]int{3, 2, 0},
		[]int{2, 1, 5},
	}
	temp(3, edges)
}

//output 0 2 -1
func TestTemp1(t *testing.T) {
	edges := [][]int{
		[]int{1, 3, 5},
		[]int{2, 3, 2},
		[]int{3, 2, 7},
		[]int{1, 2, 0},
		[]int{2, 3, 2},
		[]int{3, 2, 0},
		[]int{2, 2, 5},
	}
	temp(3, edges)
}

//output 0 2 5
func TestTemp2(t *testing.T) {
	edges := [][]int{
		[]int{3, 3, 5},
		[]int{2, 1, 2},
		[]int{3, 1, 7},
		[]int{1, 2, 1},
		[]int{2, 1, 2},
		[]int{2, 2, 0},
		[]int{2, 3, 5},
	}
	temp(3, edges)
}

//output 5 -1 -1
func TestTemp3(t *testing.T) {
	edges := [][]int{
		[]int{3, 3, 5},
	}
	temp(3, edges)
}

//output 0
func TestTemp4(t *testing.T) {
	edges := [][]int{
		[]int{1, 1, 0},
	}
	temp(1, edges)
}

//output 0 3 -1
func TestTemp5(t *testing.T) {
	subproc := exec.Command("./temp")
	stdin, _ := subproc.StdinPipe()
	stdout, _ := subproc.StdoutPipe()
	defer stdin.Close()
	edges := [][]int{
		[]int{1, 2, 0},
		[]int{1, 1, 2},
		[]int{2, 2, 3},
	}
	subproc.Start()
	io.WriteString(stdin, "3 3\n")
	for i := range edges {
		for j := range edges[i] {
			io.WriteString(stdin, fmt.Sprint(edges[i][j]))
			if j == len(edges[i])-1 {
				io.WriteString(stdin, "\n")
			} else {
				io.WriteString(stdin, " ")
			}
		}
	}
	reader := bufio.NewReader(stdout)
	s, _ := reader.ReadString('\n')
	subproc.Wait()
	fmt.Println(s)
}

//output 0 -1...-1
func TestTemp6(t *testing.T) {
	subproc := exec.Command("./temp")
	stdin, _ := subproc.StdinPipe()
	stdout, _ := subproc.StdoutPipe()
	defer stdin.Close()
	n := 1000
	m := 100000
	edges := make([][]int, m)
	for i := range edges {
		s := (i + 1) % n
		t := 1
		c := i
		edges[i] = []int{s, t, c}
	}
	subproc.Start()
	io.WriteString(stdin, "1000 100000\n")
	for i := range edges {
		for j := range edges[i] {
			io.WriteString(stdin, fmt.Sprint(edges[i][j]))
			if j == len(edges[i])-1 {
				io.WriteString(stdin, "\n")
			} else {
				io.WriteString(stdin, " ")
			}
		}
	}
	reader := bufio.NewReader(stdout)
	s, _ := reader.ReadString('\n')
	subproc.Wait()
	fmt.Println(s)
}

//output 0 1 2...998 -1
func TestTemp7(t *testing.T) {
	n := 1000
	m := 100000
	edges := make([][]int, m)
	for i := range edges {
		s := (i + 1) % n
		t := n - (i+1)%n
		c := i
		edges[i] = []int{s, t, c}
	}
	temp(n, edges)
}

//output:16 35 391 576 816 -1 -1 -1

func TestTempRand(t *testing.T) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(1527819448507607163))
	n := r.Intn(10) + 5
	m := r.Intn(50) + 1
	edges := make([][]int, m)
	for i := range edges {
		s := r.Intn(n) + 1
		t := r.Intn(n) + 1
		c := r.Intn(1000)
		edges[i] = []int{s, t, c}
	}
	temp(n, edges)
	t.Log("n", n)
	for i := range edges {
		t.Log(edges[i])
	}
	t.Log(seed)
}

//0 45 45 51 -1
func TestTempRand1(t *testing.T) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(1527834800255631545))
	n := r.Intn(3) + 5
	m := r.Intn(50) + 1
	edges := make([][]int, m)
	for i := range edges {
		s := r.Intn(n) + 1
		t := r.Intn(n) + 1
		c := r.Intn(100)
		edges[i] = []int{s, t, c}
	}
	temp(n, edges)
	t.Log("n", n)
	for i := range edges {
		t.Log(edges[i])
	}
	t.Log(seed)
}

//func TestTempRand2(t *testing.T) {
//	seed := time.Now().UnixNano()
//	r := rand.New(rand.NewSource(seed))
//	n := r.Intn(1000) + 1
//	m := r.Intn(100000) + 1
//	edges := make([][]int, m)
//	for i := range edges {
//		s := r.Intn(n) + 1
//		t := r.Intn(n) + 1
//		c := r.Intn(1000000000)
//		edges[i] = []int{s, t, c}
//	}
//	temp(n, edges)
//	t.Log("n", n)
//	t.Log(seed)
//}

//func TestTempRand3(t *testing.T) {
//	seed := time.Now().UnixNano()
//	r := rand.New(rand.NewSource(seed))
//	n := 1000
//	m := 100000
//	edges := make([][]int, m)
//	for i := range edges {
//		s := r.Intn(999) + 1
//		t := r.Intn(999) + 1
//		c := r.Intn(1000000000)
//		edges[i] = []int{s, t, c}
//	}
//	temp(n, edges)
//	t.Log("n", n)
//	t.Log(seed)
//}

func BenchmarkTemp(b *testing.B) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := 1000
		m := 100000
		edges := make([][]int, m)
		for i := range edges {
			s := r.Intn(999) + 1
			t := r.Intn(999) + 1
			c := r.Intn(1000000000)
			edges[i] = []int{s, t, c}
		}

		b.StartTimer()
		temp(n, edges)
	}
}

func BenchmarkMain(b *testing.B) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := 1000
		m := 100000
		edges := make([][]int, m)
		for i := range edges {
			s := r.Intn(500) + 1
			t := r.Intn(500) + 1
			c := r.Intn(1000000000)
			edges[i] = []int{s, t, c}
		}
		b.StartTimer()
		subproc := exec.Command("./temp")
		stdin, _ := subproc.StdinPipe()
		defer stdin.Close()
		subproc.Start()
		io.WriteString(stdin, fmt.Sprintf("%d %d\n", n, m))
		for i := range edges {
			for j := range edges[i] {
				io.WriteString(stdin, fmt.Sprint(edges[i][j]))
				if j == len(edges[i])-1 {
					io.WriteString(stdin, "\n")
				} else {
					io.WriteString(stdin, " ")
				}
			}
		}
		subproc.Wait()
	}
}
