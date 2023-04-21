package drivers

import (
	"database/sql"
	"time"

    _ "github.com/jackc/pgconn"
    _ "github.com/jackc/pgx"
    _ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{} 

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const MaxDbLifetime = 5 *time.Minute

// Creates database pool for postgress
func ConnectSQL(dsn string) (*DB, error){
     d, err := NewDatabase(dsn)
     if err != nil{
        panic(err)
     }

     d.SetMaxOpenConns(maxOpenDbConn)
     d.SetMaxIdleConns(maxIdleDbConn)
     d.SetConnMaxLifetime(MaxDbLifetime)
     dbConn.SQL=d
     err =testDB(dbConn.SQL)
     if err != nil {
        return nil, err
     }
     return dbConn, nil
}

// testdb pings the database
func testDB(d *sql.DB) ( error){
        err := d.Ping()
        if err != nil{
        return err
        }

    return nil
}
//creates new database
func NewDatabase(dsn string)(*sql.DB,error){
    db, err := sql.Open("pgx",dsn)
    if err != nil {
        return  nil , err 
    }

    if err = db.Ping();err !=nil {
        return nil , err
    }


    return db, nil
   

}
