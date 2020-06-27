package db

import (
	"database/sql"
	"fmt"
	"strings"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

//数据库配置
const (
	userName = "root"
	password = "201592009"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "wallpaper"
)
//Db数据库连接池
var MysqlClient *sql.DB

//注意方法名大写，就是public
func InitDB()  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	var err error
	MysqlClient, err = sql.Open("mysql", path)
	if err != nil {
		fmt.Printf("WRONG!!! %v\n", err)
	}
	//设置数据库最大连接数
	MysqlClient.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	MysqlClient.SetMaxIdleConns(10)
	//验证连接
	if err := MysqlClient.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func CreateTables() {
	_,_ = MysqlClient.Exec("create table user (\nuid int primary key auto_increment,\nname varchar(200),\npassword varchar(200),\npid varchar(200)\n);")
	_,_ = MysqlClient.Exec("create table image (\npid int primary key auto_increment,\nname varchar(200),\nclassify varchar(20),\nfilename varchar(200),\nuid int\n);")
	_,_ = MysqlClient.Exec("create table model (\npid int primary key auto_increment,\nname varchar(200),\nclassify varchar(20),\nmodel_path varchar(200),\nimages_path varchar(200),\nuid int\n);")
}

func InitSQLiteDB() (err error) {
	MysqlClient, err = sql.Open("sqlite3","wallpaper.db")
	if err != nil {
		fmt.Printf("WRONG!!! %v\n", err)
		return
	}
	//CreateTables()
	return
}

func CheckUser(username string) bool {
	var uid int
	query := fmt.Sprintf("SELECT uid FROM user WHERE name = '%+v'", username)
	fmt.Println(query)
	err := MysqlClient.QueryRow(query).Scan(&uid)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func Validate(username string, password string) int64 {
	var uid int64
	query := fmt.Sprintf("SELECT uid FROM user WHERE name = '%v' and password = '%v'", username, password)
	err := MysqlClient.QueryRow(query).Scan(&uid)
	fmt.Println(query)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return -1
	}
	if err != nil {
		return -1
	}
	return uid
}

func InsertUser(username string, password string)(error, int64) {
	tx, err := MysqlClient.Begin()
	if err != nil {
		return err, 0
	}
	stmt, err := tx.Prepare("INSERT INTO user (name, password, pid) VALUES (?, ?, ?)")
	if err != nil {
		return err, 0
	}
	res, err := stmt.Exec(username, password, "")
	if err != nil {
		return err, 0
	}
	tx.Commit()
	uid, err := res.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, uid
}

func SaveImg(filename string, name string, classify string, uid int64) (error, int64) {
	tx, err := MysqlClient.Begin()
	if err != nil {
		return err, 0
	}
	stmt, err := tx.Prepare("INSERT INTO image (name, classify, filename, uid) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err, 0
	}
	res, err := stmt.Exec(name, classify, filename, uid)
	if err != nil {
		return err, 0
	}
	tx.Commit()
	pid, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return err, 0
	}
	return nil, pid
}

func GetUserImg(uid int64) string {
	var pid string
	query := fmt.Sprintf("SELECT pid FROM user WHERE uid = %+v", uid)
	err := MysqlClient.QueryRow(query).Scan(&pid)
	if err == sql.ErrNoRows {
		return ""
	}
	return pid
}

func UpdateUserImg(uid int64, pid int64) error {
	oldPid := GetUserImg(uid)
	var newPid string
	if oldPid == "" {
		newPid = fmt.Sprintf("%v", pid)
	} else {
		newPid = fmt.Sprintf("%v,%v", oldPid, pid)
	}
	//开启事务
	tx, err := MysqlClient.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET pid = ? WHERE uid = ?")
	if err != nil{
		return err
	}
	//设置参数以及执行sql语句
	_, err = stmt.Exec(newPid, uid)
	if err != nil{
		return err
	}
	//提交事务
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func FetchImg(uid int64, classify string) ([]string, error){

	query := fmt.Sprintf("SELECT name, filename FROM image where uid = %v and classify = '%v' ", uid, classify)
	rows, err := MysqlClient.Query(query)
	if err != nil {
		return nil, err
	}
	var filename string
	var imgname string
	result := make([]string, 0)
	for rows.Next(){
		err := rows.Scan(&imgname, &filename)
		if err != nil {
			return nil, err
		}
		message := fmt.Sprintf("%s@%s", imgname, filename)
		result = append(result, message)
	}
	return result, nil
}

func FetchClass(uid int64) ([]string, error) {
	query := fmt.Sprintf("SELECT distinct classify FROM image where uid = %v ", uid)
	rows, err := MysqlClient.Query(query)
	if err != nil {
		return nil, err
	}
	var class string
	result := make([]string, 0)
	for rows.Next() {
		err := rows.Scan(&class)
		if err != nil {
			return nil, err
		}
		result = append(result, class)
	}
	return result, nil
}

func InsertModel(uid int64,modelPath ,imagesPath,classify ,name string) error {
	tx, err := MysqlClient.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO model (name, classify, model_path,image_path, uid) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, classify, modelPath,imagesPath,uid)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

type ModelInfo struct {
	Name string `json:"name"`
	Classify string `json:"classify"`
	Image_path string `json:"image_path"`
}

func FetchModel(uid int64) ([]ModelInfo, error) {
	query := fmt.Sprintf("SELECT name,classify,image_path FROM model where uid = %v ", uid)
	rows, err := MysqlClient.Query(query)
	if err != nil {
		return nil, err
	}
	var modelInfo ModelInfo
	result := make([]ModelInfo, 0)
	for rows.Next() {
		err := rows.Scan(&modelInfo.Name,&modelInfo.Classify,&modelInfo.Image_path)
		if err != nil {
			return nil, err
		}
		result = append(result, modelInfo)
		fmt.Printf("FetchModel modelInfo=%+v",modelInfo)
	}

	return result, nil
}

