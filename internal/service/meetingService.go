package service

import (
	"context"
	v1 "meeting/api/meeting/v1"
	"meeting/internal/biz"
)

// GreeterService is a greeter service.
type MeetingService struct {
	v1.UnimplementedGreeterServer

	uc *biz.MeetingGreeterUsecase
}

// NewGreeterService new a greeter service.
func NewMeetingService(uc *biz.MeetingGreeterUsecase) *MeetingService {
	return &MeetingService{uc: uc}
}

func (meetingService *MeetingService) Create(ctx context.Context, req *v1.MeetingRequest) (*v1.MeetingReploy, error) {
	return meetingService.uc.Create(ctx, req)
}
