package fsm_test

import (
	"fsm"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSettingInitialState(t *testing.T) {
	s := fsm.State{Name: "foo"}
	sm := fsm.NewStateMachine(s)
	require.Equal(t, s, sm.State())
}
