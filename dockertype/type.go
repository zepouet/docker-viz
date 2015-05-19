package dockertype

import "github.com/emirpasic/gods/sets/hashset"

// Generic type for images and containers values
type DockerType interface {
	GetId() string
	GetFatherId() string
	GetName() string
	GetSize() string
	GetLink() *hashset.Set
}