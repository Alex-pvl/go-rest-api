package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"rest-api/model"
)

type UserRepositoryImpl struct {
	db             *sql.DB
	UserRepository UserRepository
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) FindAll() []model.User {
	users := make([]model.User, 0)
	rows, err := ur.db.Query("SELECT id, login, password FROM users")
	if err != nil {
		return users
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Password); err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	return users
}

func (ur *UserRepositoryImpl) FindOneById(id uint64) (model.User, error) {
	row := ur.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	user := model.User{}
	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.Token)

	return user, err
}

func (ur *UserRepositoryImpl) FindOneByLogin(login string) (model.User, error) {
	row := ur.db.QueryRow("SELECT * FROM users WHERE login = $1", login)

	user := model.User{}
	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.Token)

	return user, err
}

func (ur *UserRepositoryImpl) Create(login, password string) (uint64, error) {
	query := "INSERT INTO users(login, password, token) VALUES ($1, $2, $3) RETURNING id"
	row := ur.db.QueryRow(query, login, password, "")

	var userId uint64
	err := row.Scan(&userId)

	return userId, err
}

func (ur *UserRepositoryImpl) DeleteById(id uint64) error {
	_, err := ur.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
