package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/icecrime/api/api"
)

// New creates a new Docker remote API service instance in the form of an
// http.Handler.
func New(impl api.Service, swaggerFilePath string) http.Handler {
	baseSrv := newBaseServer(impl)
	containersSrv := newContainersServer(impl)

	container := restful.NewContainer()
	container.Add(baseSrv.WebService)
	container.Add(containersSrv.WebService)

	swaggerConf := swagger.Config{
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/docs/apidocs.json",
		SwaggerPath:     "/docs/swagger/",
		SwaggerFilePath: swaggerFilePath,
	}
	swagger.RegisterSwaggerService(swaggerConf, container)

	return container
}
