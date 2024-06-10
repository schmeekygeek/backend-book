package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

type Person struct {
  ID        int8  `db:"id" json:"id"` 
  FirstName string `db:"first_name" json:"first_name"`
  LastName  string `db:"last_name" json:"last_name"`
  Username  string `json:"username"`
  Email     string `json:"email"`
  Password  string `json:"password"`
  Liked     []BlogID `db:"liked" json:"liked"`
}

type BlogID int8

type Blog struct {
  Title       string
  Description string
  Content     string
}

const (
  CREATEPERSONTABLE = `
    CREATE TABLE IF NOT EXISTS PERSON (
      id serial unique,
      first_name text,
      last_name  text,
      email      text unique,
      username   text unique,
      password   text,
      liked int8[]
    );
  `
)

func Init() error {
  db, err := sqlx.Connect(
    "postgres",
    "user=postgres dbname=blogspace password=1213 sslmode=disable",
  )
  if err != nil {
    return err
  }
  DB = db
  DB.MustExec(CREATEPERSONTABLE)
  return nil
}

func CreatePerson(p *Person) error {
  tx := DB.MustBegin()
  tx.MustExec(
    "INSERT INTO person (first_name, last_name, email, username, password, liked) VALUES ($1, $2, $3, $4, $5, $6)",
    p.FirstName,
    p.LastName,
    p.Email,
    p.Username,
    p.Password,
    pq.Array(p.Liked),
  )
  tx.Commit()
  return nil
}
