package models
import "time"
//套餐表
type Set_meal struct {
	ID                 int              `gorm:"primary_key" json:"id"`      //id
	Name               string           `json:"name"`                       //套餐名称
	Fit_number         string           `json:"fit_number"`                 //适合人数
	Price              float64          `json:"price"`                      //套餐价格
	Describe           string           `json:"describe"`                   //套餐描述
	Picture1           string           `json:"picture"`                    //图片路径1
	Picture2           string           `json:"picture"`                    //图片路径2
	Picture3           string           `json:"picture"`                    //图片路径3
	Item_category      []Item_category  `json:"item_category"`              //单品分类表中的数据

	Is_deleted         time.Time        `json:"is_deleted"`                 //软删除
	Create_ts          time.Time        `json:"create_ts"`                  //创建时间
	Update_ts          time.Time        `json:"update_ts"`                  //修改时间
}
