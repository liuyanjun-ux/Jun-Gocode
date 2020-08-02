package models

import "time"

//单品类别表
type Category struct {
	ID            int              `gorm:"primary_key" json:"id"`      //id
	Name          string           `json:"name"`                       //菜单类别名称
	Tags          string           `json:"tags"`                       //标签
	Remark        string           `json:"remark"`                     //评论
	Is_deleted    time.Time        `json:"is_deleted"`                 //软删除
	Create_ts     time.Time        `json:"create_ts"`                  //创建时间
	Update_ts     time.Time        `json:"update_ts"`                  //修改时间
}
