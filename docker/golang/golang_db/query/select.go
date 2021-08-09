package query

import (
	"database/sql"
	"fmt"
)

type Account struct {
	id               int
	email            string
	phone            string
	username         string
	password         string
	create_at        int
	create_ip_at     string
	last_login_at    int
	last_login_ip_at string
	login_times      int
	status           int
}

func Select(db *sql.DB, input int) string {
	account := new(Account)
	row := db.QueryRow("select * from account_user where id=?", input)
	if err := row.Scan(
		&account.id,
		&account.email,
		&account.phone,
		&account.username,
		&account.password,
		&account.create_at,
		&account.create_ip_at,
		&account.last_login_at,
		&account.last_login_ip_at,
		&account.login_times,
		&account.status,
	); err != nil {
		fmt.Printf("Fail: ", err)
		return "Fail"
	}
	fmt.Println("Success: ", *account)

	return account.username

}
