package flare

import "github.com/Treeptik/docker-viz/dockertype"

func MakeMatriceJson(dockerList map[string]dockertype.DockerType) string {
	var flare string
	nbdock := len(dockerList)
	var i int = 0
	flare += "{\"nodes\":["
	for _, dock := range dockerList {
		i++
		flare += "{\"name\":\""+ dock.GetName() +"\"}"
		if i < nbdock {
			flare += ", "
		}
	}
	flare += "],\"links\":["
	flare += "]}"

	return flare
}

func MiserablesFlare() string {
	dockerList := GenerateDockerContainerList()

	return  MakeMatriceJson(dockerList)
}