package flare

import (
	"github.com/samalba/dockerclient"
	"log"
)

type(
	FlareContainers interface {
		BubbleContainers() string;
	}
)

// Load all containers information in Docker
func GenerateDockerContainerList(dockerClient *string) map[string]dockerclient.Container {
	docker, _ := dockerclient.NewDockerClient(*dockerClient, nil)

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