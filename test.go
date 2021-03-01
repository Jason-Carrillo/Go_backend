type Person struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int    `db:"age"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "codeup"
	dbName := "go_chal_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic("dbConn:" + err.Error())
	}
	fmt.Println("TEST")
	return db
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM person")
	ErrorCheck(err)
	per := Person{}
	var res []Person
	for selDB.Next() {
		var ID int
		var FirstName, LastName string
		var Age int

		err = selDB.Scan(&ID, &FirstName, &LastName, &Age)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		res = append(res, per)
		fmt.Println(res)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}