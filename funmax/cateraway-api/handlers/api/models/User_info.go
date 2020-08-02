package models
import "time"

//用户初始化表
type User_info struct {
	ID                          int              `gorm:"primary_key" json:"id"`                //id
	User_id                     int              `json:"user_id"`                              //用户Id
	Nickname                    string           `json:"nickname"`                             //昵称
	Tel                         string           `json:"tel"`                                  //电话
	Email                       string           `json:"email"`                                //邮箱
	Address                     string           `json:"address"`                              //地址
	Gender                      string           `json:"gender"`                               //性别
	Is_deleted                  time.Time        `json:"is_deleted"`                           //软删除
	Create_ts                   time.Time        `json:"create_ts"`                            //创建时间
	Update_ts                   time.Time        `json:"update_ts"`                            //修改时间
}
