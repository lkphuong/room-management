package reservation

import "database/sql"

type ReservationQueryAll struct {
	KeyWork string `query:"keywork" json:"keywork" validate:"omitempty,max=255"`
}

type ReservationResponse struct {
	Id           string         `boil:"id" json:"id"`
	CustomerName sql.NullString `boil:"customer_name" json:"customer_name"`
	Amount       string         `boil:"amount" json:"amount"`
	Mobile       string         `boil:"mobile" json:"mobile"`
	Date       	 string         `boil:"date" json:"date"`
	Hour       	 string         `boil:"hour" json:"hour"`
	Status       int16          `boil:"status" json:"status"`
}

type ReservationResponseData struct {
	Id           string  `boil:"id" json:"id"`
	CustomerName *string `boil:"customer_name" json:"customer_name"`
	Amount       string  `boil:"amount" json:"amount"`
	Mobile       string  `boil:"mobile" json:"mobile"`
	Date         string  `boil:"date" json:"date"`
	Hour         string  `boil:"hour" json:"hour"`
	Status       int16   `boil:"status" json:"status"`
}