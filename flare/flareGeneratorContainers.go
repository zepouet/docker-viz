package flare

import (
	"github.com/samalba/dockerclient"
	"log"
	"strconv"
)

type(
	FlareContainers interface {
		BubbleContainers() string;
	}
)

// Load all containers information in Docker
func GenerateDockerContainerList() map[string]dockerclient.Container {
	docker := DockerEngineConnection()

	// load containers list
	listContainers, err := docker.ListContainers(true, true, "")
	if err != nil {
		log.Fatal(err)
	}

	containers := make(map[string]dockerclient.Container)
	for _, c := range listContainers {
		containers[c.Id] = c
	}

	return containers
}

// Create the json architecture
func MakeJsonContainers(dockerContainerList map[string]dockerclient.Container) string {
	var flare string
	nbFils := len(dockerContainerList)
	var i int = 0
	for _, container := range dockerContainerList {
		i++
		virtualSize := strconv.Itoa(int(container.SizeRw))
		flare += "{\"name\": \"" + container.Names[0] + "\", \"size\": " + virtualSize + "}"
		if i < nbFils {
			flare += ", "
		}
	}

	return flare
}

// Returns the full json for docker containers bubble diagram
func BubbleContainers() string {
	dockerContainerList := GenerateDockerContainerList()

	return  "{\"name\": \"Docker\", \"children\": [" + MakeJsonContainers(dockerContainerList) + "]}"
}