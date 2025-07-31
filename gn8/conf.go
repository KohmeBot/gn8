package gn8

type Config struct {
	// 免打扰时间段 (起止时间)
	DNDDuration []int `yaml:"dnd_duration"`
}
