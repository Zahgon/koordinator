/*
Copyright 2022 The Koordinator Authors.

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

package nodenumaresource

import (
	nrtv1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	nrtinformers "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions"
	fwktype "k8s.io/kube-scheduler/framework"
)

type nodeResourceTopologyEventHandler struct {
	topologyManager TopologyOptionsManager
}

func registerNodeResourceTopologyEventHandler(informerFactory nrtinformers.SharedInformerFactory, topologyManager TopologyOptionsManager) error {
	_ = "STUB: not implemented"
	return nil
}

func initNRTInformerFactory(handle fwktype.Handle) (nrtinformers.SharedInformerFactory, error) {
	_ = "STUB: not implemented"
	return *new(nrtinformers.SharedInformerFactory), nil
}

func (m *nodeResourceTopologyEventHandler) OnAdd(obj interface{}, isInInitialList bool) {
	_ = "STUB: not implemented"
	return
}

func (m *nodeResourceTopologyEventHandler) OnUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (m *nodeResourceTopologyEventHandler) OnDelete(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (m *nodeResourceTopologyEventHandler) updateNodeResourceTopology(oldNodeResTopology, newNodeResTopology *nrtv1alpha1.NodeResourceTopology) {
	_ = "STUB: not implemented"
	return
}

// Give other plugins a chance to customize a different MaxRefCount
