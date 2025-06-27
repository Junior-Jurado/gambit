package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Junior_Jurado/gambit/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza Registro Orders")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	query := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES ('"
	query += o.Order_UserUUID + "', " + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + ", " + strconv.Itoa(o.Order_AddId) + ")"

	var result sql.Result
	result, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	for _, od := range o.OrdersDetails{
		query = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES("
		query += strconv.Itoa(int(LastInsertId)) + ", " + strconv.Itoa(od.OD_ProdId) + ", " + strconv.Itoa(od.OD_Quantity) + ", "
		query += strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ")"

		fmt.Println(query)
		_, err = Db.Exec(query)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("Insert Order > Ejecución Exitosa")
	return LastInsertId, nil
}
