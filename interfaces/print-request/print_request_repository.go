package print_request

import "threedee/entity"

type PrintRequestRepositoryInterface interface {
	GetAll() ([]*entity.PrintRequest, error)
	GetById(id int) (*entity.PrintRequest, error)
	Insert(model *entity.PrintRequest) (int, error)
	Update(model *entity.PrintRequest) (bool, error)
	Delete(id int) (bool, error)
}
