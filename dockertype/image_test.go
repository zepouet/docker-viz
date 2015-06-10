package dockertype

import (
"testing"
"github.com/stretchr/testify/assert"
"github.com/fsouza/go-dockerclient"
)

func TestContainerStruct(t *testing.T) {
	APIImg := docker.APIImages{}
	APIImg.ID = "THIS_ID"
	APIImg.RepoTags = []string{"THIS_NAME", "THIS_OTHER_NAME"}
	APIImg.ParentID = "THIS_IMAGE"
	APIImg.VirtualSize = 100

	Img := Image{APIImg}

	assert.Equal(t, Img.GetId(), "THIS_ID")
	assert.Equal(t, Img.GetName(), "THIS_NAME")
	assert.Equal(t, Img.GetFatherId(), "THIS_IMAGE")
	assert.Equal(t, Img.GetSize(), "100")
	assert.Equal(t, Img.GetLink().Size(), 0)
	assert.Equal(t, Img.GetVolumeFrom().Size(), 0)
}