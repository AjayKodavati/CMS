package db

import "time"

type Coupon struct {
	CouponID              int       `json:"couponID"`
	CouponCode            string    `json:"couponCode"`
	DiscountType          string    `json:"discountType"`
	DiscountValue         float64   `json:"discountValue"`
	TermsAndConditions    string    `json:"termsAndConditions"`
	UsageType             string    `json:"usageType"`
	ValidFrom             time.Time `json:"validFrom"`
	ValidUntil            time.Time `json:"validUntil"`
	MinimumPurchaseAmount float64   `json:"minimumPurchaseAmount"`
}

type Medicine struct {
	MedicineID   int    `json:"medicineID"`
	MedicineName string `json:"medicineName"`
	CategoryID   int    `json:"categoryID"`
}

type CouponMedicine struct {
	CouponID   int
	MedicineID int
}

type Category struct {
	CategoryID   int    `json:"categoryID"`
	CategoryName string `json:"categoryName"`
}
