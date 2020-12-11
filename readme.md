# Mandalorian
Mandalorian is a scheduler scaffold based on kubernetes code v1.20.
It is used to facilitate developers to quickly implement their 
own scheduler algorithm and adapt to the existing scheduling 
strategy of kubernetes scheduler.


## Get Started 
### 1. Implement your filtering method
```go
// pkg/scheduler/mandalorian/filtering.go

func (m *Mandalorian) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.Infof("Filter Node: %v while Scheduling Pod: %v/%v. ",nodeInfo.Node().GetName(),pod.GetNamespace(),pod.GetName())
	// TODO: Write Your Filter Policy here.
	// ..
	return nil
}
```

### 2. Implement your scoring method

```go
// pkg/scheduler/mandalorian/scoring.go

func (m *Mandalorian) Score(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.Infof("Scoring Node: %v while scheduling Pod: %v/%v",nodeName,pod.GetNamespace(),pod.GetName())
	// TODO: Write Your Score Policy here.
	// ...
	return 0,nil
}
```

### 3. Run Your Scheduler
- edit scheduler config file `./config/config.yaml`
```yaml
apiVersion: kubescheduler.config.k8s.io/v1beta1
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig:  "your k8s config path"
leaderElection:
  leaderElect: true
  resourceName: mandalorian
  resourceNamespace: kube-system
profiles:
  - schedulerName: mandalorian
```
- run your scheduler
```shell 
go run ./cmd/kube-scheduler/scheduler.go --config=./config/config.yaml --v=3
```
