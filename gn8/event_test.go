package gn8

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPluginGN8_IsDND(t *testing.T) {
	t.Run("未跨天", func(t *testing.T) {
		p := PluginGN8{Conf: Config{DNDDuration: []int{1, 8}}}

		now := time.Now()
		ti := time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.Local)

		assert.True(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local)
		assert.True(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local)
		assert.False(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		assert.False(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, time.Local)
		assert.False(t, p.IsDND(ti))

	})

	t.Run("跨天", func(t *testing.T) {
		p := PluginGN8{Conf: Config{DNDDuration: []int{23, 8}}}

		now := time.Now()
		ti := time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.Local)

		assert.True(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local)
		assert.True(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local)
		assert.False(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		assert.True(t, p.IsDND(ti))

		ti = time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.Local)
		assert.False(t, p.IsDND(ti))

	})
}
