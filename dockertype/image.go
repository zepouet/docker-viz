package dockertype

import (
	"github.com/samalba/dockerclient"
	"strconv"
	"github.com/emirpasic/gods/sets/hashset"
)

type Image struct {
	dockerclient.Image
}

func (i Image) GetId() string {
	return i.Id
}

func (i Image) GetFatherId() string {
	return i.ParentId
}

func (i Image) GetName() string {
	return i.RepoTags[0]
}

func (i Image) GetSize() string {
	return strconv.FormatInt(i.VirtualSize, 10)
}

func (i Image) GetLink() *hashset.Set {
	return hashset.New()
}

func (i Image) GetVolumeFrom() *hashset.Set {
	return hashset.New()
}