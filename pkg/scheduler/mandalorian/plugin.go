package mandalorian

import (
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	Name = "Mandalorian"
	ScoreWeight = 2
)


var _ framework.FilterPlugin = &Mandalorian{}
var _ framework.ScorePlugin = &Mandalorian{}

type Mandalorian struct {
	handle    framework.Handle
}

func (m *Mandalorian) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &Mandalorian{
		handle: h,
	}, nil
}
