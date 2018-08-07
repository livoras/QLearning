package main

import (
	"QLearning/envs"
	"QLearning"
	"time"
	"log"
)

func main()  {
	env := envs.NewFrozenSnake()
	env.Size = 5

	agent := QLearning.NewAgent(len(env.Actions))

	for i := 0; i < 10000; i++ {
		isDone := false
		env.Reset()
		for !isDone {
			oldState := env.GetState()
			action := agent.ChooseAction(oldState)
			newState, reward, done := env.Step(action)
			agent.UpdateQ(oldState, action, reward, newState)
			isDone = done
			if i >= 100 {
				env.Render()
				time.Sleep(300 * time.Millisecond)
				if isDone {
					log.Print("回合结束！")
					time.Sleep(2000 * time.Millisecond)
				}
			}
		}
	}
}
