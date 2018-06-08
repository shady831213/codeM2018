package main

import (
	"math"
	"bufio"
	"os"
	"strconv"
	"strings"
	"container/heap"
)

func main() {

	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	s, _, _ := input.ReadLine()
	line := strings.Split(string(s), " ")
	n, _ := strconv.Atoi(line[0])
	m, _ := strconv.Atoi(line[1])
	g := new(graph).init(n)
	h := &edgesHeap{}
	xCntMap, yCntMap := make([]bool, n+1, n+1), make([]bool, n+1, n+1)
	xCnt, yCnt := 0, 0
	for i := 0; i < m; i++ {
		arr := make([]int, 3)
		s, _, _ := input.ReadLine()
		line := strings.Split(string(s), " ")
		arr[0], _ = strconv.Atoi(line[0])
		arr[1], _ = strconv.Atoi(line[1])
		arr[2], _ = strconv.Atoi(line[2])
		heap.Push(h, arr)
		if !xCntMap[arr[0]] {
			xCntMap[arr[0]] = true
			xCnt++
		}
		if !yCntMap[arr[1]] {
			yCntMap[arr[1]] = true
			yCnt++
		}
	}
	l := n
	if xCnt < l {
		l = xCnt
	}
	if yCnt < l {
		l = yCnt
	}
	for i := 0; i < n; i++ {
		result := -1
		for g.matches < i+1 {
			if h.Len() == 0 || i >= l {
				result = -1
				break
			}
			lastResult := result
			edge := heap.Pop(h).([]int)
			result = edge[2]
			g.addEdge(edge[0], edge[1])
			if g.xCnt < i+1 || g.yCnt < i+1 {
				continue
			}

			if result != lastResult || g.matches < i+1 {
				g.maxMatch()
			}
		}
		output.WriteString(strconv.Itoa(result))
		if i == n-1 {
			output.WriteString("\n")
		} else {
			output.WriteString(" ")
		}
	}
	output.Flush()
}

func temp(n int, edges [][]int) {
	g := new(graph).init(n)
	h := &edgesHeap{}
	output := bufio.NewWriter(os.Stdout)
	xCntMap, yCntMap := make([]bool, n+1, n+1), make([]bool, n+1, n+1)
	xCnt, yCnt := 0, 0
	for i := range edges {
		heap.Push(h, edges[i])
		if !xCntMap[edges[i][0]] {
			xCntMap[edges[i][0]] = true
			xCnt++
		}
		if !yCntMap[edges[i][1]] {
			yCntMap[edges[i][1]] = true
			yCnt++
		}
	}
	l := n
	if xCnt < l {
		l = xCnt
	}
	if yCnt < l {
		l = yCnt
	}
	for i := 0; i < n; i++ {
		result := -1
		for g.matches < i+1 {
			if h.Len() == 0 || i >= l {
				result = -1
				break
			}
			lastResult := result
			edge := heap.Pop(h).([]int)
			result = edge[2]
			g.addEdge(edge[0], edge[1])
			if g.xCnt < i+1 || g.yCnt < i+1 {
				continue
			}

			if result != lastResult || g.matches < i+1 {
				g.maxMatch()
			}
		}
		output.WriteString(strconv.Itoa(result))
		if i == n-1 {
			output.WriteString("\n")
		} else {
			output.WriteString(" ")
		}
	}
	output.Flush()
}

type edgesHeap [][]int

func (h edgesHeap) Len() int           { return len(h) }
func (h edgesHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h edgesHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *edgesHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *edgesHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0: n-1]
	return x
}

//Hopcroft-Karp
type graph struct {
	n                int
	x                [][]int
	dis              int
	xMatch, yMatch   []int
	xLevel, yLevel   []int
	matches          int
	queue            []int
	yVisit           []bool
	xCntMap, yCntMap []bool
	xCnt, yCnt       int
}

