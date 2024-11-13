package gn8sdk

import (
	"fmt"
	"github.com/kohmebot/plugin"
	"reflect"
	"time"
)

type Invoker struct {
	val reflect.Value
}

func NewInvoker(env plugin.Env) (*Invoker, error) {
	p, ok := env.GetPlugin("GN8")
	if !ok {
		return nil, fmt.Errorf("GN8 plugin not found")
	}

	val := reflect.ValueOf(p)
	return &Invoker{val: val}, nil
}

// IsDND 判断是否在免打扰时间段内
func (i *Invoker) IsDND(t time.Time) bool {
	return i.val.MethodByName("IsDND").Call([]reflect.Value{reflect.ValueOf(t)})[0].Bool()
}
