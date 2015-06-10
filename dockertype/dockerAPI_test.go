package dockertype

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestDockerEngineConnection(t *testing.T) {
	dockerClient, err := DockerEngineConnection()

	if err != nil {
		log.Fatal(err)
	}

	_, err = dockerClient.Version()

	if err != nil {
		log.Fatal(err)
	}
}

func TestLoader(t *testing.T) {
	images := LoadDockerImages()
	assert.Equal(t, len(images), 48)

	containers := LoadDockerContainers()
	assert.Equal(t, len(containers), 13)
}