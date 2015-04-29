package flare

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

var dockerClient string

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
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)

	dendrogam := DendrogamAndBubbleImages()
	assert.Equal(t, strings.Count(dendrogam, "["), strings.Count(dendrogam, "]"))
	assert.Equal(t, strings.Count(dendrogam, "{"), strings.Count(dendrogam, "}"))
	assert.Equal(t, strings.Count(dendrogam, "\\"), 0)
}