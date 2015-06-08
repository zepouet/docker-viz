package dockertype

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/fsouza/go-dockerclient"
	"strconv"
	"strings"
	"fmt"
)

type Container struct {
	docker.APIContainers
}

// return the ID of Container
func (c Container) GetId() string {
	return c.ID
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
	fmt.Println(c)
	return strconv.FormatInt(c.SizeRw, 10)
}

// return the links of Container
func (c Container) GetLink() *hashset.Set {
	i, _ := LoadContainerInfos(c.ID)
	links := hashset.New()
	for _, link := range i.HostConfig.Links {
		linkSlpit := strings.Split(link, ":")
		containerLinked, err := LoadContainerInfos(linkSlpit[0][1:len(linkSlpit[0])])

		if(err) {
			continue
		}

		links.Add(containerLinked.ID)
	}

	return links
}

// return the volume link (volume_from) of Container
func (c Container) GetVolumeFrom() *hashset.Set {
	i, _ := LoadContainerInfos(c.ID)
	links := hashset.New()

	for _, link := range i.HostConfig.VolumesFrom {
		containerLinked, err := LoadContainerInfos(link)

		if(err) {
			continue
		}
		links.Add(containerLinked.ID)
	}

	return links
}