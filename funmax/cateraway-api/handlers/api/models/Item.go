package models

import "time"

//单品表
type Item struct {
	ID            int              `gorm:"primary_key" json:"id"`      //id
	Name          string           `json:"name"`                       //名字
	Details       string           `json:"details"`                    //细节
	Price         float64          `json:"price"`                      //价格
	Remark        string           `json:"remark"`                     //评论
	Is_deleted    time.Time        `json:"is_deleted"`                 //软删除
	Create_ts     time.Time        `json:"create_ts"`                  //创建时间
	Update_ts     time.Time        `json:"update_ts"`                  //修改时间
}