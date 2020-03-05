package model

import (
	"time"
)

type ObjectMeta struct {
	Name              string
	Namespace         string
	CreationTimestamp time.Time
	DeletionTimestamp *time.Time
	Labels            map[string]string
}

type Namespace struct {
	ObjectMeta
	NamespacePhase string
}

type NamespaceList struct {
	Items []Namespace
}
