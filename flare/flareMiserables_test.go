package flare

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
)

func TestMiserablesFlare(t *testing.T) {
	images := MiserablesFlare()
	assert.Equal(t, strings.Count(images, "["), strings.Count(images, "]"))
	assert.Equal(t, strings.Count(images, "{"), strings.Count(images, "}"))
	assert.Equal(t, strings.Count(images, "\\"), 0)
}