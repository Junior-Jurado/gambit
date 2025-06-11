package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	// "strconv"
	// "strings"
	"github.com/Junior_Jurado/gambit/models"
	"github.com/Junior_Jurado/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/Junior_Jurado/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza Registro de InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	query := fmt.Sprintf("INSERT INTO category (Categ_Name, Categ_Path) VALUES ('%s', '%s')", c.CategName, c.CategPath)
	
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

	fmt.Println("Insert Category > Ejecución Exitosa")
	return LastInsertId, err2
}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza Registro de UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "UPDATE category SET "

	if len(c.CategName) > 0 {
		query += "CategName = '"+ tools.EscapeString(c.CategName)+ "'"
	}

	if len(c.CategName) > 0 {
		if !strings.HasSuffix(query, "SET") {
			query += ", "
		}
		query += "Categ_Path = '" + tools.EscapeString(c.CategPath)+"' "
	}

	query += "WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución Exitosa")
	return nil
}