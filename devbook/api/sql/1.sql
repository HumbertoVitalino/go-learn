CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id INT auto_increment PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL unique,
    email VARCHAR(50) NOT NULL unique,
    passkeyword VARCHAR(20) not null unique,
    createdAt TIMESTAMP default current_timestamp()
) ENGINE=INNODB;