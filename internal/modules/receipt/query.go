package receipt

const (
	GET_RECEIPT = `
		SELECT
			HoaDon.cuahang_id AS store_code,
			ISNULL(sum(HoaDon.total + HoaDon.phuphi + HoaDon.vat), 0) AS revenue
		FROM
			HoaDon
			JOIN Ca ON Ca.id = HoaDon.ca_id
		WHERE
			ca.date = CONVERT(varchar, GETDATE (), 23)
			AND HoaDon.trangthai = 4
		GROUP BY
			HoaDon.cuahang_id
	`

	SELECT_RECEIPT_BY_STORE = `
		SELECT
			HoaDon.room_id AS room_code,
			ISNULL(sum(HoaDon.total + HoaDon.phuphi + HoaDon.vat), 0) AS revenue
		FROM
			HoaDon
			JOIN Ca ON Ca.id = HoaDon.ca_id
		WHERE
			ca.date = cast(getdate () AS date)
			AND HoaDon.trangthai = 4
			AND HoaDon.cuahang_id = '%s'
		GROUP BY
			HoaDon.room_id
	`

	BILL_DETAIL = `
		SELECT
			goods_id AS id,
			goods_name AS name,
			quantity,
			price,
			total,
			category_id
		FROM
			HoaDon_ChiTietHangHoa
		WHERE
			bill_id = (
				SELECT
					top 1 HoaDon.id
				FROM
					HoaDon
					JOIN Ca ON Ca.id = HoaDon.ca_id
				WHERE
					ca.date = cast(getdate () AS date)
					AND HoaDon.trangthai = 4
					AND HoaDon.cuahang_id = '%s'
					AND hoadon.finish IS NULL
					AND hoadon.closed = 0
					AND room_id = '%s')
				AND trangthai = 4
	`
)
