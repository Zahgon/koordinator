package core

import (
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
)

type GangSummary struct {
	Name                   string           `json:"name"`
	WaitTime               time.Duration    `json:"waitTime"`
	CreateTime             time.Time        `json:"createTime"`
	Mode                   string           `json:"mode"`
	GangMatchPolicy        string           `json:"gangMatchPolicy"`
	MinRequiredNumber      int              `json:"minRequiredNumber"`
	TotalChildrenNum       int              `json:"totalChildrenNum"`
	GangGroup              []string         `json:"gangGroup"`
	Children               sets.Set[string] `json:"children"`
	PendingChildren        sets.Set[string] `json:"pendingChildren"`
	WaitingForBindChildren sets.Set[string] `json:"waitingForBindChildren"`
	BoundChildren          sets.Set[string] `json:"boundChildren"`
	OnceResourceSatisfied  bool             `json:"onceResourceSatisfied"`
	GangGroupInfo          *GangGroupInfo   `json:"gangGroupInfo"`
	GangFrom               string           `json:"gangFrom"`
	HasGangInit            bool             `json:"hasGangInit"`
}

func (gang *Gang) GetGangSummary() *GangSummary { _ = "STUB: not implemented"; return nil }
