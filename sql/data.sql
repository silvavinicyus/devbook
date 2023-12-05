INSERT INTO
  users (name, nick, email, password)
VALUES
  (
    "user1",
    "user_1",
    "user1@gmail.com",
    "$2a$10$GzBkIJ0EIjrJ5hR7in284OpPAUhhqvaTj0L97torOzFPx/SBRX0XG"
  ),
    (
    "user2",
    "user_2",
    "user2@gmail.com",
    "$2a$10$GzBkIJ0EIjrJ5hR7in284OpPAUhhqvaTj0L97torOzFPx/SBRX0XG"
  )
  ,  (
    "user3",
    "user_3",
    "user3@gmail.com",
    "$2a$10$GzBkIJ0EIjrJ5hR7in284OpPAUhhqvaTj0L97torOzFPx/SBRX0XG"
  );

INSERT INTO
  followers (user_id, follower_id)
VALUES
  (1, 2),
  (3, 1),
  (1, 3);

INSERT INTO 
  posts (title, content, creator_id)
VALUES
  ("Publicação do usuário 1", "Essa é a publicação do usuário 1", 1),
  ("Publicação do usuário 2", "Essa é a publicação do usuário 2", 2),
  ("Publicação do usuário 3", "Essa é a publicação do usuário 3", 3);