package main

import (
	"database/sql"
	"fmt"
	"log"
	"mssqllib/pkg/mssql"
)

const (
	spName = "Reporting_C3.dbo.usp_GetMetaDataCMFL"
	isProd = false
)

func main() {
	qry := "exec " + spName
	if isProd == true {
		qry = qry + "_p"
	} else {
		qry = qry + "_t"
	}
	qry = qry + " @n=?"
	fmt.Println(qry)

	cfg := mssql.Config{URL: "sqlserver://user:pass@YourDbServer?database=ThatDBName"}
	db, err := mssql.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	var sqlversion string

	defer db.Database.Close()
	// rows, err := db.Database.Query("select @@version")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for rows.Next() {
	// 	err := rows.Scan(&sqlversion)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(sqlversion)
	// }

	row := db.Database.QueryRow("select @@version")
	err = row.Scan(&sqlversion)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No Rows Returned \n")
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Println(sqlversion)
	}

}
