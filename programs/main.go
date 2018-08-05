package main

import (
	"math"
	"log"
	"QLearning"
	"time"
)


type Env struct {
	state float64
	length float64
}

func (e *Env) step (action int) (newState float64, reward float64, done bool) {
	if action == 0 {
		e.state = math.Max(0, e.state - 1)
	} else {
		e.state = math.Min(e.length - 1, e.state + 1)
	}
	done = false
	newState = e.state
	reward = -1
	if e.state == e.length - 1 {
		done = true
	}
	return
}

func (e *Env) render() {
	graph := ""
	graph += "|"
	currentIndex := int(e.state)
	for i, length := 0, int(e.length); i < length; i++ {
		if i == currentIndex {
			if i % 2 == 0 {
				graph += "\\"
			} else  {
				graph += "/"
			}
		} else {
			graph += "-"
		}
	}
	graph += "|"
	log.Print(graph)
}

func (e *Env) reset() {
	e.state = 0
}

func NewEnv (length float64) *Env {
	return &Env{
		state: 0,
		length: length,
	}
}

func main()  {
	env := NewEnv(20)
	env.render()

	agent := QLearning.NewAgent(2)
	agent.EpsilonDecay = 0.998
	doneCount := 0

	for i := 0; i < 10000; i++ {
		log.Printf("第 %v 回合", i)
		isDone := false
		env.reset()
		for !isDone {
			action := agent.ChooseAction(env.state)
			oldState := env.state
			newState, reward, done := env.step(action)
			isDone = done
			agent.UpdateQ(oldState, action, reward, newState)
			if doneCount >= 300 {
				env.render()
				time.Sleep(300 * time.Millisecond)
				log.Print(agent.Q)
			}
			if done {
				log.Print("游戏结束！")
				doneCount += 1
			}
		}
	}
}
