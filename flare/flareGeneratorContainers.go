package flare

// Returns the full json for docker containers bubble diagram
func BubbleContainers() string {
	dockerContainerList := GenerateDockerContainerList()
	dockerImagesChilds := GenerateDockerChild(dockerContainerList)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerContainerList) + EndJson
}