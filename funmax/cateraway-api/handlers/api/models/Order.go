package models
import "time"
//订单表
type Order struct {
	ID                  int              `gorm:"primary_key" json:"id"`      //id
	Order_id            string           `json:"order_id"`                   //订单编号
	Name                string           `json:"name"`                       //联系人用户名称
	Tel                 string           `json:"tel"`                        //电话
	Address             string           `json:"address"`                    //地址
	Order_date          string           `json:"order_date"`                 //订单时间
	Delivery_date       string           `json:"delivery_date"`              //送货时间
	Order_price         float64          `json:"order_price"`                //订单价格
	Order_state         int              `json:"order_state"`                //订单状态(1是确认,2是取消)
	Payment_method      int              `json:"payment_method"`             //订单支付模式(1是信用卡;2是入数;3是PayMe)
	Payment_state       int              `json:"payment_state"`              //订单支付状态(1是确认;2是取消)
	Upload_data         string           `json:"upload_data"`                //上传资料
	Delivery_state      int              `json:"delivery_state"`             //送货状态(1是已送出;2是未送出)
	Printing            int              `json:"printing"`                   //订单打印状态(1是已打印;2是未打印)
	Information_state   int              `json:"information_state"`          //信息回访状态(1是确认,2是取消)
	Is_deleted          time.Time        `json:"is_deleted"`                 //软删除
	Create_ts           time.Time        `json:"create_ts"`                  //创建时间
	Update_ts           time.Time        `json:"update_ts"`                  //修改时间
}
