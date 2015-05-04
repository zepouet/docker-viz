package flare

// Returns the full json for docker images Dendrogam & bubble diagram
func DendrogamAndBubbleImages() string {
	dockerImagesList := GenerateDockerImageList()
	dockerImagesChilds := GenerateDockerChild(dockerImagesList)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerImagesList) + EndJson
}