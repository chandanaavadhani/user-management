package userDB

import (
	"github.com/chandanaavadhani/usermanagement/models"
	_ "github.com/go-sql-driver/mysql"
)

func Insertdetails(user models.User) error {

	db, err := Dbconnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Execute the query
	query, err := db.Prepare(`INSERT INTO users VALUES(?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(user.Fullname, user.Contactno, user.Email, user.Username, user.PWord)
	if err != nil {
		return err
	}
	return err
}

func Updatepassword(user models.User) error {
	db, err := Dbconnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Query(`Update users.users set PWord = ? where UserName = ? `, user.NewPassword, user.Username)
	if err != nil {
		return err
	}
	return err
}

func Deleteuser(username string) error {
	db, err := Dbconnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query, err := db.Query(`DELETE FROM users.users where UserName = ?`, username)
	if err != nil {
		return err
	}
	defer query.Close()
	return err
}

func Validationsquery(user models.User) (int, string, error) {

	var count int
	var password string
	db, err := Dbconnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query, err := db.Query(`select count(*), PWord from users.users where UserName = ? Group By UserName`, user.Username)
	if err != nil {
		return 0, "", err
	}
	defer query.Close()

	for query.Next() {
		if err := query.Scan(&count, &password); err != nil {
			return 0, "", err
		}
	}
	return count, password, nil
}

func StringValidations(user string) (int, string, error) {

	var count int
	var password string
	db, err := Dbconnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query, err := db.Query(`select count(*), PWord from users.users where UserName = ? Group By UserName`, user)
	if err != nil {
		return 0, "", err
	}
	defer query.Close()

	for query.Next() {
		if err := query.Scan(&count, &password); err != nil {
			return 0, "", err
		}
	}
	return count, password, nil
}
