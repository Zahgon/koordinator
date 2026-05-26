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

package system

type CoreSched struct{}

func NewCoreSched() CoreSchedInterface { _ = "STUB: not implemented"; return *new(CoreSchedInterface) }

func NewCoreSchedExtended() CoreSchedExtendedInterface {
	_ = "STUB: not implemented"
	return *new(CoreSchedExtendedInterface)
}

func (s *CoreSched) Get(pidType CoreSchedScopeType, pid uint32) (uint64, error) {
	_ = "STUB: not implemented"
	// NOTE: pidType only support Thread type.
	return 0, nil
}

func (s *CoreSched) Create(pidType CoreSchedScopeType, pid uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CoreSched) ShareTo(pidType CoreSchedScopeType, pid uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CoreSched) ShareFrom(pidType CoreSchedScopeType, pid uint32) error {
	_ = "STUB: not implemented"
	// NOTE: pidTypeFrom only support Thread type.
	return nil
}

type CoreSchedExtendedResult struct {
	FailedPIDs []uint32
	Error      error
}

func (s *CoreSched) clear(pidType CoreSchedScopeType, pids ...uint32) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CoreSched) Clear(pidType CoreSchedScopeType, pids ...uint32) ([]uint32, error) {
	_ = "STUB: not implemented"
	// keep the outside goroutine with cookie 0, then we can reset the target pid's cookie by ShareTo the cookie of
	// the new goroutine
	// TODO: directly use syscall when the kernel supports Clear (0x1000)
	return nil, nil
}

func (s *CoreSched) assign(pidTypeFrom CoreSchedScopeType, pidFrom uint32, pidTypeTo CoreSchedScopeType, pidsTo ...uint32) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CoreSched) Assign(pidTypeFrom CoreSchedScopeType, pidFrom uint32, pidTypeTo CoreSchedScopeType, pidsTo ...uint32) ([]uint32, error) {
	_ = "STUB: not implemented"
	// keep the outside goroutine with cookie 0, then we can assign the pidFrom's cookie to pidTo by
	// NOTE: pidTypeFrom only support Thread type.
	// 1. ShareFrom the pidFrom's cookie to the new goroutine
	// 2. ShareTo the new goroutine's cookie to the target pidTo
	return nil, nil
}

// ProbeCoreSchedIfEnabled checks if the MAINLINE kernel support the core scheduling.
// Since there's no direct API, we probe by calling prctl and checking its return value.
func ProbeCoreSchedIfEnabled() bool { _ = "STUB: not implemented"; return false }
