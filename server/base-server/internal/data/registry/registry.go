package registry

type ArtifactRegistry interface {
	CreateProject(projectReq *ProjectReq) error
}
