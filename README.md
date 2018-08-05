# QLearning
A simple QLeaning Agent in Golang

## Usage

```golang
func main()  {
	env := NewEnv(20)
	env.render()

    /** Initialize an agent with actions space */
	agent := QLearning.NewAgent(2)
	
	for i := 0; i < 10000; i++ {
		isDone := false
		env.reset()
		for !isDone {
		    /** Choose action per to state */
			action := agent.ChooseAction(env.state)
			oldState := env.state
			newState, reward, done := env.step(action)
			isDone = done
			/** Learn from interaction with the environment */
			agent.UpdateQ(oldState, action, reward, newState)
		}
	}
}

```
## License
MIT
