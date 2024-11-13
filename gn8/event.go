package gn8

import "time"

// IsDND 判断当前是否在免打扰时间段内
func (p *PluginGN8) IsDND(t time.Time) bool {
	durs := p.Conf.DNDDuration
	if len(durs) < 2 {
		return false
	}

	if durs[0] < durs[1] {
		return t.Hour() >= durs[0] && t.Hour() <= durs[1]
	}
	return t.Hour() >= durs[0] || t.Hour() <= durs[1]

}
