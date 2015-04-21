package flare

import (
	"testing"

	"github.com/samalba/dockerclient"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"os"
	"log"
	"strings"
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

func TestGenerateDockerImageAndJson(t *testing.T) {
	images := GenerateDockerImageList(&dockerClient)
	assert.Equal(t, len(images), 48)

	childs := GenerateDockerImageChild(images)
	assert.Equal(t, len(childs), 12)
	assert.Equal(t, len(childs["Docker"]), 1)

	json := MakeJson("Docker", childs, images)
	assert.Equal(t, len(json), 2210)
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)

	dendrogam := DendrogamAndBubbleImages(&dockerClient)
	assert.Equal(t, len(dendrogam), len(json)+34)
	assert.Equal(t, strings.Count(dendrogam, "["), strings.Count(dendrogam, "]"))
	assert.Equal(t, strings.Count(dendrogam, "{"), strings.Count(dendrogam, "}"))
	assert.Equal(t, strings.Count(dendrogam, "\\"), 0)
}