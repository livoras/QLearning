package QLearning

import (
	"math"
	"math/rand"
	"time"
)

type State interface {}

type Agent struct {
	/** 学习率 */
	Alpha float64
	/** 衰减因子 */
	Gamma float64
	/** 探索因子 */
	Epsilon float64
	/** 探索因子衰减度 */
	EpsilonDecay float64
	/** 最小探索因子 */
	EpsilonMin float64
	/** actions 的数量 */
	ActionsSpace int
	/** Q Table */
	Q map[State][]float64
	/** 初始化 Q 值表默认值 */
	QDefaultValue float64
}

/** Agent 构造函数 */
func NewAgent(actionSpace int) *Agent {
	return &Agent{
		Alpha:        0.2,
		Gamma:        0.8,
		Epsilon:      1,
		EpsilonDecay: 0.995,
		EpsilonMin:   0.05,
		/** maps are reference types */
		Q:             map[State][]float64{},
		QDefaultValue: 0,
		ActionsSpace:  actionSpace,
	}
}

/** 更新 Q 值表 */
func (v *Agent) UpdateQ(state State, action int, reward float64, newState State) {
	currentActionsQ := v.getActionsOrSetActionsByState(state)
	_, maxNextStateQVal := argmax(v.getActionsOrSetActionsByState(newState))
	oldQValue := currentActionsQ[action]
	currentActionsQ[action] = (1 - v.Alpha) * oldQValue + v.Alpha* (reward + v.Gamma* maxNextStateQVal)
}

func (v *Agent) ChooseAction(state State) int {
	rand.Seed(time.Now().Unix())
	/** 探索因子会缓慢进行变化 */
	v.Epsilon = math.Max(v.Epsilon * v.EpsilonDecay, v.EpsilonMin)
	if rand.Float64() < v.Epsilon {
		/** 进行随机选择 */
		return rand.Intn(v.ActionsSpace)
	} else {
		/** 选择最大的 */
		actions := v.getActionsOrSetActionsByState(state)
		maxActionIndex, _ := argmax(actions)
		return maxActionIndex
	}
}

func (v *Agent) SaveModel() {
	// TODO
}

func (v *Agent) LoadModel() {
	// TODO
}

/** 根据 state 获取 actions 的 Q 值列表，如果不存在则初始化一个 */
func (v *Agent) getActionsOrSetActionsByState (state State) []float64  {
	/** 如果有值，直接返回 actions */
	if actions, ok := v.Q[state]; ok {
		return actions
	} else {
		v.Q[state] = make([]float64, v.ActionsSpace)
		return v.Q[state]
	}
}

func argmax(nums []float64) (int, float64) {
	var max = nums[0]
	var maxIndex = 0
	for i, num := range nums {
		if num > max {
			max = num
			maxIndex = i
		}
	}
	return maxIndex, max
}

