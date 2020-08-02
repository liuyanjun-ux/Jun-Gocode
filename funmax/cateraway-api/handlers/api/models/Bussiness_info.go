package models

import "time"

//商业信息表
type Bussiness_info struct {
	ID                          int              `gorm:"primary_key" json:"id"`                //id
	User_id                     int              `json:"user_id"`                              //用户Id
	Type                        string           `json:"type"`                                 //类型
	Name                        string           `json:"name"`                                 //名称
	Tel                         string           `json:"tel"`                                  //电话
	Contact_person              string           `json:"contact_person"`                       //联系人
	Contact_tel                 string           `json:"contact_tel"`                          //联系电话
	Address                     string           `json:"address"`                              //地址
	Is_deleted                  time.Time        `json:"is_deleted"`                           //软删除
	Create_ts                   time.Time        `json:"create_ts"`                            //创建时间
	Update_ts                   time.Time        `json:"update_ts"`                            //修改时间
}