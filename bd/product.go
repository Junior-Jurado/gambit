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

	query := fmt.Sprintf("INSERT INTO products (%s) VALUES (%s)",
		strings.Join(columns, ", "),
		strings.Join(values, ", "),
	)

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

func UpdateProduct(p models.Product) error {
	fmt.Println("Comienza Update")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "UPDATE products SET "
	query = tools.ArmoSentencia(query, "Prod_Title", "S", 0, 0, p.ProdTitle)
	query = tools.ArmoSentencia(query, "Prod_Description", "S", 0, 0, p.ProdDescription)
	query = tools.ArmoSentencia(query, "Prod_Price", "F", 0, p.ProdPrice, "")
	query = tools.ArmoSentencia(query, "Prod_CategoryId", "N", p.ProdCategId, 0, "")
	query = tools.ArmoSentencia(query, "Prod_Stock", "N", p.ProdStock, 0, "")
	query = tools.ArmoSentencia(query, "Prod_Path", "S", 0, 0, p.ProdPath)

	query += " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)

	fmt.Println(query)
	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Product > Ejecución Exitosa")
	
	return nil
}

func DeleteProduct(id int) error {
	fmt.Println("Comienza DELETE Product")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	
	fmt.Println("Delete Product > Ejecución Exitosa")
	return nil
}

func SelectProduct(p models.Product, choice string, page int, pageSize int, orderType string, orderField string) (models.ProductResp, error) {
	fmt.Println("Comienza SelectProduct")
	var Resp models.ProductResp
	var Prod []models.Product

	err := DbConnect()
	if err != nil {
		return Resp, err
	}
	defer Db.Close()

	var queryP string
	var queryQ string
	var where, limit string

	queryP = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products "
	queryQ = "SELECT COUNT(*) as registros FROM products "

	switch choice {
	case "P":
		where = "WHERE Prod_Id = " + strconv.Itoa(p.ProdId)
	case "S":
		where = "WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%' "
	case "C":
		where = "WHERE Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
	case "U":
		where = "WHERE UCASE(Prod_Path) LIKE '%" + strings.ToUpper(p.ProdPath) + "%' "
	case "K":
		join := "JOIN category ON Prod_CategoryId = Categ_Id AND Categ_Path LIKE '%" + strings.ToUpper(p.ProdCategPath) + "%' "
		queryP += join
		queryQ += join
	}

	queryQ += where

	// ⚠️ Ejecutar queryQ y validar error antes de usar rows
	rowsCount, err := Db.Query(queryQ)
	if err != nil {
		fmt.Println("Error en queryQ:", err.Error())
		return Resp, err
	}
	defer rowsCount.Close()

	var regis sql.NullInt32
	if rowsCount.Next() {
		err = rowsCount.Scan(&regis)
		if err != nil {
			fmt.Println("Error al escanear registros:", err.Error())
			return Resp, err
		}
	} else {
		fmt.Println("No se encontraron registros")
		return Resp, nil
	}
	registros := int(regis.Int32)

	if page > 0 {
		if registros > pageSize {
			limit = "LIMIT " + strconv.Itoa(pageSize)
			if page > 1 {
				offset := pageSize * (page - 1)
				limit += " OFFSET " + strconv.Itoa(offset)
			}
		}
	}

	var orderBy string
	if len(orderField) > 0 {
		switch orderField {
		case "I":
			orderBy = " ORDER BY Prod_Id "
		case "T":
			orderBy = " ORDER BY Prod_Title "
		case "D":
			orderBy = " ORDER BY Prod_Description "
		case "F":
			orderBy = " ORDER BY Prod_CreatedAt "
		case "P":
			orderBy = " ORDER BY Prod_Price "
		case "S":
			orderBy = " ORDER BY Prod_Stock "
		case "C":
			orderBy = " ORDER BY Prod_CategoryId "
		}
	}

	if orderType == "D" {
		orderBy += " DESC "
	}

	queryP += where + orderBy + " " + limit

	fmt.Println("Query Products > " + queryP)
	fmt.Println("Query Quantity Products >" + queryQ)

	rows, err := Db.Query(queryP)
	if err != nil {
		fmt.Println("Error en queryP:", err.Error())
		return Resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		var ProdId sql.NullInt32
		var ProdTitle sql.NullString
		var ProdDescription sql.NullString
		var ProdCreatedAt sql.NullTime
		var ProdUpdated sql.NullTime
		var ProdPrice sql.NullFloat64
		var ProdPath sql.NullString
		var ProdCategoryId sql.NullInt32
		var ProdStock sql.NullInt32

		err := rows.Scan(&ProdId, &ProdTitle, &ProdDescription, &ProdCreatedAt, &ProdUpdated, &ProdPrice, &ProdPath, &ProdCategoryId, &ProdStock)
		if err != nil {
			return Resp, err
		}

		p.ProdId = int(ProdId.Int32)
		p.ProdTitle = ProdTitle.String
		p.ProdDescription = ProdDescription.String
		p.ProdCreatedAt = ProdCreatedAt.Time.String()
		p.ProdUpdated = ProdUpdated.Time.String()
		p.ProdPrice = ProdPrice.Float64
		p.ProdPath = ProdPath.String
		p.ProdCategId = int(ProdCategoryId.Int32)
		p.ProdStock = int(ProdStock.Int32)

		Prod = append(Prod, p)
	}

	Resp.TotalItems = registros
	Resp.Data = Prod

	fmt.Println("Select Product > Ejecución Exitosa")
	return Resp, nil
}
