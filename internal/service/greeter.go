package service

import (
	"context"

	v1 "point-server/api/helloworld/v1"
	"point-server/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// GetName implements helloworld.GetName
func (s *GreeterService) GetName(ctx context.Context, in *v1.GetRequest) (*v1.GetResponse, error) {
	// s.log.WithContext(ctx).Infof("GetName Received: %v", in.GetName())

	// if in.GetPathRepeatedFloatValue() == "error" {
	// 	return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	// }
	// GetPathRepeatedFloatValue
	return &v1.GetResponse{Coordinates: in.GetCoordinates()}, nil
}
