package service

import (
	"github.com/dangerousvasil/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"lib-transport/proto"
	"lib-transport/swagger"

	"golang.org/x/exp/slog"
	"io/fs"
	"net/http"
)

const swaggerURL = "/api/v1/swagger/"
const protocURL = "/api/v1/proto/"

func mountSwagger(mux *runtime.ServeMux) error {
	slog.Info("mountSwagger")
	static, _ := fs.Sub(swaggerui.Swagfs, "embed")
	staticServerSwaggerUI := http.FileServer(http.FS(static))
	err := fs.WalkDir(swagger.Assets(), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		file, err := fs.ReadFile(swagger.Assets(), path)
		if err != nil {
			return err
		}

		slog.Info("Find swagger", "path", swaggerURL+path+"/swagger_spec")
		err = mux.HandlePath("GET", swaggerURL+path+"/**", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			req.URL.Path = req.URL.Path[len(swaggerURL+path):]
			staticServerSwaggerUI.ServeHTTP(w, req)
		})
		if err != nil {
			return err
		}
		err = mux.HandlePath("GET", swaggerURL+path+"/swagger_spec", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			_, err = w.Write(file)
		})
		if err != nil {
			return err
		}
		return nil
	})
	slog.Info("mountSwagger DONE", "err", err)
	if err != nil {
		return err
	}
	staticServerSwagger := http.FileServer(http.FS(swagger.Assets()))
	err = mux.HandlePath("GET", swaggerURL+"*", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		req.URL.Path = req.URL.Path[len(swaggerURL):]
		staticServerSwagger.ServeHTTP(w, req)
	})

	slog.Info("staticServerSwagger DONE", "err", err)
	return err
}

func mountProtoc(mux *runtime.ServeMux) error {
	staticServerProto := http.FileServer(http.FS(proto.Assets()))
	return mux.HandlePath("GET", protocURL+"**", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		req.URL.Path = req.URL.Path[len(protocURL):]
		staticServerProto.ServeHTTP(w, req)
	})
}
