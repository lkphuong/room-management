package auth

const (
	LOGIN = `select id, code, name from NhanVien2 where code = '%s' and password = '%s' and trangthai = 4 and is_locked = 0`
)
