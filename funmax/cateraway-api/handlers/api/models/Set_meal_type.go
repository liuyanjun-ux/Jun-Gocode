package models
import "time"
//套餐类型
type Set_meal_type struct {
	ID            int              `gorm:"primary_key" json:"id"`      //id
	Name          string           `json:"category"`                   //套餐类型名称
	Picture       string           `json:"picture"`                    //图片路径
	Set_meal      []Set_meal       `json:"set_meal"`                   //Set_meal表中的数据

	Is_deleted    time.Time        `json:"is_deleted"`                 //软删除
	Create_ts     time.Time        `json:"create_ts"`                  //创建时间
	Update_ts     time.Time        `json:"update_ts"`                  //修改时间

}

