package flare

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

var docker string

func TestInit(t *testing.T) {
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

func TestGenerateDockerContainerAndJson(t *testing.T) {
	containers := GenerateDockerContainerList()
	assert.Equal(t, len(containers), 13)

	json := MakeJsonContainers(containers)
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)

	bubble := BubbleContainers()
	assert.Equal(t, len(bubble), len(json)+34)
	assert.Equal(t, strings.Count(bubble, "["), strings.Count(bubble, "]"))
	assert.Equal(t, strings.Count(bubble, "{"), strings.Count(bubble, "}"))
	assert.Equal(t, strings.Count(bubble, "\\"), 0)
}