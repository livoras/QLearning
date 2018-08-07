package envs

import (
	"log"
	"fmt"
)

var defaultActions = [][]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

type FrozenSnake struct {
	State []int // "x,y"
	Actions [][]int
	Size int
	Holes [][]int
}

func NewFrozenSnake () *FrozenSnake {
	size := 10
	return &FrozenSnake{
		State: []int{0, 0},
		Actions: defaultActions,
		Size: 10,
		Holes: [][]int{
			{size / 2 - 2, size / 2 - 2},
			{1, 1},
		},
	}
}

func (v *FrozenSnake) Step (actionIndex int) (string, float64, bool) {
	action := v.Actions[actionIndex]
	v.State[0] += action[0]
	v.State[1] += action[1]

	if v.State[0] < 0 { v.State[0] = 0 }
	if v.State[1] < 0 { v.State[1] = 0 }
	if v.State[0] >= v.Size { v.State[0] = v.Size - 1 }
	if v.State[1] >= v.Size { v.State[1] = v.Size - 1 }

	strState := v.GetState()
	if v.checkIsInHole(v.State[0], v.State[1]) {
		return strState, -100, true
	}
	if v.State[0] == v.Size - 1 && v.State[1] == v.Size - 1 {
		return strState, 100, true
	}
	return strState, -1, false
}

func (v *FrozenSnake) Render () {
	graph := "\n"
	for i := 0; i < v.Size; i++ {
		for j := 0; j < v.Size; j++ {
			graph += v.getCharByPosition(i, j)
		}
		graph += "\n"
	}
	log.Print(graph)
}

func (v *FrozenSnake) Reset ()  {
	v.State[0] = 0
	v.State[1] = 0
}

func (v *FrozenSnake) getCharByPosition(x int, y int) string {
	//if x == 0 && y == 0 {
	//	return "O"
	//}
	if v.State[0] == x && v.State[1] == y {
		return "~"
	}
	if x == v.Size - 1 && y == v.Size - 1 {
		return "G"
	}
	if v.checkIsInHole(x, y) {
		return "H"
	}
	return "X"
}

func (v *FrozenSnake) checkIsInHole (x int, y int) bool {
	for _, hole := range v.Holes {
		if hole[0] == x && hole[1] == y{
			return true
		}
	}
	return false
}

func (v *FrozenSnake) GetState() string {
	return fmt.Sprintf("(%v, %v)", v.State[0], v.State[1])
}
