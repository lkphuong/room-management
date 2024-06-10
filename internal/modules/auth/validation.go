package auth

func (p LoginParam) Validate() error {
	return validation.Struct(p)
}
