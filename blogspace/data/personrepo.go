package data

import (
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)

var DB *sqlx.DB

type Person struct {
  ID        int64  `db:"id" json:"id"` 
  FirstName string `db:"first_name" json:"first_name"`
  LastName  string `db:"last_name" json:"last_name"`
  Username  string `json:"username"`
  Email     string `json:"email"`
  Password  string `json:"password"`
  Liked     []BlogID `json:"liked"`
}

type BlogID int64

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
      email      text,
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

  return nil
}
