package main

import (
	"net"
	"net/http"

	local_grpc "github.com/costiss/go-rest-grpc/grpc"
	"github.com/costiss/go-rest-grpc/rest"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ginHandler *gin.Engine
}

func main() {

	lis, _ := net.Listen("tcp", ":8080")
	go func() {
		ginServer := rest.MakeRestServer()
		h := &Handler{
			ginHandler: ginServer,
		}
		http.Serve(lis, h)
	}()

	grpcServer := local_grpc.MakeGrpcServer()
	grpcServer.Serve(lis)

}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ginHandler.ServeHTTP(w, r)
}
