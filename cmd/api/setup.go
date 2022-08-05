package main

import "database/sql"

type app struct {
	DB *sql.DB
	routes newRouter()
}