package flare

import (
	"github.com/Treeptik/docker-viz/dockertype"
)

// return all json for Bubble diagram
func BubbleFlare(who string) string {
	dockerList := make(map[string]dockertype.DockerType)

	switch who {
		case "images":
			dockerList = dockertype.GenerateDockerImageList()
		case "containers":
			dockerList = dockertype.GenerateDockerContainerList()
		default:
	}

	dockerImagesChilds := dockertype.GenerateDockerChild(dockerList)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerList) + EndJson
}