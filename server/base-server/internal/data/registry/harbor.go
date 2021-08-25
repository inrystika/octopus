package registry

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"server/base-server/internal/conf"
	"server/common/errors"
	"strconv"
	"strings"

	harborV1 "server/common/harbor/v1/swagger"
	harborV2 "server/common/harbor/v2/swagger"
	"server/common/log"
)

const (
	apiVersion string = "v2.0"
)

func NewRegistry(confData *conf.Data, logger log.Logger) ArtifactRegistry {
	return newHarborRegistry(confData.Harbor, logger)
}

func newHarborRegistry(config *conf.Harbor, logger log.Logger) *harborRegistry {
	log := log.NewHelper("Harbor", logger)
	registry := &harborRegistry{
		config:     config,
		apiVersion: config.ApiVersion,
		log:        log,
	}

	if config.ApiVersion == "" || config.ApiVersion == apiVersion {
		harborCfg := &harborV2.Configuration{
			BasePath: registry.GetBaseURL(),
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			},
		}
		registry.ArtifactRegistry = &v2Registry{
			client: harborV2.NewAPIClient(harborCfg),
			basicAuth: harborV2.BasicAuth{
				UserName: config.Username,
				Password: config.Password,
			},
			log: log,
		}
	} else {
		harborCfg := &harborV1.Configuration{
			BasePath: registry.GetBaseURL(),
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			},
		}
		registry.ArtifactRegistry = &v1Registry{
			client: harborV1.NewAPIClient(harborCfg),
			basicAuth: harborV1.BasicAuth{
				UserName: config.Username,
				Password: config.Password,
			},
			log: log,
		}
	}
	return registry
}

type harborRegistry struct {
	log        *log.Helper
	config     *conf.Harbor
	apiVersion string
	ArtifactRegistry
}

func (h *harborRegistry) GetBaseURL() string {
	urlStr := h.config.Host
	if h.config.UseSSL {
		urlStr = "https://" + urlStr
	} else {
		urlStr = "http://" + urlStr
	}
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	var err error
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	var baseURLStr string

	if h.apiVersion == "" || h.apiVersion == apiVersion {
		baseURLStr = baseURL.String() + "api/" + "v2.0"
	} else {
		baseURLStr = baseURL.String() + "api/"
	}
	return baseURLStr
}

type v1Registry struct {
	client    *harborV1.APIClient
	basicAuth harborV1.BasicAuth
	log       *log.Helper
}

type v2Registry struct {
	client    *harborV2.APIClient
	basicAuth harborV2.BasicAuth
	log       *log.Helper
}

func (h *v1Registry) CreateProject(projectReq *ProjectReq) error {
	var public int32
	if projectReq.Public {
		public = 1
	}
	req := harborV1.ProjectReq{
		ProjectName: projectReq.ProjectName,
		Public:      public,
	}
	response, err1 := h.client.ProductsApi.ProjectsPost(h.GetAuth(nil), req)
	StatusCode := response.StatusCode
	switch StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusConflict:
		return errors.Errorf(nil, errors.ErrorHarborProjectExists)
	default:
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		h.log.Error(context.Background(), string(body))
		return errors.Errorf(err1, errors.ErrorHarborCheckProjectFailed)
	}
}

func (h *v1Registry) GetAuth(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	ctx = context.WithValue(ctx, harborV1.ContextBasicAuth, h.basicAuth)
	return ctx
}

func (h *v2Registry) GetAuth(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	ctx = context.WithValue(ctx, harborV2.ContextBasicAuth, h.basicAuth)
	return ctx
}

func (h *v2Registry) CreateProject(projectReq *ProjectReq) error {
	req := harborV2.ProjectReq{
		ProjectName: projectReq.ProjectName,
		Metadata: &harborV2.ProjectMetadata{
			Public: strconv.FormatBool(projectReq.Public),
		},
	}
	response, err1 := h.client.ProjectApi.CreateProject(h.GetAuth(nil), req, nil)
	StatusCode := response.StatusCode
	switch StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusConflict:
		return errors.Errorf(nil, errors.ErrorHarborProjectExists)
	default:
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		h.log.Error(context.Background(), string(body))
		return errors.Errorf(err1, errors.ErrorHarborCheckProjectFailed)
	}
}
