package flare

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
)

func TestDendrogamFlare(t *testing.T) {
	images := DendrogamFlare()
	assert.Equal(t, strings.Count(images, "["), strings.Count(images, "]"))
	assert.Equal(t, strings.Count(images, "{"), strings.Count(images, "}"))
	assert.Equal(t, strings.Count(images, "\\"), 0)
}