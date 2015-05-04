package flare

import "github.com/Treeptik/docker-viz/dockertype"

func BubbleFlare(who string) string {
	dockerList := make(map[string]dockertype.DockerType)

	if who == "images" {
		dockerList = GenerateDockerImageList()
	} else {
		dockerList = GenerateDockerContainerList()
	}

	dockerImagesChilds := GenerateDockerChild(dockerList)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerList) + EndJson
}