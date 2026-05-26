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

package framework

import (
	"flag"

	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/component-base/featuregate"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

var (
	DefaultMutableQOSExtPluginFG featuregate.MutableFeatureGate = featuregate.NewFeatureGate()
	DefaultQOSExtPluginsFG       featuregate.FeatureGate        = DefaultMutableQOSExtPluginFG

	defaultQOSExtPluginsFG = map[featuregate.Feature]featuregate.FeatureSpec{}

	globalExtensionPlugins = map[featuregate.Feature]ExtensionPlugin{}
)

func RegisterQOSExtPlugin(feature featuregate.Feature, featureSpec featuregate.FeatureSpec,
	plugin ExtensionPlugin) error {
	_ = "STUB: not implemented"
	return nil
}

type ExtensionPlugin interface {
	InitFlags(fs *flag.FlagSet)
	Setup(client clientset.Interface, metricCache metriccache.MetricCache, statesInformer statesinformer.StatesInformer)
	Run(stopCh <-chan struct{})
}

type QOSExtensionConfig struct {
	FeatureGates map[string]bool
}

func (c *QOSExtensionConfig) InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func SetupPlugins(client clientset.Interface, metricCache metriccache.MetricCache, statesInformer statesinformer.StatesInformer) {
	_ = "STUB: not implemented"
	return
}

func StartPlugins(cfg *QOSExtensionConfig, stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}
