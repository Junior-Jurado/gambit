package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/Junior_Jurado/gambit/models"
	"github.com/Junior_Jurado/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Comienza Registro InsertProduct")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	columns := []string{"Prod_Title"}
	values := []string{"'" + tools.EscapeString(p.ProdTitle) + "'"}

	if len(p.ProdDescription) > 0 {
		columns = append(columns, "Prod_Description")
		values = append(values, "'"+tools.EscapeString(p.ProdDescription)+"'")
	}
	if p.ProdPrice > 0 {
		columns = append(columns, "Prod_Price")
		values = append(values, strconv.FormatFloat(p.ProdPrice, 'f', -1, 64))
	}
	if p.ProdCategId > 0 {
		columns = append(columns, "Prod_CategoryId")
		values = append(values, strconv.Itoa(p.ProdCategId))
	}
	if p.ProdStock > 0 {
		columns = append(columns, "Prod_Stock")
		values = append(values, strconv.Itoa(p.ProdStock))
	}
	if len(p.ProdPath) > 0 {
		columns = append(columns, "Prod_Path")
		values = append(values, "'"+tools.EscapeString(p.ProdCategPath)+"'")
	}

	query := fmt.Sprintf("INSERT INTO productos (%s) VALUES (%s)",
		strings.Join(columns, ", "),
		strings.Join(values, ", "),
	)
	fmt.Println(query)

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

	fmt.Println("Insert Product > Ejecución Exitosa")

	return LastInsertId, nil

}
