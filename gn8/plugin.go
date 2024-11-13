package gn8

import (
	"fmt"
	"github.com/kohmebot/pkg/command"
	"github.com/kohmebot/pkg/version"
	"github.com/kohmebot/plugin"
	"github.com/wdvxdr1123/ZeroBot"
)

type PluginGN8 struct {
	Conf Config
}

func NewPlugin() *PluginGN8 {
	return new(PluginGN8)
}

func (p *PluginGN8) Init(engine *zero.Engine, env plugin.Env) error {
	err := env.GetConf(&p.Conf)
	if err != nil {
		return err
	}

	return nil
}

func (p *PluginGN8) Name() string {
	return "GN8"
}

func (p *PluginGN8) Description() string {
	return "早晚安，免打扰控制"
}

func (p *PluginGN8) Commands() fmt.Stringer {
	return command.NewCommands()
}

func (p *PluginGN8) Version() uint64 {
	return uint64(version.NewVersion(0, 0, 1))
}

func (p *PluginGN8) OnBoot() {

}
