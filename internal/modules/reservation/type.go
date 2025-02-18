package reservation

import "database/sql"

type ReservationQueryAll struct {
	KeyWork string `query:"keywork" json:"keywork" validate:"omitempty,max=255"`
}

type ReservationResponse struct {
	Id           string         `boil:"id" json:"id"`
	CustomerName sql.NullString `boil:"customer_name" json:"customer_name"`
	CuaHangId    string         `boil:"cuahang_id" json:"cuahang_id"`
	Phong        sql.NullString `boil:"phong" json:"phong"`
	Amount       string         `boil:"amount" json:"amount"`
	Mobile       string         `boil:"mobile" json:"mobile"`
	Date       	 string         `boil:"date" json:"date"`
	Email        sql.NullString `boil:"email" json:"email"`
	StoreNote    sql.NullString `boil:"store_note" json:"store_note"`
	CustomerNote sql.NullString `boil:"customer_note" json:"customer_note"`
	ApprovedBy   sql.NullString `boil:"approved_by" json:"approved_by"`
	ApprovedDate sql.NullString `boil:"approved_date" json:"approved_date"`
	PickupBy     sql.NullString `boil:"pickup_by" json:"pickup_by"`
	PickupDate   sql.NullString `boil:"pickup_date" json:"pickup_date"`
	Reason       sql.NullString `boil:"reason" json:"reason"`
}

type ReservationResponseData struct {
	Id           string  `boil:"id" json:"id"`
	CustomerName *string `boil:"customer_name" json:"customer_name"`
	CuaHangId    string  `boil:"cuahang_id" json:"cuahang_id"`
	Phong        *string `boil:"phong" json:"phong"`
	Amount       string  `boil:"amount" json:"amount"`
	Mobile       string  `boil:"mobile" json:"mobile"`
	Date         string  `boil:"date" json:"date"`
	Email        *string `boil:"email" json:"email"`
	StoreNote    *string `boil:"store_note" json:"store_note"`
	CustomerNote *string `boil:"customer_note" json:"customer_note"`
	ApprovedBy   *string `boil:"approved_by" json:"approved_by"`
	ApprovedDate *string `boil:"approved_date" json:"approved_date"`
	PickupBy     *string `boil:"pickup_by" json:"pickup_by"`
	PickupDate   *string `boil:"pickup_date" json:"pickup_date"`
	Reason       *string `boil:"reason" json:"reason"`
}