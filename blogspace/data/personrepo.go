package db

type Person struct {
  FirstName string
  LastName  string
  Email     string
  Password  string
  Liked     []BlogID
}

type BlogID string

type Blog struct {
  Title       string
  Description string
  Content     string
}

func CreatePerson() {

}
