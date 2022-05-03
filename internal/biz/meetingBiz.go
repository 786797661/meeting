package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "meeting/api/meeting/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type MeetingRepo interface {
	Create(context.Context, *v1.MeetingRequest) (*v1.MeetingReploy, error)
	Register(context.Context, *v1.RegisterRequest) (*v1.RegisterReploy, error)
}

// GreeterUsecase is a Greeter usecase.
type MeetingGreeterUsecase struct {
	repo MeetingRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewMeetingGreeterUsecase(repo MeetingRepo, logger log.Logger) *MeetingGreeterUsecase {
	return &MeetingGreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *MeetingGreeterUsecase) Create(ctx context.Context, req *v1.MeetingRequest) (*v1.MeetingReploy, error) {
	//uc.log.WithContext(ctx).Infof("Create: %v", req.Meeting.Name)
	return uc.repo.Create(ctx, req)
}

func (uc *MeetingGreeterUsecase) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReploy, error) {
	//uc.log.WithContext(ctx).Infof("Create: %v", req.Meeting.Name)
	return uc.repo.Register(ctx, req)
}
