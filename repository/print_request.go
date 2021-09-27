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

	rows, err := db.Query("select " +
		"a.id," +
		"a.item_name," +
		"a.est_weight," +
		"a.est_filament_length," +
		"a.est_duration," +
		"a.file_url," +
		"a.requestor," +
		"a.status " +
		"from tbl_m_3d_print_request a")
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

	rows, err := db.Query("select "+
		"a.id,"+
		"a.item_name,"+
		"a.est_weight,"+
		"a.est_filament_length,"+
		"a.est_duration,"+
		"a.file_url,"+
		"a.requestor,"+
		"a.status "+
		"from tbl_m_3d_print_request a where a.id = $1", id)
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

func (*PrintRequestRepository) Insert(model *entity.PrintRequest) (int, error) {
	db, err := database.NewPostgresql()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var lastInsertId *int
	err = db.QueryRow("INSERT INTO tbl_m_3d_print_request("+
		"item_name,"+
		"est_weight,"+
		"est_filament_length,"+
		"est_duration,"+
		"file_url,"+
		"requestor) "+
		"VALUES "+
		"($1,"+
		"$2,"+
		"$3,"+
		"$4,"+
		"$5,"+
		"$6) "+
		"RETURNING id;",
		model.ItemName,
		model.EstimatedWeight,
		model.EstimatedFilamentLength,
		model.EstimatedDuration,
		model.FileUrl,
		model.Requestor).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return *lastInsertId, nil
}

func (*PrintRequestRepository) Update(model *entity.PrintRequest) (bool, error) {
	db, err := database.NewPostgresql()
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, err = db.Query("UPDATE tbl_m_3d_print_request SET "+
		"item_name = $1,"+
		"est_weight = $2,"+
		"est_filament_length = $3,"+
		"est_duration = $4,"+
		"file_url = $5,"+
		"requestor = $6,"+
		"status = $7 "+
		"WHERE id = $8;",
		model.ItemName,
		model.EstimatedWeight,
		model.EstimatedFilamentLength,
		model.EstimatedDuration,
		model.FileUrl,
		model.Requestor,
		model.Status,
		model.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (*PrintRequestRepository) Delete(id int) (bool, error) {
	db, err := database.NewPostgresql()
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, err = db.Query("UPDATE tbl_m_3d_print_request SET "+
		"is_active = false "+
		"WHERE id = $1;",
		id)
	if err != nil {
		return false, err
	}
	return true, nil
}
