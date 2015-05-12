package dockertype

import "github.com/emirpasic/gods/sets/hashset"

type DockerType interface {
	GetId() string
	GetFatherId() string
	GetName() string
	GetSize() string
	GetRam() string
	GetLink() *hashset.Set
}