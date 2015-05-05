package flare

import "github.com/Treeptik/docker-viz/dockertype"

func MakeMatriceJson(dockerList map[string]dockertype.DockerType) string {
	var flare string
	nbdock := len(dockerList)
	var i int = 0

	for _, dock := range dockerList {
		i++
		flare += "{\"name\":\""+ dock.GetName() +"\"}"
		if i < nbdock {
			flare += ", "
		}
	}

	return flare
}

func MakeLinkJson(dockerList map[string]dockertype.DockerType) string {
	return ""
}

func MiserablesFlare() string {
	dockerList := dockertype.GenerateDockerContainerList()

	return  "{\"nodes\":[" + MakeMatriceJson(dockerList) + "],\"links\":[" + MakeLinkJson(dockerList) + "]}"
}