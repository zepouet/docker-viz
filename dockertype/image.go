package dockertype

import (
	"strconv"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/fsouza/go-dockerclient"
)

type Image struct {
	docker.APIImages
}

// return the ID of Image
func (i Image) GetId() string {
	return i.ID
}

// return the ID of father's Image
func (i Image) GetFatherId() string {
	return i.ParentID
}

// return the Name of Image
func (i Image) GetName() string {
	return i.RepoTags[0]
}

// return the Hard Disk size of Image
func (i Image) GetSize() string {
	return strconv.FormatInt(i.VirtualSize, 10)
}

// return a void Hashset
func (i Image) GetLink() *hashset.Set {
	return hashset.New()
}

// return a void Hashset
func (i Image) GetVolumeFrom() *hashset.Set {
	return hashset.New()
}