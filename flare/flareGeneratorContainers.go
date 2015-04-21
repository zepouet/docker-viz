package flare

import (
	"github.com/samalba/dockerclient"
	"log"
)

type(
	Flare interface {
		BubbleContainers() string;
	}
)

// Load all image information clone and commit in Docker
func GenerateDockerContainerList(dockerClient *string) map[string]dockerclient.Container {
	docker, _ := dockerclient.NewDockerClient(*dockerClient, nil)

	listContainers, err := docker.ListContainers(true, true, "")
	if err != nil {
		log.Fatal(err)
	}

	containers := make(map[string]dockerclient.Container)
	for _, c := range listContainers {
		containers[c.Id] = *c
	}

	return containers
}

// Returns the full json for docker containers bubble diagram
func BubbleContainers(dockerClient *string) string {
	return false
}