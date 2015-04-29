package flare

import (
	"github.com/samalba/dockerclient"
	"strconv"
)

type(
	FlareContainers interface {
		BubbleContainers() string;
	}
)

// Load all containers information in Docker
func GenerateDockerContainerList() map[string]dockerclient.Container {

	containers := make(map[string]dockerclient.Container)
	for _, c := range LoadDockerContainers() {
		containers[c.Id] = c
	}

	return containers
}

// Create the json architecture
func MakeJsonContainers(dockerContainerList map[string]dockerclient.Container) string {
	var flare string
	var i int = 0
	for _, container := range dockerContainerList {
		virtualSize := strconv.Itoa(int(container.SizeRw))
		if virtualSize == "0" {
			continue
		}
		if i != 0 {
			flare += ", "
		}
		flare += "{\"name\": \"" + container.Names[0] + "\", \"size\": " + virtualSize + "}"
		i++
	}

	return flare
}

// Returns the full json for docker containers bubble diagram
func BubbleContainers() string {
	dockerContainerList := GenerateDockerContainerList()

	return  "{\"name\": \"Docker\", \"children\": [" + MakeJsonContainers(dockerContainerList) + "]}"
}