func (g *graph) init(n int) *graph {
	g.n = n
	g.x = make([][]int, g.n+1)
	for i := range g.x {
		g.x[i] = make([]int, 0, g.n+1)
	}
	g.matches = 0
	g.xMatch, g.yMatch = make([]int, g.n+1, g.n+1), make([]int, g.n+1, g.n+1)
	g.queue = make([]int, g.n, g.n)
	g.xCntMap, g.yCntMap = make([]bool, n+1, n+1), make([]bool, n+1, n+1)
	g.xCnt, g.yCnt = 0, 0
	g.yVisit = make([]bool, g.n+1, g.n+1)
	g.xLevel, g.yLevel = make([]int, g.n+1, g.n+1), make([]int, g.n+1, g.n+1)
	return g
}

func (g *graph) addEdge(i, j int) {
	g.x[i] = append(g.x[i], j)
	if !g.xCntMap[i] {
		g.xCntMap[i] = true
		g.xCnt++
	}
	if !g.yCntMap[j] {
		g.yCntMap[j] = true
		g.yCnt++
	}
}

func (g *graph) bfs() bool {
	g.dis = math.MaxInt32
	for i := range g.xLevel {
		g.xLevel[i] = 0
		g.yLevel[i] = 0
	}
	pf, pe := 0, 0
	//use queue

	for x := 1; x < g.n+1; x++ {
		if g.xMatch[x] == 0 && len(g.x[x]) != 0 {
			g.queue[pe] = x
			g.xLevel[x] = 0
			pe ++
		}
	}
	for pf != pe {
		s := g.queue[pf]
		pf++
		if g.xLevel[s] > g.dis {
			break
		}
		for _, y := range g.x[s] {
			if g.yLevel[y] == 0 {
				g.yLevel[y] = g.xLevel[s] + 1
				if g.yMatch[y] == 0 {
					g.dis = g.yLevel[y]
				} else {
					g.xLevel[g.yMatch[y]] = g.yLevel[y] + 1;
					g.queue[pe] = g.yMatch[y]
					pe ++
				}
			}
		}
	}
	//fmt.Println(g.dis, g.xLevel, g.yLevel, g.xMatch, g.yMatch)
	return g.dis != math.MaxInt32
}

func (g *graph) dfs(x int) bool {
	for _, y := range g.x[x] {
		//fmt.Println(x, y, yVisit, g.yLevel[y], g.xLevel[x], g.xMatch, g.yMatch)
		if !g.yVisit[y] && g.yLevel[y] == g.xLevel[x]+1 {
			g.yVisit[y] = true
			if g.yMatch[y] != 0 && g.yLevel[y] == g.dis {
				continue
			}
			if g.yMatch[y] == 0 || g.dfs(g.yMatch[y]) {
				g.xMatch[x] = y
				g.yMatch[y] = x
				return true
			}
		}
	}
	return false
}

func (g *graph) maxMatch() {
	for g.bfs() {
		for i := range g.yVisit {
			g.yVisit[i] = false
		}
		for x := 1; x < g.n+1; x++ {
			if g.xMatch[x] == 0 && len(g.x[x]) != 0 {
				if g.dfs(x) {
					g.matches++
				}
				//fmt.Println(matches, k, x, g.xMatch, g.yMatch, yVisit, g.yLevel, g.xLevel)
				//fmt.Println(g.x)
				//if g.matches >= k {
				//	return
				//}
			}
		}
	}
}

//func (g *graph) dfs1(x int) bool {
//	for _, y := range g.x[x] {
//		//fmt.Println(x, y, yVisit, g.yLevel[y], g.xLevel[x], g.xMatch, g.yMatch)
//		if !g.yVisit[y] {
//			g.yVisit[y] = true
//			if g.yMatch[y] == 0 || g.dfs1(g.yMatch[y]) {
//				g.xMatch[x] = y
//				g.yMatch[y] = x
//				return true
//			}
//		}
//	}
//	return false
//}
//
//func (g *graph) maxMatch1() {
//	for x := 1; x < g.n+1; x++ {
//		for i := range g.yVisit {
//			g.yVisit[i] = false
//		}
//		if g.xMatch[x] == 0 && len(g.x[x]) != 0 {
//			if g.dfs1(x) {
//				g.matches++
//			}
//		}
//	}
//}
