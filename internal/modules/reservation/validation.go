package reservation

func (p ReservationQueryAll) Validate() error {
	return validation.Struct(p)
}
