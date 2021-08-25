package registry

// https://github.com/goharbor/harbor/blob/release-1.4.0/docs/swagger.yaml#L2178
type ProjectReq struct {
	ProjectName string
	Public      bool
}
