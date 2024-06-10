package store

const (
	SELECT_STORES = `
		SELECT
            store_id,
            store_name,
            count(store_id) AS room_count
        FROM
            store
            JOIN phong ON (store_id = cuahang_id AND phong.trangthai = 4)
		WHERE store_id NOT IN ('30','2','26')
        GROUP BY
            store_id,
            store_name
	`
)
