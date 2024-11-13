package gn8sdk

import (
	"github.com/kohmebot/gn8/gn8"
	"github.com/kohmebot/plugin"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockEnv struct {
	plugin.Env
}

func (m *mockEnv) GetPlugin(name string) (p plugin.Plugin, ok bool) {
	return &gn8.PluginGN8{Conf: gn8.Config{DNDDuration: []int{1, 8}}}, true
}

func TestInvoker_IsDND(t *testing.T) {

	invoker, _ := NewInvoker(&mockEnv{})
	now := time.Now()
	res := invoker.IsDND(time.Date(now.Year(), now.Month(), now.Day(), 5, 0, 0, 0, now.Location()))
	assert.True(t, res)
}
