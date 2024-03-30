package models

import "github.com/filipesiota/studygram/db"

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Biography string `json:"biography"`
	IsActive  bool   `json:"isActive"`
}

func Create(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql :=
		`INSERT INTO user
			(name, username, email, biography)
		VALUES
			($1, $2, $3, $4)
		RETURNING id`

	err = conn.QueryRow(sql, user.Name, user.Username, user.Email, user.Biography).Scan(&id)

	return
}

func Get(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql :=
		`SELECT
			id,
			name,
			username,
			email,
			biography,
			is_active AS isActive
		FROM user
		WHERE
			id = $1`

	row := conn.QueryRow(sql, id)

	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Biography, &user.IsActive)

	return
}

func GetAll() (users []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql :=
		`SELECT
			id,
			name,
			username,
			email,
			biography,
			is_active AS isActive
		FROM user`

	rows, err := conn.Query(sql)

	if err != nil {
		return
	}

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Biography, &user.IsActive)

		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return
}

func Update(user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql :=
		`UPDATE user
		SET
			name = $1
			username = $2
			email = $3
			biography = $4
		WHERE
			id = $5`

	res, err := conn.Exec(sql, user.Name, user.Username, user.Email, user.Biography, user.ID)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql :=
		`UPDATE user
		SET
			is_active = 0
		WHERE
			id = $5`

	res, err := conn.Exec(sql, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
