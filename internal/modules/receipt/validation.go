package receipt

func (p ReceiptDetailParam) Validate() error {
	return validation.Struct(p)
}
