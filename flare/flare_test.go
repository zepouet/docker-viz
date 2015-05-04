package flare

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
)

func TestMakeJsonFatherWithImage(t *testing.T) {
	images := GenerateDockerImageList()
	assert.Equal(t, len(images), 48)

	childs := GenerateDockerChild(images)
	assert.Equal(t, len(childs), 12)
	assert.Equal(t, len(childs["Docker"]), 1)

	json := MakeJsonFather("Docker", childs, images)
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)
}

func TestMakeJsonFatherWithContainer(t *testing.T) {
	container := GenerateDockerContainerList()
	assert.Equal(t, len(container), 13)

	childs := GenerateDockerChild(container)
	json := MakeJsonFather("Docker", childs, container)
	assert.Equal(t, strings.Count(json, "["), strings.Count(json, "]"))
	assert.Equal(t, strings.Count(json, "{"), strings.Count(json, "}"))
	assert.Equal(t, strings.Count(json, "\\"), 0)
}