package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type storage interface{
     CreateShortUrl(*ShortUrl) error
	 GetUrl(string)(*ShortUrl,error)
	 GetlongUrl(string)(*ShortUrl,error)
}
type PostgressStore struct{
	db *sql.DB
}
func NewPostgressStore() (*PostgressStore,error){
	err:= godotenv.Load()
	dblink:= os.Getenv("DBLINK")
	conStr:= dblink
	db, err:= sql.Open("postgres",conStr)
	if err!=nil{
		panic(err)
	}
	if err:=db.Ping();err!=nil{
		return nil,err
	}
	return &PostgressStore{
		db: db,
	},nil
}

func (s *PostgressStore) init() (error) {
    return s.CreateURLTable()
}
func (s *PostgressStore) CreateURLTable() error{
	query:= `create table if not exists URL ( 
		id serial primary key, 
        alias varchar(50),
		long_url varchar(300)
		)`
		_,err := s.db.Exec(query)
		return err
}
func (s *PostgressStore)CreateShortUrl(u *ShortUrl)error{
    qeuery:= `insert into URL
	(alias, long_url)
	values($1,$2)`
	_,err:= s.db.Query(
		qeuery,
		u.Alias,
		u.LongURL,
	)
	if err!=nil{
		return err
	}
	return nil
}
func (s *PostgressStore)GetUrl(ali string)(*ShortUrl,error){
	rows,err:=s.db.Query(`SELECT * from URL
	where alias = $1`,ali)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		return ScanintoURLs(rows)
	}
	return nil,fmt.Errorf("Failed to get card")

}
func (s *PostgressStore)GetlongUrl(url string)(*ShortUrl,error){
	rows,err:=s.db.Query(`SELECT * from URL
	where long_url = $1`,url)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		return ScanintoURLs(rows)
	}
	return nil,fmt.Errorf("Failed to get card")

}
func ScanintoURLs(rows *sql.Rows) (*ShortUrl,error){
	url:=new(ShortUrl)
	err:=rows.Scan(&url.ID,&url.Alias,&url.LongURL)
	if err!=nil{
		return nil,err
	}
    return url,nil
}