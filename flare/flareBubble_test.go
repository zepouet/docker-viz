package flare

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
)

func TestBubbleFlare(t *testing.T) {
	images := BubbleFlare("images", "size")
	assert.Equal(t, strings.Count(images, "["), strings.Count(images, "]"))
	assert.Equal(t, strings.Count(images, "{"), strings.Count(images, "}"))
	assert.Equal(t, strings.Count(images, "\\"), 0)

	containers := BubbleFlare("containers", "size")
	assert.Equal(t, strings.Count(containers, "["), strings.Count(containers, "]"))
	assert.Equal(t, strings.Count(containers, "{"), strings.Count(containers, "}"))
	assert.Equal(t, strings.Count(containers, "\\"), 0)
}