// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type MeetingHTTPServer interface {
	Create(context.Context, *MeetingRequest) (*MeetingReploy, error)
}

func RegisterMeetingHTTPServer(s *http.Server, srv MeetingHTTPServer) {
	r := s.Route("/")
	r.POST("/meeting/create", _Meeting_Create0_HTTP_Handler(srv))
}

func _Meeting_Create0_HTTP_Handler(srv MeetingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MeetingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.Meeting/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*MeetingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MeetingReploy)
		return ctx.Result(200, reply)
	}
}

type MeetingHTTPClient interface {
	Create(ctx context.Context, req *MeetingRequest, opts ...http.CallOption) (rsp *MeetingReploy, err error)
}

type MeetingHTTPClientImpl struct {
	cc *http.Client
}

func NewMeetingHTTPClient(client *http.Client) MeetingHTTPClient {
	return &MeetingHTTPClientImpl{client}
}

func (c *MeetingHTTPClientImpl) Create(ctx context.Context, in *MeetingRequest, opts ...http.CallOption) (*MeetingReploy, error) {
	var out MeetingReploy
	pattern := "/meeting/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/helloworld.v1.Meeting/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
