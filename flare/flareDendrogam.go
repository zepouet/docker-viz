package flare

import "github.com/Treeptik/docker-viz/dockertype"

func DendrogamFlare() string {
	dockerList := dockertype.GenerateDockerImageList()
	dockerImagesChilds := dockertype.GenerateDockerChild(dockerList)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerList) + EndJson
}