package apis

import (
	"database/sql"
	"fmt"
	"github.com/garylailai/cateraway-api/handlers/api/database"
	"github.com/garylailai/cateraway-api/handlers/api/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

//查询展示首页食物列表数据
func FoodAllQuery(c *gin.Context){
    db :=database.DbManager()
    item :=&models.Item{}
	order :=&models.Order{}
	bussiness_info :=&models.Bussiness_info{}
	category :=&models.Category{}
	item_category :=&models.Item_category{}
	price_type :=&models.Price_type{}
	user :=&models.User{}
	user_info :=&models.User_info{}
	meal_selection_mode :=&models.Meal_selection_mode{}
	set_meal :=&models.Set_meal{}
	set_meal_type :=&models.Set_meal_type{}
	if err := db.Find(item).Error; err !=nil {
		//迁移模式加载数据库表
		db.AutoMigrate(item)
		db.AutoMigrate(order)
		db.AutoMigrate(bussiness_info)
		db.AutoMigrate(category)
		db.AutoMigrate(item_category)
		db.AutoMigrate(price_type)
		db.AutoMigrate(user)
		db.AutoMigrate(user_info)
		db.AutoMigrate(meal_selection_mode)
		db.AutoMigrate(set_meal)
		db.AutoMigrate(set_meal_type)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到的都表已加载到数据库",
		})
	}


	//查询套餐类型
	//var set_meal_types models.Set_meal_type
	var rows *sql.Rows
	var set_meal_types []*models.Set_meal_type
	//db.Find(&set_meal_types)
	rows, _ = db.Raw(" SELECT set_meal_types.`id`,set_meals.`id`,set_meals.`name`,fit_number,price,set_meals.describe,picture1,picture2,picture3 FROM set_meals,set_meal_types WHERE set_meal_types.id=set_meals.`set_meal_type_id`").Rows()
	//rows, _ = db.Raw("SELECT set_meal_types.`id`,set_meal_types.name,picture,set_meals.`id`,set_meals.`name`,fit_number,price,set_meals.describe,picture1,picture2,picture3 FROM set_meals INNER JOIN set_meal_types ON set_meal_types.id=set_meals.`set_meal_type_id` WHERE set_meal_types.id=1").Rows()

	defer rows.Close()
	for rows.Next() {
		var set_meal_type models.Set_meal_type
		db.ScanRows(rows, &set_meal_type)
		var set_meals []models.Set_meal
		db.Find(&set_meals)
		var csli []models.Item_category



		csli = append(csli, models.Item_category{})
		set_meal_type.Set_meal=append(set_meals,models.Set_meal{Item_category:csli})
		set_meal_types = append(set_meal_types, &set_meal_type)

	}

	result := set_meal_types


	fmt.Println("Query_All输出:", result)

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
	})
}


