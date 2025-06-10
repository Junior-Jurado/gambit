package bd

import(
	"fmt"
	"database/sql"
	// "strconv"
	// "strings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Junior_Jurado/gambit/models"
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