package flare

import "github.com/Treeptik/docker-viz/dockertype"

// return all json for Dendrogram graph
func DendrogamFlare() string {
	dockerList := dockertype.GenerateDockerImageList()
	dockerImagesChilds := dockertype.GenerateDockerChild(dockerList)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerList, "size") + EndJson
}