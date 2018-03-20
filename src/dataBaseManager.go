package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:lexco123@tcp(127.0.0.1:3306)/sgdea")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT numero_radicacion FROM radicacion WHERE pdf_idPDF = ?")
	if err != nil {
		panic(err.Error())
	}
	printSelectResult(stmtOut,err)
	stmtOut.Close()

	stmtInsert, err := db.Prepare("INSERT INTO roles (nom_rol_usuario,esta_rol_usuario,DEPENDENCIAS_idDEPENDENCIAS) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	res, err := stmtInsert.Exec("((delete * from roles))","1","109")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res)

	stmtInsert.Close()

}

func printSelectResult(stmtOut *sql.Stmt,err error )  {
	var squareNum string

	err = stmtOut.QueryRow(40).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("el numero de radicacion asociado al pdf # 40 es : %d", squareNum)
}