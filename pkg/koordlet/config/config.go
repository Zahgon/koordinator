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

package config

import (
	"flag"

	"k8s.io/client-go/rest"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/audit"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	maframework "github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/prediction"
	qmframework "github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks"
	statesinformerimpl "github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer/impl"
)

const (
	DefaultKoordletConfigMapNamespace = "koordinator-system"
	DefaultKoordletConfigMapName      = "koordlet-config"

	CMKeyQoSPluginExtraConfigs = "qos-plugin-extra-configs"
)

type Configuration struct {
	ConfigMapName      string
	ConfigMapNamesapce string
	KubeRestConf       *rest.Config
	StatesInformerConf *statesinformerimpl.Config
	CollectorConf      *maframework.Config
	MetricCacheConf    *metriccache.Config
	QOSManagerConf     *qmframework.Config
	RuntimeHookConf    *runtimehooks.Config
	AuditConf          *audit.Config
	PredictionConf     *prediction.Config

	FeatureGates map[string]bool
}

func NewConfiguration() *Configuration { _ = "STUB: not implemented"; return nil }

func (c *Configuration) InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (c *Configuration) InitKubeConfigForKoordlet(kubeAPIQPS float64, kubeAPIBurst int) error {
	_ = "STUB: not implemented"
	return nil
}
