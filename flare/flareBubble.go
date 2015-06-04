package flare

import (
	"github.com/Treeptik/docker-viz/dockertype"
	"log"
	"time"
)

// return all json for Bubble diagram
func BubbleFlare(who string) string {
	dockerList := make(map[string]dockertype.DockerType)

	start := time.Now()
	switch who {
		case "images":
			dockerList = dockertype.GenerateDockerImageList()
		case "containers":
			dockerList = dockertype.GenerateDockerContainerList()
		default:
	}
	elapsed := time.Since(start)
	log.Printf("GenerateDockerContainerList took %s", elapsed)

	start = time.Now()
	dockerImagesChilds := dockertype.GenerateDockerChild(dockerList)
	elapsed = time.Since(start)
	log.Printf("GenerateDockerChild took %s", elapsed)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerList) + EndJson
}