package flare

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
	"os/exec"
)

func init() {
	_, err := exec.Command("./dockerInit.sh").Output()
	if err != nil {
		log.Fatalf("Init: %s", err)
	}

}

func TestDockerEngineConnection(t *testing.T) {
	docker := DockerEngineConnection()

	_, err := docker.ListImages()

	if err != nil {
		log.Fatal(err)
	}
}

func TestLoader(t *testing.T) {
	images := LoadDockerImages()
	assert.Equal(t, len(images), 48)

	containers := LoadDockerContainers()
	assert.Equal(t, len(containers), 13)
}