package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CouponMedicinesRepository interface {
	GetCouponIdByMedicineID(ctx context.Context, medicineID int) (int, error)
}

type CouponMedicinesRepositoryService struct {
	pool *pgxpool.Pool	
}

func NewCouponMedicinesRepository(pool *pgxpool.Pool) *CouponMedicinesRepositoryService {
	return &CouponMedicinesRepositoryService{
		pool: pool,
	}
}

func (c *CouponMedicinesRepositoryService) GetCouponIdByMedicineID(ctx context.Context, medicineID int) (int, error) {
	var couponID int
	query := `SELECT coupon_id FROM coupon_medicines WHERE medicine_id = $1`
	err := c.pool.QueryRow(ctx, query, medicineID).Scan(&couponID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, nil // No coupon found for this medicine
		}
		return 0, err // Return the error if something else went wrong
	}
	return couponID, nil 
}