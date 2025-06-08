package repository

import (
	"context"

	"github.com/AjayKodavati/CMS/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

/**
 * CouponRepository defines the interface for coupon-related database operations.
 * It includes methods for creating, deleting, updating, and retrieving coupons.
*/

type CouponRepository interface {
	CreateCoupon(ctx context.Context, coupon *db.Coupon) error
	DeleteCoupon(ctx context.Context, couponCode string) error
	UpdateCoupon(ctx context.Context, coupon *db.Coupon) error
	GetCouponByID(ctx context.Context, couponCode string) (*db.Coupon, error)
}

/**
 * couponRepositoryService implements the CouponRepository interface following the dependency inversion principle.
 * It provides methods to interact with the database for coupon operations.
*/

type couponRepositoryService struct {
	pool *pgxpool.Pool	
}

func NewCouponRepository(pool *pgxpool.Pool) CouponRepository {
	return &couponRepositoryService{
		pool: pool,
	}
}

func (c *couponRepositoryService) CreateCoupon(ctx context.Context, coupon *db.Coupon) error {
	query := `INSERT INTO coupons (coupon_code, discount_type, discount_value, terms_and_conditions, usage_type, valid_from, valid_until, minimum_purchase_amount) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING coupon_id`
	_, err := c.pool.Exec(ctx, query,
		coupon.CouponCode,
		coupon.DiscountType,
		coupon.DiscountValue,
		coupon.TermsAndConditions,
		coupon.UsageType,
		coupon.ValidFrom,
		coupon.ValidUntil,
		coupon.MinimumPurchaseAmount,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponRepositoryService) DeleteCoupon(ctx context.Context, couponCode string) error {
	query := `DELETE FROM coupons WHERE coupon_code = $1`
	_, err := c.pool.Exec(ctx, query, couponCode)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponRepositoryService) UpdateCoupon(ctx context.Context, coupon *db.Coupon) error {
	query := `UPDATE coupons SET coupon_code = $1, discount_type = $2, discount_value = $3, terms_and_conditions = $4, usage_type = $5, valid_from = $6, valid_until = $7, minimum_purchase_amount = $8 
			  WHERE coupon_id = $9`
	_, err := c.pool.Exec(ctx, query,
		coupon.CouponCode,
		coupon.DiscountType,
		coupon.DiscountValue,
		coupon.TermsAndConditions,
		coupon.UsageType,
		coupon.ValidFrom,
		coupon.ValidUntil,
		coupon.MinimumPurchaseAmount,
		coupon.CouponID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponRepositoryService) GetCouponByID(ctx context.Context, couponCode string) (*db.Coupon, error) {
	query := `SELECT coupon_id, coupon_code, discount_type, discount_value, terms_and_conditions, usage_type, valid_from, valid_until, minimum_purchase_amount 
			  FROM coupons WHERE coupon_code = $1`
	row := c.pool.QueryRow(ctx, query, couponCode)

	var coupon db.Coupon
	err := row.Scan(
		&coupon.CouponID,
		&coupon.CouponCode,
		&coupon.DiscountType,
		&coupon.DiscountValue,
		&coupon.TermsAndConditions,
		&coupon.UsageType,
		&coupon.ValidFrom,
		&coupon.ValidUntil,
		&coupon.MinimumPurchaseAmount,
	)
	if err != nil {
		return nil, err
	}
	return &coupon, nil
}