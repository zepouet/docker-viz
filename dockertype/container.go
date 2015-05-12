package dockertype

import (
	"github.com/samalba/dockerclient"
	"strconv"
	"strings"
	"github.com/emirpasic/gods/sets/hashset"
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

func (c Container) GetLink() *hashset.Set {
	i := *LoadContainerInfos(c.Id)
	links := hashset.New()
	for _, link := range i.HostConfig.Links {
		linkSlpit := strings.Split(link, ":")
		containerLinked := *LoadContainerInfos(linkSlpit[0])
		links.Add(containerLinked.Id)
	}

	return links
}