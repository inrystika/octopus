package registry

type ArtifactRegistry interface {
	CreateProject(projectReq *ProjectReq) error
	DeleteArtifact(projectName string, repositoryName string, reference string) error
}
