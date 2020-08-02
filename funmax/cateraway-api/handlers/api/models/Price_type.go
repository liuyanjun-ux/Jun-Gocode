package models
import "time"
//定价分类表
type Price_type struct {
	ID            int              `gorm:"primary_key" json:"id"`      //id
	Price_type    int              `json:"price_type"`                 //套餐类型(1是按人数计算;2是按套餐计算)
	Is_deleted    time.Time        `json:"is_deleted"`                 //软删除
	Create_ts     time.Time        `json:"create_ts"`                  //创建时间
	Update_ts     time.Time        `json:"update_ts"`                  //修改时间
}
