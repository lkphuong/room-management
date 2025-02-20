package auth

const (
	LOGIN = `
		SELECT
			NhanVien2.id,
			NhanVien2.code,
			NhanVien2.name,
			NhanVien_CuaHang.cuahang_id as store_id
		FROM
			NhanVien2
			JOIN NhanVien_CuaHang ON NhanVien2.code = NhanVien_CuaHang.user_id
		WHERE
			code = '%s'
			AND password = '%s'
			AND NhanVien2.trangthai = 4
			AND NhanVien_CuaHang.trangthai = 4
			AND is_locked = 0
	`
)
