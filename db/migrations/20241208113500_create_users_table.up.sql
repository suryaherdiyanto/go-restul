
CREATE TABLE users(
    id int AUTO_INCREMENT,
    first_name varchar(49),
    last_name varchar(49) null,
	password varchar(100),
    email varchar(29),
    created_at timestamp,
    updated_at timestamp,

    PRIMARY KEY(id),
    UNIQUE(email)
);
