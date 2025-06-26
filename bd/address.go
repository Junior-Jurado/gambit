package bd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Junior_Jurado/gambit/models"
	_ "github.com/go-sql-driver/mysql"
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

func AddressExist(User string, id int) (error, bool) {
	fmt.Println("Comienza AddressExist")
	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	query := "SELECT 1 FROM addresses WHERE Add_Id = " + strconv.Itoa(id) + " AND Add_UserId = '" + User + "'"
	fmt.Println(query)

	rows, err := Db.Query(query)
	if err != nil {
		return err, false
	}
	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("AddressExist > Ejecución Exitosa - valor devuelto " + valor)

	if valor == "1" {
		return nil, true
	}
	return nil, false

}

func UpdateAddress(t models.Address, User string) error {
	fmt.Println("Comienza Update Address")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	
	query := "UPDATE addresses SET "
	if t.AddAddress != "" {
		query += "Add_Address = '" + t.AddAddress + "', "
	}

	if t.AddCity != "" {
		query += "Add_City = '" + t.AddCity + "', "
	}

	if t.AddName != "" {
		query += "Add_Name = '" + t.AddName + "', "
	}

	if t.AddPostalCode != "" {
		query += "Add_PostalCode = '" + t.AddPostalCode + "', "
	}

	if t.AddPhone != "" {
		query += "Add_Phone = '" + t.AddPhone + "', "
	}

	if t.AddState != "" {
		query += "Add_State = '" + t.AddState + "', "
	}

	if t.AddTitle != "" {
		query += "Add_Title = '" + t.AddTitle + "', "
	}

	query, _ = strings.CutSuffix(query, ", ")
	query += " WHERE Add_Id = " + strconv.Itoa(t.AddId)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(query)
	fmt.Println("Update Address > Ejecución Exitosa")
	return nil
}

func DeleteAddress(id int) error {
	fmt.Println("Comienza Delete Address")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "DELETE FROM addresses WHERE Add_addId = " + strconv.Itoa(id)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	
	fmt.Println("Delete Address > Ejecución Exitosa")
	return nil
}