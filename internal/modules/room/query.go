package room

const (
	SELECT_ROOMS = `
		SELECT
			CuaHang.code AS store_code,
			CuaHang.name AS store_name,
			phong.code AS room_code,
			ISNULL(t.start, '') AS start
		FROM (
			SELECT
				*
			FROM
				phong
			WHERE
				trangthai = 4) phong
			JOIN CuaHang ON phong.cuahang_id = CuaHang.code
			LEFT JOIN (
				SELECT
					hoadon.room_id, hoadon.cuahang_id, hoadon. [start]
				FROM
					hoadon
					JOIN Ca ON hoadon.ca_id = Ca.id
				WHERE
					ca.date = CONVERT(varchar, GETDATE (), 23)
					AND hoadon.finish IS NULL
					AND hoadon.closed = 0) t ON (phong.code = t.room_id
				AND phong.cuahang_id = t.cuahang_id)
		ORDER BY
			t.start DESC,
			CuaHang.code ASC
	`

	SELECT_ROOMS_BY_STORE = `
		SELECT
			CuaHang.code AS store_code,
			CuaHang.name AS store_name,
			phong.code AS room_code,
			t.start
		FROM (
			SELECT
				*
			FROM
				phong
			WHERE
				trangthai = 4) phong
			JOIN CuaHang ON phong.cuahang_id = CuaHang.code
			LEFT JOIN (
				SELECT
					hoadon.room_id, hoadon.cuahang_id, hoadon. [start]
				FROM
					hoadon
					JOIN Ca ON hoadon.ca_id = Ca.id
				WHERE
					ca.date = CONVERT(varchar, GETDATE (), 23)
					AND hoadon.finish IS NULL
					AND hoadon.closed = 0) t ON (phong.code = t.room_id
				AND phong.cuahang_id = t.cuahang_id)
		WHERE phong.cuahang_id = '%s'
		ORDER BY
			t.start DESC,
			CuaHang.code ASC
	`
)
