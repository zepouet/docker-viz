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

func (c Container) GetLink() *hashset.Set {
	i, _ := *LoadContainerInfos(c.Id)
	links := hashset.New()


	for _, link := range i.HostConfig.Links {
		linkSlpit := strings.Split(link, ":")
		containerLinked, err := *LoadContainerInfos(linkSlpit[0])

		if(err) {
			continue
		}
		links.Add(containerLinked.Id)
	}

	return links
}

func (c Container) GetVolumeFrom() *hashset.Set {
	i, _ := *LoadContainerInfos(c.Id)
	links := hashset.New()

	for _, link := range i.HostConfig.VolumesFrom {
		containerLinked, err := *LoadContainerInfos(link)

		if(err) {
			continue
		}
		links.Add(containerLinked.Id)
	}

	return links
}