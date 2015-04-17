package flare

import (
	"github.com/samalba/dockerclient"
	"github.com/stretchr/testify/assert"
	"testing"
	"os/exec"
	"os"
	"log"
)

func init() {

	_, err := exec.Command("./dockerInit.sh").Output()
	if err != nil {
		log.Fatalf("Init: %s", err)
	}

}

func TestInitImagesAndContainers(t *testing.T) {
	dockerClient := os.Getenv("DOCKER_HOST")
	if dockerClient == "" {
		dockerClient = "unix:///var/run/docker.sock"
	}

	docker, err := dockerclient.NewDockerClient(dockerClient, nil)
	if err != nil {
		t.Fatal("Cannot init the docker client")
	}

	images, err := docker.ListImages()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(images), 48)

	containers, err := docker.ListContainers(true, true, "")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(containers), 13)
}