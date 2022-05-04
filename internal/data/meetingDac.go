package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "meeting/api/meeting/v1"
	"meeting/internal/biz"
	"strconv"
	"time"
)

type Meeting struct {
	ID        int64     `gorm:"column:id;primarykey"`
	Name      string    `gorm:"column:name"`
	Address   string    `gorm:"column:Address"`
	State     string    `gorm:"column:state"`
	AppDeatil string    `gorm:"column:appdetail"`
	UserID    int64     `gorm:"column:user_id"`
	CreatTime time.Time `gorm:"column:createTime"`
}

type MeetingUser struct {
	ID         int64 `gorm:"column:id;primarykey"`
	ConsumerId int64 `gorm:"column:consumer_id"`
	MeetingId  int64 `gorm:"column:meeting_id"`
	State      int   `gorm:"column:state"`
}

func (Meeting) TableName() string {
	return "scholar_meeting"
}
func (MeetingUser) TableName() string {
	return "scholar_consumer_meeting"
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
		UserID:    ctx.Value("UserId").(int64),
		CreatTime: time.Now(),
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

	if err := dac.data.db.Save(&m).Error; err != nil {
		res.Success = "false"
		res.Msg = "会议信息创建失败！"
		res.SaveRecode = 0
		return &res, errors.New(500, "Meet_Create", err.Error())
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

	//var meet Meeting
	//if err2 := dac.data.db.Model(&meet).Where("name=?", req.Meeting.Name ).Updates(map[string]interface{}{"name": "小样"}).Error; err2 != nil {
	//	res.Success = "false"
	//	res.Msg = "会议信息创建失败！"
	//	res.SaveRecode = 0
	//	return &res, errors.New(500, "Meet_Create", err2.Error())
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
	return &res, nil
}

func (dac *meetingRepo) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReploy, error) {

	var meets []Meeting
	dac.data.db.Where("name=?", req.Meeting.Name).Find(&meets)

	c, _ := strconv.Atoi(dac.data.rdb.Get(req.Meeting.Name).Val())

	var res v1.RegisterReploy
	if c > 25 {
		res.Success = "false"
		res.Msg = "超过与会人数"
		return &res, nil
	}
	c++
	dac.data.rdb.Set(req.Meeting.Name, c, 0)
	var meet Meeting
	dac.data.db.Where("name=?", req.Meeting.Name).Find(&meet)

	var meetUser = MeetingUser{
		MeetingId:  meet.ID,
		ConsumerId: ctx.Value("UserId").(int64),
		State:      1,
	}
	dac.data.db.Save(meetUser)

	res.Success = "true"
	res.Msg = "注册成功"
	return &res, nil
}
