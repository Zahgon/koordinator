//go:build linux
// +build linux

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

package tc

func (p *tcPlugin) EnsureIptables() error { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) DelIptables() error { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) iptablesExisted() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// looks like this one:
// -A POSTROUTING -m set --match-set mid_class src -j CLASSIFY --set-class 0001:0003
