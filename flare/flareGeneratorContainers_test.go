package flare

import (
	"testing"

	"github.com/samalba/dockerclient"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"os"
	"log"
)

var docker string

func init() {
	docker = os.Getenv("DOCKER_HOST")
	if docker == "" {
		docker = "unix:///var/run/docker.sock"
	}

	_, err := exec.Command("./dockerInit.sh").Output()
	if err != nil {
		log.Fatalf("Init: %s", err)
	}

}

func TestInit(t *testing.T) {

	docker, err := dockerclient.NewDockerClient(docker, nil)
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

func TestGenerateDockerContainerAndJson(t *testing.T) {
	containers := GenerateDockerContainerList(&docker)
	assert.Equal(t, len(containers), 13)
}