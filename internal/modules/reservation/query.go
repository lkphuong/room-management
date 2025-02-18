package reservation

const (
	SELECT_RESERVATION = `
	SELECT * 
	FROM Reservation 
	WHERE 
		(customer_name LIKE '%%' + '%s' + '%%' OR mobile LIKE '%%' + '%s' + '%%')
			AND Date >= CAST(GETDATE() AS DATE)
			AND Date < DATEADD(DAY, 1, CAST(GETDATE() AS DATE))
		ORDER BY 
			CASE 
				WHEN Date >= CAST(GETDATE() AS DATE) THEN 1
				ELSE 2
			END ASC,
			Date ASC;
	`
)