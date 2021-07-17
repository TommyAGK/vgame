package vgame_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	vgame "github.com/tommyagk/vgame"
)

func Test_Start_Loop(t *testing.T) {
	l := vgame.Game_loop()
	assert.Equal(t, true, l)
}
