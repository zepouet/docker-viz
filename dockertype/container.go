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

// return the ID of Container
func (c Container) GetId() string {
	return c.Id
}

// return the ID of Container image base
func (c Container) GetFatherId() string {
	return c.Image
}

// return the Name of Container
func (c Container) GetName() string {
	return c.Names[0]
}

// return the Hard disk size of Container
func (c Container) GetSize() string {
	return strconv.FormatInt(c.SizeRw, 10)
}

// return the links of Container
func (c Container) GetLink() *hashset.Set {
	i, _ := LoadContainerInfos(c.Id)
	links := hashset.New()


	for _, link := range i.HostConfig.Links {
		linkSlpit := strings.Split(link, ":")
		containerLinked, err := LoadContainerInfos(linkSlpit[0])

		if(err) {
			continue
		}
		links.Add(containerLinked.Id)
	}

	return links
}

// return the volume link (volume_from) of Container
func (c Container) GetVolumeFrom() *hashset.Set {
	i, _:= LoadContainerInfos(c.Id)
	links := hashset.New()

	for _, link := range i.HostConfig.VolumesFrom {
		containerLinked, err := LoadContainerInfos(link)

		if(err) {
			continue
		}
		links.Add(containerLinked.Id)
	}

	return links
}