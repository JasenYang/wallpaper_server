package db

import "testing"

func TestInsert(t *testing.T)  {
	InitSQLiteDB()
	CreateTables()
	InsertModel(0, "/static/model/Pikachu.stl", "", "public", "pikachu.stl")
	InsertModel(0, "/static/model/Eevee.stl", "", "public", "eevee.stl")
}
