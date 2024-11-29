package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/Kpatoc452/Dceorator/config"
	"github.com/Kpatoc452/Dceorator/models"
	"github.com/jackc/pgx/v5"
)

var ctx = context.Background()

type DB interface{
	GetInfo(id int) (models.Info, error)
	AddInfo(info models.Info) error
	DeleteInfo(id int) error	
}

type database struct{
	conn *pgx.Conn
}

func New(cfg config.Config) (*database, error){
	urlDB := cfg.GetPSQL()
	conn, err := pgx.Connect(ctx, urlDB)
	if err != nil{
		log.Println("PostgreSQL. Cannot connect to Database.")
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil{
		log.Println("PostgreSQL. Error in Ping Database.")
		return nil, err
	}

	return &database{
		conn: conn,
	}, nil

}

func(db *database) GetInfo(id int) (models.Info, error) {
	var info models.Info
	queryStr := fmt.Sprintf("SELECT * FROM tasks WHERE TaskID=%d", id)
	row, err := db.conn.Query(ctx, queryStr)
	if err != nil{
		log.Println("PostgreSQL. Error GetInfo by id.")
		return info, err
	}

	err = row.Scan(&info)
	if err != nil {
		log.Println("PostgreSQL. Error write in models.Info")
		return info, err
	}
	return info, nil
}

func(db *database) AddInfo(info string) error{
	_, err := db.conn.Exec(ctx,  "insert into tasks(description) values($1)", info)
	if err != nil {
		log.Println("PostgreSQL. Error Add Info.")
	}
	return err
}

func (db *database) DeleteInfo(id int) error{
	fmt.Println("person deleted")
	return nil
}

func (db *database) GetAllInfo() ([]string, error) {
	rows, err := db.conn.Query(ctx, "select description from tasks")
	if err != nil {
		log.Println("GetAllInfo PSQL. Error with Query")
		return nil, err
	}

	var description string
	var res []string

	for rows.Next() {
		rows.Scan(&description)
		res = append(res, description)
	}
	return res, err
}