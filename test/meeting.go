package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	v1 "meeting/api/meeting/v1"
)

var meetingClient v1.MeetingClient
var conn *grpc.ClientConn

func main() {
	Init()
	TestCreateMeeting()
	conn.Close()
}

// Init 初始化 grpc 链接
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9300", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	meetingClient = v1.NewMeetingClient(conn)
}
func TestCreateMeeting() {
	meeting := &v1.MeetingRequest_Meeting{
		Name:      "NewRrandMeet2",
		Address:   "祥原路",
		AppDeatil: "高峰论坛",
	}
	rsp, err := meetingClient.Create(context.Background(), &v1.MeetingRequest{
		Meeting: meeting,
	})
	if err != nil {
		panic("grpc 创建失败" + err.Error())
	}
	fmt.Println(rsp)
}
