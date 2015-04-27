package flare

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"os/exec"
	"log"
	"strings"
)

var dockerClient string

func init() {
	_, err := exec.Command("./dockerInit.sh").Output()
	if err != nil {
		log.Fatalf("Init: %s", err)
	}

}

func TestInitImagesAndContainers(t *testing.T) {

	docker := DockerEngineConnection()

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
	images := GenerateDockerImageList()
	assert.Equal(t, len(images), 48)

	childs := GenerateDockerImageChild(images)
	assert.Equal(t, len(childs), 12)
	assert.Equal(t, len(childs["Docker"]), 1)

	json := MakeJsonImages("Docker", childs, images)
	assert.Equal(t, len(json), 2210)
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)

	dendrogam := DendrogamAndBubbleImages()
	assert.Equal(t, len(dendrogam), len(json)+34)
	assert.Equal(t, strings.Count(dendrogam, "["), strings.Count(dendrogam, "]"))
	assert.Equal(t, strings.Count(dendrogam, "{"), strings.Count(dendrogam, "}"))
	assert.Equal(t, strings.Count(dendrogam, "\\"), 0)
}