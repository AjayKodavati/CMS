package repository

import "github.com/jackc/pgx/v5/pgxpool"

type DBRepositories struct {
	CouponRepository                 CouponRepository
	CategoriesRepository             CategoriesRepository
	MedicineRepositoryService        MedicineRepository
	CouponCategoriesRepository       CouponCategoriesRepository
	CouponMedicinesRepositoryService CouponMedicinesRepository
}

func SetUpDBRepositories(pool *pgxpool.Pool) *DBRepositories {
	return &DBRepositories{
		CouponRepository:                 NewCouponRepository(pool),
		CategoriesRepository:             NewCategoriesRepository(pool),
		MedicineRepositoryService:        NewMedicineRepository(pool),
		CouponCategoriesRepository:       NewCouponCategoriesRepository(pool),
		CouponMedicinesRepositoryService: NewCouponMedicinesRepository(pool),
	}
}