package golang_db

import (
	"database/sql"
	"fmt"
)

type Account struct {
	id               int
	email            string
	phone            string
	username         string
	Password         string
	create_at        int
	create_ip_at     string
	last_login_at    int
	last_login_ip_at string
	login_times      int
	status           int
}

func SelectById(db *sql.DB, id int) (Account, error) {
	account := Account{}
	row := db.QueryRow("select * from account_user where id=?", id)
	if err := row.Scan(
		&account.id,
		&account.email,
		&account.phone,
		&account.username,
		&account.Password,
		&account.create_at,
		&account.create_ip_at,
		&account.last_login_at,
		&account.last_login_ip_at,
		&account.login_times,
		&account.status,
	); err != nil {
		fmt.Printf("Fail: ", err)
		return account, err
	}

	return account, nil

}
