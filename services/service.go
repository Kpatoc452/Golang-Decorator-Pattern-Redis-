package service

import "github.com/Kpatoc452/Dceorator/models"


type Info interface{
	GetInfo(id int) (models.Info, error)
	AddInfo(info string) error
	DeleteInfo(id int) error
	GetAllInfo() ([]string, error)
}

type Service struct{
	Info
}

func New() *Service{
	return &Service{
		
	}
}