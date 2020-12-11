package mandalorian

import (
	"context"

	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

func (m *Mandalorian) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.Infof("Filter Node: %v while Scheduling Pod: %v/%v. ",nodeInfo.Node().GetName(),pod.GetNamespace(),pod.GetName())
	// TODO: Write Your Filter Policy here.
	// ..
	return nil
}
