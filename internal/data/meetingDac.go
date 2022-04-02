package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "meeting/api/meeting/v1"
	"meeting/internal/biz"
	"time"
)

type Meeting struct {
	Id        int64  `gorm:"primarykey"`
	Name      string `gorm:"column:name"`
	Address   string `gorm:"column:Address"`
	state     string `gorm:"column:state"`
	appdetail string `gorm:"column:appdetail"`
}

type meetingRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewMeetingRepo(data *Data, logger log.Logger) biz.MeetingRepo {
	return &meetingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (dac *meetingRepo) Create(ctx context.Context, req *v1.MeetingRequest) (*v1.MeetingReploy, error) {
	m := Meeting{
		Name:      req.Meeting.Name,
		Address:   req.Meeting.Address,
		state:     "1",
		appdetail: "备注",
	}
	var res v1.MeetingReploy

	if err := dac.data.db.Save(m).Error; err != nil {
		res.Success = "false"
		res.Msg = "会议信息创建失败！"
		res.SaveRecode = 0
		return &res, errors.New(500, "Meet_Create", err.Error())
	} else {
		go func(meeting Meeting) {
			if me, err := json.Marshal(meeting); err != nil {
				log.Info(err)
				return
			} else {
				dac.data.rdb.Set(req.Meeting.Name, string(me), 1*time.Second).Err()
			}
		}(m)
		return &res, err
	}
}
