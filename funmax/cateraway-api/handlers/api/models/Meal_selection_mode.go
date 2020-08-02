package models
import "time"
//选餐模式表
type Meal_selection_mode struct {
	ID                           int              `gorm:"primary_key" json:"id"`      //id
	Meal_selection_mode          int              `json:"meal_selection_mode"`        //选餐模式(1自选;2半自选;3套餐)
	Is_deleted                   time.Time        `json:"is_deleted"`                 //软删除
	Create_ts                    time.Time        `json:"create_ts"`                  //创建时间
	Update_ts                    time.Time        `json:"update_ts"`                  //修改时间
}
