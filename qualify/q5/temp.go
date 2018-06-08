package main

import (
	"container/list"
	"math"
	"fmt"
	"strings"
	"strconv"
)

type train struct {
	s, t           int
	c              int
	ts, td         int
	hasSubstituted bool
	valid          bool
}

func (tr *train) init(s, t int, c int, ts, td int) *train {
	tr.s = s
	tr.t = t
	tr.c = c
	tr.ts = ts
	tr.td = td
	tr.hasSubstituted = false
	tr.valid = false
	return tr
}

type graph struct {
	len         int
	path        []map[int]map[int]map[int]int
	reversePath []map[int]map[int]map[int]int
	valid       []map[int]bool
}

func (g *graph) init(n int) *graph {
	g.len = n + 1
	g.path, g.reversePath = make([]map[int]map[int]map[int]int, g.len), make([]map[int]map[int]map[int]int, g.len)
	g.valid = make([]map[int]bool, g.len)
	for i := range g.path {
		g.path[i] = make(map[int]map[int]map[int]int)
		g.reversePath[i] = make(map[int]map[int]map[int]int)
		g.valid[i] = make(map[int]bool)
	}
	return g
}

func (g *graph) lazyInit(p []map[int]map[int]map[int]int, s, t, ts, td int) {
	if _, ok := p[s][ts]; !ok {
		p[s][ts] = make(map[int]map[int]int)
	}
	if _, ok := p[s][ts][t]; !ok {
		p[s][ts][t] = make(map[int]int)
	}
	if _, ok := p[s][ts][t]; !ok {
		p[s][ts][t] = make(map[int]int)
	}
}

func (g *graph) addTrain(s, t, c, ts, td int) {
	g.lazyInit(g.path, s, t, ts, td)
	if t == g.len-1 {
		g.path[s][ts][t][0] = c
	} else {
		for i := 48; i > td; i-- {
			g.path[s][ts][t][i] = c

		}
	}
	if s == 1 {
		g.lazyInit(g.path, 0, s, 0, ts)
		g.path[0][0][s][ts] = 0
	}

}

func (g *graph) addReverseTrain(s, t, c, ts, td int) {
	if t == g.len-1 {
		g.lazyInit(g.reversePath, t, s, 0, ts)
		g.reversePath[t][0][s][ts] = c
	} else {
		for i := 48; i > td; i-- {
			g.lazyInit(g.reversePath, t, s, i, ts)
			g.reversePath[t][i][s][ts] = c
		}
	}
	if s == 1 {
		g.lazyInit(g.reversePath, s, 0, ts, 0)
		g.reversePath[s][ts][0][0] = 0
	}
}

func (g *graph) reverseDfs(s, ts int) {
	g.valid[s][ts] = true
	for nt := range g.reversePath[s][ts] {
		for ntd := range g.reversePath[s][ts][nt] {
			if !g.valid[nt][ntd] {
				g.reverseDfs(nt, ntd)
			}
		}
	}
}

func (g *graph) removePath() {
	for s := 1; s < g.len; s++ {
		for i := 48; i >= 0; i -- {
			if g.valid[s][i] {
				delete(g.path[s], i)
				break
			}
		}
	}
}

func (g *graph) minCost() int {
	d := make([][]int, g.len)
	for i := range d {
		d[i] = make([]int, 49)
		for j := range d[i] {
			d[i][j] = math.MaxInt32
		}
	}
	d[0][0] = 0
	//use queue
	queue := list.New()
	queue.PushBack(struct{ s, ts int }{0, 0})
	for queue.Len() != 0 {
		s := queue.Front().Value.(struct{ s, ts int })
		for t := range g.path[s.s][s.ts] {
			for td := range g.path[s.s][s.ts][t] {
				if d[t][td] > d[s.s][s.ts]+g.path[s.s][s.ts][t][td] {
					d[t][td] = d[s.s][s.ts] + g.path[s.s][s.ts][t][td]
					queue.PushBack(struct{ s, ts int }{t, td})
				}
			}
		}
		queue.Remove(queue.Front())
	}

	if d[g.len-1][0] == math.MaxInt32 {
		return -1
	}
	return d[g.len-1][0]
}

func (g *graph) run() int {
	g.reverseDfs(g.len-1, 0)
	if !g.valid[0][0] {
		return -1
	}
	//fmt.Println("before remove", g.path)
	g.removePath()
	//fmt.Println(g.path)
	return g.minCost()
}

func temp(n int, trains []*train) {
	g := new(graph).init(n)
	for _, t := range trains {
		g.addTrain(t.s, t.t, t.c, t.ts, t.td)
		g.addReverseTrain(t.s, t.t, t.c, t.ts, t.td)
	}
	fmt.Println(g.run())
}

func main() {
	n, path := 0, 0
	fmt.Scan(&n, &path)
	g := new(graph).init(n)
	for i := 0; i < path; i++ {
		s, t, c, ts, td := 0, 0, 0, "", ""
		fmt.Scan(&s, &t, &c, &ts, &td)
		g.addTrain(s, t, c, fmtTime(ts), fmtTime(td))
		g.addReverseTrain(s, t, c, fmtTime(ts), fmtTime(td))
	}
	fmt.Println(g.run())
}
func fmtTime(s string) int {
	timeS := strings.Split(s, ":")
	timeI0, _ := strconv.Atoi(timeS[0])
	timeI1, _ := strconv.Atoi(timeS[1])
	timeF := timeI0 << 1
	if timeI1 > 0 {
		return timeF + 1
	}
	return timeF
}
