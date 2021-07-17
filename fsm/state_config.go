package fsm

type edge struct {
	trigger Trigger
	state   State
	preds   []func() bool
}

type StateConfig struct {
	owner       *StateMachine
	onEnter     func()
	onEnterFrom map[Trigger]func(interface{})
	onExit      func()
	state       State
	parent      *StateConfig
	permitted   map[string]*edge
}

func newStateConfig(sm *StateMachine, s State) *StateConfig {
	return &StateConfig{
		owner:       sm,
		state:       s,
		onEnterFrom: make(map[Trigger]func(interface{})),
		permitted:   make(map[string]*edge),
	}
}

func (c *StateConfig) Permit(t Trigger, s State) *StateConfig {
	c.owner.registerStateConfig(s)
	c.permitted[t.Key] = &edge{trigger: t, state: s}
	return c
}

func (c *StateConfig) PermitIf(t Trigger, s State, pred func() bool) *StateConfig {
	c.owner.registerStateConfig(s)

	val, ok := c.permitted[t.Key]
	if !ok {
		val = &edge{trigger: t, state: s}
		c.permitted[t.Key] = val
	}

	val.preds = append(val.preds, pred)
	return c
}

func (c *StateConfig) OnEnter(f func()) *StateConfig {
	c.OnEnter = f
	return c
}

func (c *StateConfig) OnEnterFrom(t Trigger, f func(interface{})) *StateConfig {
	c.OnEnterFrom[t] = f
	return c
}

func (c *StateConfig) OnExit(f func()) *StateConfig {
	c.OnExit = f
	return c
}

func (c *StateConfig) SubstateOf(s State) *StateConfig {
	cfg := c.owner.registerStateConfig(s)
	c.parent = cfg
	return c
}
