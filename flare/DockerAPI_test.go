package flare

import (
	"testing"
	"log"
)

func TestDockerEngineConnection(t *testing.T) {
	docker := DockerEngineConnection()

	_, err := docker.ListImages()

	if err != nil {
		log.Fatal(err)
	}
}