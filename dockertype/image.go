package dockertype

import (
	"github.com/samalba/dockerclient"
	"strconv"
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

func (i Image) GetRam() string {
	return "0"
}