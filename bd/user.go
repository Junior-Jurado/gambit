package bd

import (
	// "database/sql"
	"fmt"
	// "strconv"
	// "strings"
	// "time"
	// "errors"

	"github.com/Junior_Jurado/gambit/models"
	"github.com/Junior_Jurado/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateUser(UField models.User, User string) error {
	fmt.Println("Comienza UpdateUser")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "UPDATE users SET "
	coma := ""
	
	if len(UField.UserFirstName) > 0 {
		coma = ","
		query += "User_FirtsName = '" + UField.UserFirstName + "'"
	}

	if len(UField.UserLastName) > 0 {
		query += coma + " User_LastName = '" + UField.UserLastName + "'"
	}

	query += ", User_DataUpd = '" + tools.FechaMySQL() + "' WHERE User_UUID = '" + User + "'"

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update User > Ejecución Exitosa")
	return nil

}