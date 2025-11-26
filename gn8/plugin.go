package gn8

import (
	"github.com/kohmebot/plugin/v2"
	"github.com/wdvxdr1123/ZeroBot"
)

type PluginGN8 struct {
	Conf Config
}

func NewPlugin() *PluginGN8 {
	return new(PluginGN8)
}

func (p *PluginGN8) OnInit(engine plugin.Engine, env plugin.Env) error {
	err := env.GetConf(&p.Conf)
	if err != nil {
		return err
	}

	return nil
}

func (p *PluginGN8) Name() string {
	return "gn8"
}

func (p *PluginGN8) OnHelp(ctx *zero.Ctx) {

}

func (p *PluginGN8) Version() string {
	return "v0.0.5"
}

func (p *PluginGN8) OnBoot() {

}
