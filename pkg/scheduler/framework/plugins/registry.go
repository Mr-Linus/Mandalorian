/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugins

import (
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/defaultbinder"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/defaultpreemption"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/imagelocality"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/interpodaffinity"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/nodeaffinity"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/nodelabel"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/nodename"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/nodeports"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/nodepreferavoidpods"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/noderesources"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/nodeunschedulable"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/nodevolumelimits"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/podtopologyspread"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/queuesort"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/selectorspread"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/serviceaffinity"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/tainttoleration"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/volumebinding"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/volumerestrictions"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/plugins/volumezone"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/framework/runtime"
	"github.com/NJUPT-ISL/Mandalorian/pkg/scheduler/mandalorian"
)

// NewInTreeRegistry builds the registry with all the in-tree plugins.
// A scheduler that runs out of tree plugins can register additional plugins
// through the WithFrameworkOutOfTreeRegistry option.
func NewInTreeRegistry() runtime.Registry {
	return runtime.Registry{
		selectorspread.Name:                        selectorspread.New,
		imagelocality.Name:                         imagelocality.New,
		tainttoleration.Name:                       tainttoleration.New,
		nodename.Name:                              nodename.New,
		nodeports.Name:                             nodeports.New,
		nodepreferavoidpods.Name:                   nodepreferavoidpods.New,
		nodeaffinity.Name:                          nodeaffinity.New,
		podtopologyspread.Name:                     podtopologyspread.New,
		nodeunschedulable.Name:                     nodeunschedulable.New,
		noderesources.FitName:                      noderesources.NewFit,
		noderesources.BalancedAllocationName:       noderesources.NewBalancedAllocation,
		noderesources.MostAllocatedName:            noderesources.NewMostAllocated,
		noderesources.LeastAllocatedName:           noderesources.NewLeastAllocated,
		noderesources.RequestedToCapacityRatioName: noderesources.NewRequestedToCapacityRatio,
		volumebinding.Name:                         volumebinding.New,
		volumerestrictions.Name:                    volumerestrictions.New,
		volumezone.Name:                            volumezone.New,
		nodevolumelimits.CSIName:                   nodevolumelimits.NewCSI,
		nodevolumelimits.EBSName:       nodevolumelimits.NewEBS,
		nodevolumelimits.GCEPDName:     nodevolumelimits.NewGCEPD,
		nodevolumelimits.AzureDiskName: nodevolumelimits.NewAzureDisk,
		nodevolumelimits.CinderName:    nodevolumelimits.NewCinder,
		interpodaffinity.Name:          interpodaffinity.New,
		nodelabel.Name:                 nodelabel.New,
		serviceaffinity.Name:           serviceaffinity.New,
		queuesort.Name:                 queuesort.New,
		defaultbinder.Name:             defaultbinder.New,
		defaultpreemption.Name:         defaultpreemption.New,
		mandalorian.Name:               mandalorian.New,
	}
}
