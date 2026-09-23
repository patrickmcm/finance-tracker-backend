package controllers

import (
	"google.golang.org/grpc"
)

type Controller interface {
	RegisterController(s *grpc.Server)
}
