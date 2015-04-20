package flare

import (
	"testing"

	"github.com/samalba/dockerclient"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"os"
	"log"
)

var dockerClient string

func init() {
	dockerClient = os.Getenv("DOCKER_HOST")
	if dockerClient == "" {
		dockerClient = "unix:///var/run/docker.sock"
	}

	_, err := exec.Command("./dockerInit.sh").Output()
	if err != nil {
		log.Fatalf("Init: %s", err)
	}

}

func TestInitImagesAndContainers(t *testing.T) {

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

func TestGenerateDockerImageList(t *testing.T) {
	images := GenerateDockerImageList(&dockerClient)
	assert.Equal(t, len(images), 48)

	childs := GenerateDockerImageChild(images)
	assert.Equal(t, len(childs), 12)
	assert.Equal(t, len(childs["Docker"]), 1)
}