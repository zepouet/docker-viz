package flare

import (
	"github.com/Treeptik/docker-viz/dockertype"
	"strconv"
)

// Make two colums name
func MakeMatriceJson(dockerList map[string]dockertype.DockerType) (string, map[string]int) {
	var flare string
	var i int = 0
	pos := make(map[string]int)

	for _, dock := range dockerList {
		// if not a first
		if i != 0 {
			flare += ", "
		}
		flare += "{\"name\":\""+ dock.GetName() +"\"}"
		pos[dock.GetId()] = i
		i++
	}

	return flare, pos
}

// Make link in matrice
func MakeLinkJson(dockerList map[string]dockertype.DockerType, pos map[string]int) string {
	var flare string
	comma := 0
	for _, dock := range dockerList {
		links := dock.GetLink()
		// for all link compared with "dock"
		for _, docklink := range links.Values() {
			// if not a first
			if comma != 0 {
				flare += ", "
			}
			flare += "{\"source\":" + strconv.Itoa(pos[dock.GetId()]) +",\"target\":" + strconv.Itoa(pos[docklink.(string)]) +",\"value\":2}"
			comma++
		}

		Volumelinks := dock.GetVolumeFrom()
		// for all volumelink compared with "dock"
		for _, dockVolumelinks := range Volumelinks.Values() {
			// if not a first
			if comma != 0 {
				flare += ", "
			}
			flare += "{\"source\":" + strconv.Itoa(pos[dockVolumelinks.(string)]) +",\"target\":" + strconv.Itoa(pos[dock.GetId()]) +",\"value\":200}"
			comma++
		}
	}
	return flare
}

// return the complete json for miserables matrice
func MiserablesFlare() string {
	dockerList := dockertype.GenerateDockerContainerList()

	matrice, pos := MakeMatriceJson(dockerList)
	return  "{\"nodes\":[" + matrice + "],\"links\":[" + MakeLinkJson(dockerList, pos) + "]}"
}