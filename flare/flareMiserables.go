package flare

import (
	"github.com/Treeptik/docker-viz/dockertype"
	"strconv"
)

func MakeMatriceJson(dockerList map[string]dockertype.DockerType) (string, map[string]int) {
	var flare string
	var i int = 0
	pos := make(map[string]int)

	for _, dock := range dockerList {
		if i != 0 {
			flare += ", "
		}
		flare += "{\"name\":\""+ dock.GetName() +"\"}"
		pos[dock.GetId()] = i
		i++
	}

	return flare, pos
}

func MakeLinkJson(dockerList map[string]dockertype.DockerType, pos map[string]int) string {
	var flare string
	comma := 0
	for _, dock := range dockerList {
		links := dock.GetLink()
		for _, docklink := range links.Values() {
			if comma != 0 {
				flare += ", "
			}
			flare += "{\"source\":" + strconv.Itoa(pos[dock.GetId()]) +",\"target\":" + strconv.Itoa(pos[docklink.(string)]) +",\"value\":2},"
			flare += "{\"source\":" + strconv.Itoa(pos[docklink.(string)]) +",\"target\":" + strconv.Itoa(pos[dock.GetId()]) +",\"value\":200}"
			comma++
		}
	}
	return flare
}

func MiserablesFlare() string {
	dockerList := dockertype.GenerateDockerContainerList()

	matrice, pos := MakeMatriceJson(dockerList)
	return  "{\"nodes\":[" + matrice + "],\"links\":[" + MakeLinkJson(dockerList, pos) + "]}"
}