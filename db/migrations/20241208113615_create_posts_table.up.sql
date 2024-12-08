
CREATE TABLE posts(
    id int AUTO_INCREMENT,
    title varchar(100),
    category varchar(20),
    user_id int,
    content text,
    created_at timestamp,
    updated_at timestamp,

    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES `users`(id)
)