package fsm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	fsm "github.com/tommyagk/vgame/fsm"
)

func TestSettingInitialState(t *testing.T) {
	s := fsm.State{Name: "foo"}
	sm := fsm.NewStateMachine(s)
	require.Equal(t, s, sm.State())
}
