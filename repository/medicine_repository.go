package repository

import (
	"context"

	"github.com/AjayKodavati/CMS/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicineRepository interface {
	CreateMedicine(ctx context.Context, medicine *db.Medicine) error
	DeleteMedicine(ctx context.Context, medicineID int) error
	UpdateMedicine(ctx context.Context, medicine *db.Medicine) error
	GetMedicineByID(ctx context.Context, medicineID int) (*db.Medicine, error)
	GetAllMedicines(ctx context.Context) ([]*db.Medicine, error)
	GetMedicinesByCategory(ctx context.Context, categoryID int) ([]*db.Medicine, error)
	GetMedicinesByName(ctx context.Context, name string) ([]*db.Medicine, error)
}

type MedicineRepositoryService struct{
	pool *pgxpool.Pool
}

func NewMedicineRepository(pool *pgxpool.Pool) MedicineRepository {
	return &MedicineRepositoryService{pool: pool}
}

func (m *MedicineRepositoryService) CreateMedicine(ctx context.Context, medicine *db.Medicine) error {
	query := `INSERT INTO medicines (medicine_name, category_id) VALUES ($1, $2)`
	_, err := m.pool.Exec(ctx, query, medicine.MedicineName, medicine.CategoryID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MedicineRepositoryService) DeleteMedicine(ctx context.Context, medicineID int) error {
	query := `DELETE FROM medicines WHERE medicine_id = $1`
	_, err := m.pool.Exec(ctx, query, medicineID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MedicineRepositoryService) UpdateMedicine(ctx context.Context, medicine *db.Medicine) error {
	query := `UPDATE medicines SET medicine_name = $1, category_id = $2 WHERE medicine_id = $3`
	_, err := m.pool.Exec(ctx, query, medicine.MedicineName, medicine.CategoryID, medicine.MedicineID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MedicineRepositoryService) GetMedicineByID(ctx context.Context, medicieneId int) (*db.Medicine, error) {
	query := `SELECT medicine_id, medicine_name, category_id FROM medicines WHERE medicine_id = $1`
	row := m.pool.QueryRow(ctx, query, medicieneId)
	medicine := &db.Medicine{}
	err := row.Scan(&medicine.MedicineID, &medicine.MedicineName, &medicine.CategoryID)
	if err != nil {
		return nil, err
	}
	return medicine, nil
}

func (m *MedicineRepositoryService) GetAllMedicines(ctx context.Context) ([]*db.Medicine, error) {
	query := `SELECT medicine_id, medicine_name, category_id FROM medicines`
	rows, err := m.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medicines []*db.Medicine
	for rows.Next() {
		medicine := &db.Medicine{}
		err := rows.Scan(&medicine.MedicineID, &medicine.MedicineName, &medicine.CategoryID)
		if err != nil {
			return nil, err
		}
		medicines = append(medicines, medicine)
	}
	return medicines, nil
}

func (m *MedicineRepositoryService) GetMedicinesByCategory(ctx context.Context, categoryID int) ([]*db.Medicine, error) {
	query := `SELECT medicine_id, medicine_name, category_id FROM medicines WHERE category_id = $1`
	rows, err := m.pool.Query(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medicines []*db.Medicine
	for rows.Next() {
		medicine := &db.Medicine{}
		err := rows.Scan(&medicine.MedicineID, &medicine.MedicineName, &medicine.CategoryID)
		if err != nil {
			return nil, err
		}
		medicines = append(medicines, medicine)
	}
	return medicines, nil
}

func (m *MedicineRepositoryService) GetMedicinesByName(ctx context.Context, name string) ([]*db.Medicine, error) {
	query := `SELECT medicine_id, medicine_name, category_id FROM medicines WHERE medicine_name ILIKE $1`
	rows, err := m.pool.Query(ctx, query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medicines []*db.Medicine
	for rows.Next() {
		medicine := &db.Medicine{}
		err := rows.Scan(&medicine.MedicineID, &medicine.MedicineName, &medicine.CategoryID)
		if err != nil {
			return nil, err
		}
		medicines = append(medicines, medicine)
	}
	return medicines, nil
}