package main

import (
	"github.com/kohmebot/gn8/gn8"
	"github.com/kohmebot/plugin"
)

func NewPlugin() plugin.Plugin {
	return gn8.NewPlugin()
}
