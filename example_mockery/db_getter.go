package example

func getFromDB(db DB) string {
	return db.Get("ice cream")
}
