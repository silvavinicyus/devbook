package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO users(name, nick, email, password) VALUES(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	userInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(userInsertedId), nil
}

func (repository users) FindAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) Find(userId uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", userId,
	)

	if erro != nil {
		return models.User{}, erro
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository users) Update(userId uint64, user models.User) error {
	statement, erro := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, userId); erro != nil {
		return erro
	}

	return nil
}

func (repository users) Delete(userId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userId); erro != nil {
		return erro
	}

	return nil
}

func (repository users) FindByEmail(email string) (models.User, error) {
	line, erro := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?", email,
	)

	if erro != nil {
		return models.User{}, erro
	}

	defer line.Close()

	var user models.User

	if line.Next() {
		if erro := line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}
