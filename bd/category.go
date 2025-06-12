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
		query += "Categ_Name = '"+ tools.EscapeString(c.CategName)+ "'"
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(query, "SET") {
			query += ", "
		}
		query += "Categ_Path = '" + tools.EscapeString(c.CategPath)+"' "
	}

	query += "WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución Exitosa")
	return nil
}

func DeleteCategory(id int) error {
	fmt.Println("Comienza Registro de DeleteCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := fmt.Sprintf("DELETE FROM category WHERE Categ_Id = %s", strconv.Itoa(id))

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Ejecución Exitosa")
	return nil
}

func SelectCategories(CategId int, Slug string) ([]models.Category, error){
	fmt.Println("Comienza SelectCategories")

	var Categ []models.Category

	err := DbConnect()
	if err != nil {
		return Categ, err
	}
	defer Db.Close()

	query := fmt.Sprintf("SELECT Categ_Id, Categ_Name, Categ_Path FROM category ")

	if CategId > 0 {
		query += fmt.Sprintf("WHERE Categ_Id = %s", strconv.Itoa(CategId))
	} else {
		if len(Slug)>0 {
			query += fmt.Sprintf("WHERE Categ_Path LIKE '%'%s''%")
		}
	}

	fmt.Println(query)

	var rows *sql.Rows
	rows, err = Db.Query(query)

	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)
		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String
		Categ = append(Categ, c)
	}

	fmt.Println("Select Category > Ejecución Exitosa")
	return Categ, nil
}