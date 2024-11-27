package postgres

import (
	"fmt"

	"github.com/Kpatoc452/Dceorator/models"
)

type DB interface{
	GetInfo(id int) (models.Info, error)
	AddInfo(info models.Info) error
	DeleteInfo(id int) error	
}

type database struct{
}

func New() *database{
	return &database{}
}

func(db *database) GetPerson(id int) (models.Person, error) {
	p := models.Person{
		Id: 10,
		FirstName: "mikhail",
		LastName: "Efimenko",
		City: "Vyksa",
	}
	return p, nil; 
}

func(db *database) AddPerson(info models.Info) error{
	fmt.Println("person added")
	return nil
}

func DeletePerson(id int) error{
	fmt.Println("person deleted")
	return nil
}