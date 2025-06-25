package bd

import (
	// "database/sql"
	"database/sql"
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
		query += "User_FirstName = '" + UField.UserFirstName + "'"
	}

	if len(UField.UserLastName) > 0 {
		query += coma + " User_LastName = '" + UField.UserLastName + "'"
	}

	query += ", User_DateUpg = '" + tools.FechaMySQL() + "' WHERE User_UUID = '" + User + "'"

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update User > Ejecución Exitosa")
	return nil

}

func SelectUser(UserId string)(models.User, error) {
	fmt.Println("Comienza Select User")
	User := models.User{}

	err := DbConnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	query := "SELECT * FROM users WHERE User_UUID = '"+ UserId + "'"

	var rows *sql.Rows
	rows, err = Db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return User, err
	}
	defer rows.Close()
	
	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime

	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpd = dateUpg.Time.String()

	fmt.Println("Select User > Ejecución Exitosa")
	return User, nil
}