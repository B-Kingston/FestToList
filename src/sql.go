package main

//
//func insert(db *sql.DB) {
//
//	sqlStatement := `
//	INSERT INTO queries (img_loc, query_time)
//	VALUES ($1, $2)
//	RETURNING id`
//	id := 0
//	err := db.QueryRow(sqlStatement, "/Users/bailee/Documents/Programming/FestToList/src", "1999-01-08 04:05:06").Scan(&id)
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("New record ID is:", id)
//}
