package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "meeting/api/meeting/v1"
	"meeting/internal/biz"
)

type Meeting struct {
	ID        int64  `gorm:"column:id;primarykey"`
	Name      string `gorm:"column:name"`
	Address   string `gorm:"column:Address"`
	State     string `gorm:"column:state"`
	AppDeatil string `gorm:"column:appdetail"`
}

func (Meeting) TableName() string {
	return "scholar_meeting"
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
	println("dac_Create")
	println(req)
	m := Meeting{
		Name:      req.Meeting.Name,
		Address:   req.Meeting.Address,
		State:     "1",
		AppDeatil: "备注",
	}
	resJson, _ := json.Marshal(req)
	println("2dac_Create" + string(resJson))
	var res v1.MeetingReploy
	//go func(meeting Meeting) {
	//	if me, err := json.Marshal(meeting); err != nil {
	//		log.Info(err)
	//		return
	//	} else {
	//		dac.data.rdb.Set(req.Meeting.Name, string(me), 1*time.Second).Err()
	//	}
	//}(m)
	//
	//res.Success = "true"
	//res.Msg = "会议信息创建成功！"
	//res.SaveRecode = 0

	//if err := dac.data.db.Save(&m).Error; err != nil {
	//	res.Success = "false"
	//	res.Msg = "会议信息创建失败！"
	//	res.SaveRecode = 0
	//	return &res, errors.New(500, "Meet_Create", err.Error())
	//} else {
	//	go func(meeting Meeting) {
	//		if me, err := json.Marshal(meeting); err != nil {
	//			log.Info(err)
	//			println(err.Error())
	//			return
	//		} else {
	//			println(string(me))
	//			err = dac.data.rdb.Set(meeting.Name, string(me), 0).Err()
	//			if err != nil {
	//				panic(err)
	//			}
	//		}
	//	}(m)
	//}
	var meet Meeting
	if err2 := dac.data.db.Model(&meet).Where("name=?", "NewRrandMeet").Updates(map[string]interface{}{"name": "小样"}).Error; err2 != nil {
		res.Success = "false"
		res.Msg = "会议信息创建失败！"
		res.SaveRecode = 0
		return &res, errors.New(500, "Meet_Create", err2.Error())
	} else {
		go func(meeting Meeting) {
			if me, err := json.Marshal(meeting); err != nil {
				log.Info(err)
				println(err.Error())
				return
			} else {
				println(string(me))
				err = dac.data.rdb.Set(meeting.Name, string(me), 0).Err()
				if err != nil {
					panic(err)
				}
			}
		}(m)
	}
	return &res, nil
}
