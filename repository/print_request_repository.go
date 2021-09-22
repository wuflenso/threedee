package repository

import (
	"threedee/database"
	"threedee/entity"
)

// Somehow at this point, we no longer play with pointers
// is it because we dont modify that much?

type PrintRequestRepository struct {
}

func NewPrintRequestRepository() *PrintRequestRepository {
	return &PrintRequestRepository{}
}

func (*PrintRequestRepository) GetAll() ([]*entity.PrintRequest, error) {
	db, err := database.NewPostgresql()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select a.id, a.item_name, a.est_weight, a.est_filament_length, a.est_duration,	a.file_url,	a.requestor, a.status from tbl_m_3d_print_request a")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entity.PrintRequest, 0)
	for rows.Next() {
		item := entity.NewPrintRequest()
		err := rows.Scan(
			&item.Id,
			&item.ItemName,
			&item.EstimatedWeight,
			&item.EstimatedFilamentLength,
			&item.EstimatedDuration,
			&item.FileUrl,
			&item.Requestor,
			&item.Status,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*PrintRequestRepository) GetById(id int) (*entity.PrintRequest, error) {
	db, err := database.NewPostgresql()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select a.id, a.item_name, a.est_weight, a.est_filament_length, a.est_duration,	a.file_url,	a.requestor, a.status from tbl_m_3d_print_request a where a.id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	item := entity.NewPrintRequest()
	for rows.Next() {
		err := rows.Scan(
			&item.Id,
			&item.ItemName,
			&item.EstimatedWeight,
			&item.EstimatedFilamentLength,
			&item.EstimatedDuration,
			&item.FileUrl,
			&item.Requestor,
			&item.Status,
		)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return item, nil
}
