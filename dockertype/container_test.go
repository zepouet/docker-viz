package dockertype

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/fsouza/go-dockerclient"
)

func TestImagesStruct(t *testing.T) {
	APICont := docker.APIContainers{}
	APICont.ID = "THIS_ID"
	APICont.Names = []string{"THIS_NAME", "THIS_OTHER_NAME"}
	APICont.Image = "THIS_IMAGE"
	APICont.SizeRw = 100

	Cont := Container{APICont}

	assert.Equal(t, Cont.GetId(), "THIS_ID")
	assert.Equal(t, Cont.GetName(), "THIS_NAME")
	assert.Equal(t, Cont.GetFatherId(), "THIS_IMAGE")
	assert.Equal(t, Cont.GetSize(), "100")
}