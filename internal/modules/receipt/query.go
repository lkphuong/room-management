package receipt

const (
	GET_RECEIPT = `
		SELECT
			a1.store_code,
			a2.revenue,
			a1.revenue_tmp
		FROM (
			SELECT
				HoaDon.cuahang_id AS store_code,
				sum(HoaDon_ChiTietHangHoa.total) AS revenue_tmp
			FROM
				HoaDon
				JOIN HoaDon_ChiTietHangHoa ON HoaDon.id = HoaDon_ChiTietHangHoa.bill_id
				JOIN Ca ON ca.id = HoaDon.ca_id
			WHERE
				ca.date = CONVERT(varchar, GETDATE (), 23)
				AND HoaDon.trangthai = 4
				AND HoaDon_ChiTietHangHoa.trangthai = 4
			GROUP BY
				HoaDon.cuahang_id) a1
			JOIN (
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
					HoaDon.cuahang_id) a2 ON a1.store_code = a2.store_code
	`

	SELECT_RECEIPT_BY_STORE = `
				SELECT
					a1.room_code,
					a2.revenue,
					a1.revenue_tmp
				FROM (
					SELECT
						HoaDon.room_id AS room_code,
						sum(HoaDon_ChiTietHangHoa.total) AS revenue_tmp
					FROM
						HoaDon
						JOIN HoaDon_ChiTietHangHoa ON HoaDon.id = HoaDon_ChiTietHangHoa.bill_id
						JOIN Ca ON ca.id = HoaDon.ca_id
					WHERE
						ca.date = CONVERT(varchar, GETDATE (), 23)
						AND HoaDon.trangthai = 4
						AND HoaDon_ChiTietHangHoa.trangthai = 4
						AND HoaDon.cuahang_id = '%s'
					GROUP BY
						HoaDon.room_id) a1
					JOIN (
						SELECT
							HoaDon.room_id AS room_code,
							ISNULL(sum(HoaDon.total + HoaDon.phuphi + HoaDon.vat), 0) AS revenue
						FROM
							HoaDon
							JOIN Ca ON Ca.id = HoaDon.ca_id
						WHERE
							ca.date = CONVERT(varchar, GETDATE (), 23)
							AND HoaDon.trangthai = 4
							AND HoaDon.cuahang_id = '%s'
						GROUP BY
							HoaDon.room_id) a2 ON a1.room_code = a2.room_code
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
