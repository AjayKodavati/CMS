package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CouponCategoriesRepository interface {
	GetCategoryCouponIDs(ctx context.Context, categoryID int) ([]int, error)
}

type CouponCategoriesRepositoryService struct {
	pool *pgxpool.Pool
}

func NewCouponCategoriesRepository(pool *pgxpool.Pool) CouponCategoriesRepository {
	return &CouponCategoriesRepositoryService{
		pool: pool,
	}
}

func (c *CouponCategoriesRepositoryService) GetCategoryCouponIDs(ctx context.Context, categoryID int) ([]int, error) {
	var couponIDs []int
	query := `SELECT coupon_id FROM coupon_categories WHERE category_id = $1`
	rows, err := c.pool.Query(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var couponID int
		if err := rows.Scan(&couponID); err != nil {
			return nil, err
		}
		couponIDs = append(couponIDs, couponID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return couponIDs, nil
}