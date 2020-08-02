package models

import "time"
//单品id与分类id
type Item_category struct {
	ID            int              `gorm:"primary_key" json:"id"`      //id
	Category      int              `json:"category"`                   //菜单类别
	Item          int              `json:"item"`                       //单品
	Is_deleted    time.Time        `json:"is_deleted"`                 //软删除
	Create_ts     time.Time        `json:"create_ts"`                  //创建时间
	Update_ts     time.Time        `json:"update_ts"`                  //修改时间
}

