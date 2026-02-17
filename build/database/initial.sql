CREATE DATABASE IF NOT EXISTS amsvault;

USE amsvault;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS stories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    source VARCHAR(255),
    description TEXT,
    season INT,
    episode INT,
    volume INT,
    chapter INT,
    status VARCHAR(255) NOT NULL,
    medium_image VARCHAR(255),
    large_image VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

/*
    1 - Crie o usuário explicitamente:
    CREATE USER 'amsvault'@'%' IDENTIFIED BY 'amsvaultPwd';

    2- Conceda os privilégios ao usuário:
    GRANT ALL PRIVILEGES ON *.* TO 'amsvault'@'%';
    FLUSH PRIVILEGES;
*/

INSERT INTO users (name, email, password) 
VALUES ('Greg', 'greg@hotmail.com', '$2a$10$EixZaYVK1fsbw1ZfbX3OXePaWxn96p36F8zNwZ4VGh5FZ4pGWEq9e');