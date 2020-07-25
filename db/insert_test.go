package db

import "testing"

func TestInsert(t *testing.T)  {
	InitSQLiteDB()
	CreateTables()
	InsertModel(0, "/static/model/Pikachu.stl", "", "public", "pikachu.stl")
	InsertModel(0, "/static/model/Eevee.stl", "", "public", "eevee.stl")
	//InsertModel(0, "/static/model/Catbus.stl", "", "public", "Catbus.stl")
	//InsertModel(0, "/static/model/Chinese_Dragon.stl", "", "public", "Chinese_Dragon.stl")
	//InsertModel(0, "/static/model/Dragonite.stl", "", "public", "Dragonite.stl")
}
