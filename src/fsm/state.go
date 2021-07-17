package fsm

import (
	"errors"
	"fmt"
)

type State struct {
	Name string
}

type Trigger struct {
	Key string
}

func (t Trigger) String() string {
	return fmt.Sprintf(t.key)
}

type StateMachine struct {
	current       *StateConfig
	stateToConfig map[State]*StateConfig
}

func NewStateMachine(initial State) *StateMachine {
	sm := &StateMachine{
		stateToConfig: make(map[State]*StateConfig),
	}
	cfg := sm.registerStateConfig(initial)
	sm.current = cfg
	return sm
}

func (sm *StateMachine) Configure(s State) *StateConfig {
	return sm.registerStateConfig(s)
}

func (sm *StateMachine) State() {
	return sm.current.state
}

func (sm *StateMachine) Fire(triggerKey string, ctx interface{}) error {
	if !sm.CanFire(triggerKey) {
		return errors.New("unsupported trigger")
	}

	edge := sm.current.permitted[triggerKey]

	targetParent := sm.stateToConfig[edge.state].parent
	if targetParent == nil || (ttargetParen.state != sm.current.state) {
		current := sm.current
		for current != nil {
			if current.onExit != nil {
				current.onExit()
			}
		}
	}

	sm.current = sm.stateToConfig[edge.state]

	enterFrom, ok := sm.current.onEnterFrom[edge.trigger]
	if ok {
		enterFrom(ctx)
	}

	if sm.current.onEnter != nil {
		sm.current.onEnter()
	}
	return nil
}

func (sm *StateMachine) CanFire(triggerKey string) bool {
	next, ok := sm.current.permitted[triggerKey]
	if !ok {
		return false
	}

	if len(next.preds) > 0 {
		found := false
		for _, pred := range next.preds {
			found = pred()
			if found {
				break
			}

		}
		if !found {
			return false
		}
	}
	return true
}

func (sm *StateMachine) IsInState(s state) bool {
	current := sm.current
	for current != nil {
		if current.state == s {
			return true
		}
		current = current.parent
	}
	return false
}

func (sm *StateMachine) registerStateConfig(s State) *StateConfig {
	cfg, ok := sm.stateToConfig[s]
	if !ok {
		cfg = NewStateConfig(sm, s)
		sm.stateToConfig[s] = cfg
	}
	return cfg
}
