package flare

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
	"github.com/Treeptik/docker-viz/dockertype"
)

func TestMakeJsonFatherWithImage(t *testing.T) {
	images := dockertype.GenerateDockerImageList()
	assert.Equal(t, len(images), 48)

	childs := dockertype.GenerateDockerChild(images)
	assert.Equal(t, len(childs), 12)
	assert.Equal(t, len(childs["Docker"]), 1)

	json := MakeJsonFather("Docker", childs, images, "size")
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)
}

func TestMakeJsonFatherWithContainer(t *testing.T) {
	container := dockertype.GenerateDockerContainerList()
	assert.Equal(t, len(container), 13)

	childs := dockertype.GenerateDockerChild(container)
	json := MakeJsonFather("Docker", childs, container, "size")
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)
}