package merchandise

import (
	"cmk/generated/merchandise/design/v1/designv1connect"
	"cmk/merchandise/design"
	"net/http"
)

//encore:service
type Service struct {
	routes http.Handler
}

//encore:api public raw path=/merchandise.design.v1.DesignService/*endpoint
func (s *Service) DesignService(w http.ResponseWriter, r *http.Request) {
	s.routes.ServeHTTP(w, r)
}

func initService() (*Service, error) {
	design := &design.Server{}
	mux := http.NewServeMux()
	path, handler := designv1connect.NewDesignServiceHandler(design)
	mux.Handle(path, handler)
	return &Service{routes: mux}, nil
}
