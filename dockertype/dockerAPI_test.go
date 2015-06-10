package dockertype

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
	"os"
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
	assert.Equal(t, CountDockerImages(), 48)
	assert.Equal(t, len(GenerateDockerImageList()), 48)

	containers := LoadDockerContainers()
	assert.Equal(t, len(containers), 13)
	assert.Equal(t, CountDockerContainer(), 13)
	assert.Equal(t, len(GenerateDockerContainerList()), 13)

	// test if docker engine not found
	dockerClient := os.Getenv("DOCKER_HOST")
	os.Setenv("DOCKER_HOST", "DOCKER_FAIL")

	assert.Equal(t, CountDockerImages(), 0)
	assert.Equal(t, len(GenerateDockerImageList()), 0)

	assert.Equal(t, CountDockerContainer(), 0)
	assert.Equal(t, len(GenerateDockerContainerList()), 0)

	os.Setenv("DOCKER_HOST", dockerClient)
}

func TestDockerStatut(t *testing.T) {
	if(!DockerStatut()) {
		log.Fatal("docker not detected")
	}

	// test if docker engine not found
	dockerClient := os.Getenv("DOCKER_HOST")
	os.Setenv("DOCKER_HOST", "DOCKER_FAIL")

	if(DockerStatut()) {
		log.Fatal("docker found in DOCKER_FAIL ????")
	}

	os.Setenv("DOCKER_HOST", dockerClient)
}

func TestDockerVersion(t *testing.T) {
	assert.NotEqual(t, DockerVersion(), "Docker Engine not found")

	// test if docker engine not found
	dockerClient := os.Getenv("DOCKER_HOST")
	os.Setenv("DOCKER_HOST", "DOCKER_FAIL")

	assert.Equal(t, DockerVersion(), "Docker Engine not found")

	os.Setenv("DOCKER_HOST", dockerClient)
}