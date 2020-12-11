package mandalorian

import (
	"context"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework"
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

func (m *Mandalorian) Score(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.Infof("Scoring Node: %v while scheduling Pod: %v/%v",nodeName,pod.GetNamespace(),pod.GetName())
	// TODO: Write Your Score Policy here.
	// ...
	return 0,nil
}

func (m *Mandalorian) NormalizeScore(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	var (
		highest int64 = 0
		lowest        = scores[0].Score
	)

	for _, nodeScore := range scores {
		if nodeScore.Score < lowest {
			lowest = nodeScore.Score
		}
		if nodeScore.Score > highest {
			highest = nodeScore.Score
		}
	}

	if highest == lowest {
		lowest --
	}

	// Set Range to [0-100]
	for i, nodeScore := range scores {
		scores[i].Score = (nodeScore.Score - lowest) * framework.MaxNodeScore / (highest - lowest)
		klog.Infof("Node: %v, Score: %v in Plugin: Mandalorian When scheduling Pod: %v/%v", scores[i].Name, scores[i].Score,pod.GetNamespace(), pod.GetName())
	}
	return nil
}

func (m *Mandalorian) ScoreExtensions() framework.ScoreExtensions {
	return m
}