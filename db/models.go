package db

import "time"

type Coupon struct {
	CouponID              int
	CouponCode            string
	DiscountType          string
	DiscountValue         float64
	TermsAndConditions    string
	UsageType             string
	ValidFrom             time.Time
	ValidTo               time.Time
	MinimumPurchaseAmount float64
}

type Medicine struct {
	MedicineID      int
	MedicineName    string
	CategoryID      int
}

type CouponMedicine struct {
	CouponID    int
	MedicineID  int
}