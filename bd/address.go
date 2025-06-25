package bd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Junior_Jurado/gambit/models"
)

func InsertAddress(addr models.Address, User string) error {
	fmt.Println("Comienza el Registro InsertAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name)"
	query += " VALUES ('"+ User + "', '" + addr.AddAddress + "', '" + addr.AddCity + "', '"+ addr.AddState + "', '" + addr.AddPostalCode + "', '" + addr.AddPhone + "', '" + addr.AddTitle + "', '" + addr.AddName + "')"

	fmt.Println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("INSERT Address > Ejecución Exitosa")
	return nil
}