package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "gcAdmin:gcadminpasswd@tcp(47.102.147.41:3306)/gc?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Open successfully!")
	}
	defer db.Close()
}
