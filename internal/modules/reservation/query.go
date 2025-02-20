package reservation

const (
	SELECT_RESERVATION = `
	SELECT 
		id, 
		customer_name,
		amount, 
		mobile, 
		date, 
		hour,
		status
	FROM Reservation 
	WHERE 
		(customer_name LIKE '%%' + '%s' + '%%' OR mobile LIKE '%%' + '%s' + '%%')
			AND date >= CAST(GETDATE() AS DATE)
			AND date < DATEADD(DAY, 1, CAST(GETDATE() AS DATE))
			AND trangthai = 4
		ORDER BY 
			CASE 
				WHEN date >= CAST(GETDATE() AS DATE) THEN 1
				ELSE 2
			END ASC,
			date ASC;
	`
)