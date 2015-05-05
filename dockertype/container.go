package dockertype

import (
	"github.com/samalba/dockerclient"
	"strconv"
)

type Container struct {
	dockerclient.Container
}

func (c Container) GetId() string {
	return c.Id
}

func (c Container) GetFatherId() string {
	return c.Image
}

func (c Container) GetName() string {
	return c.Names[0]
}

func (c Container) GetSize() string {
	return strconv.FormatInt(c.SizeRw, 10)
}

func (c Container) GetRam() string {
	i := *LoadContainerInfos(c.Id)
	return strconv.FormatInt(i.Config.Memory, 10)
}