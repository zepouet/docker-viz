package dockertype

import (
	"github.com/samalba/dockerclient"
	"strconv"
	"github.com/emirpasic/gods/sets/hashset"
)

type Image struct {
	dockerclient.Image
}

// return the ID of Image
func (i Image) GetId() string {
	return i.Id
}

// return the ID of father's Image
func (i Image) GetFatherId() string {
	return i.ParentId
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