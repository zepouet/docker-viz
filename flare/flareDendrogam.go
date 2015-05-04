package flare

func DendrogamFlare() string {
	dockerList := GenerateDockerImageList()
	dockerImagesChilds := GenerateDockerChild(dockerList)

	return  BeginJson + MakeJsonFather("Docker", dockerImagesChilds, dockerList) + EndJson
}