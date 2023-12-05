package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

func PostRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (repository posts) Create(post models.Post) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO posts(title, content, creator_id) VALUES(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(post.Title, post.Content, post.CreatorId)
	if erro != nil {
		return 0, erro
	}

	lastId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastId), nil
}

func (repository posts) FindById(postId uint64) (models.Post, error) {
	line, erro := repository.db.Query(`
			SELECT p.*, u.nick FROM posts p INNER JOIN users u ON u.id = p.creator_id WHERE p.id = ?
		`, postId)

	if erro != nil {
		return models.Post{}, erro
	}
	defer line.Close()

	var post models.Post

	if line.Next() {
		if erro = line.Scan(&post.ID, &post.Title, &post.Content, &post.CreatorId, &post.Likes, &post.CreatedAt, &post.CreatorNick); erro != nil {
			return models.Post{}, erro
		}
	}

	return post, nil
}

func (repository posts) Find(userId uint64) ([]models.Post, error) {
	lines, erro := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON u.id = p.creator_id
		LEFT JOIN followers s on p.creator_id = s.user_id
		WHERE u.id = ? OR s.follower_id = ?
		ORDER BY p.createdAt DESC
	`, userId, userId)
	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if erro := lines.Scan(&post.ID, &post.Title, &post.Content, &post.CreatorId, &post.Likes, &post.CreatedAt, &post.CreatorNick); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository posts) Update(postId uint64, post models.Post) error {
	statement, erro := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(post.Title, post.Content, postId); erro != nil {
		return erro
	}

	return nil
}

func (repository posts) Delete(postId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postId); erro != nil {
		return erro
	}

	return nil
}

func (repository posts) FindByUser(userId uint64) ([]models.Post, error) {
	lines, erro := repository.db.Query(`
		SELECT p.*, u.nick 
		FROM posts p 
		JOIN users u ON u.id = p.creator_id
		WHERE p.creator_id = ?
	`, userId)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if erro = lines.Scan(&post.ID, &post.Title, &post.Content, &post.CreatorId, &post.Likes, &post.CreatedAt, &post.CreatorNick); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository posts) Like(postId uint64) error {
	statement, erro := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postId); erro != nil {
		return erro
	}

	return nil
}

func (repository posts) Unlike(postId uint64) error {
	statement, erro := repository.db.Prepare(`
		UPDATE posts 
		SET likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		WHERE id = ?
	`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postId); erro != nil {
		return erro
	}

	return nil
}
