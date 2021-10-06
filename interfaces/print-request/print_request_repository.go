package print_request

import "threedee/entity"

/*
 * FOURTH LAYER => Repository package and/or Service Package
 *
 * Repository is the fourth layer of the service, and so does Service. However,
 * there is no service in threedee.
 *
 * It is better to create the repo/service interface before writing the repo/service
 * code in order to assist the actual code writing. This way, method input/output mistypes
 * can be prevented.
 *
 * In threedee, the actual repo code is written in "repository/print_request_repository.go".
 * The interface is the one that is utilized in the Handler file instead of the repo. It is
 * more like a framework for the actual repo package. We instantiate the actual repo instance
 * in threede.go (second layer). This way, if there are methods in repo package that does not
 * follow the interface, we will be notified.
 */

type PrintRequestRepositoryInterface interface {
	GetAll() ([]*entity.PrintRequest, error)
	GetById(id int) (*entity.PrintRequest, error)
	Insert(model *entity.PrintRequest) (int, error)
	Update(model *entity.PrintRequest) (bool, error)
	Delete(id int) (bool, error)
}
