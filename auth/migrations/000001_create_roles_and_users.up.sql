-- Создание таблицы roles
CREATE TABLE IF NOT EXISTS roles(
    id INT AUTO_INCREMENT PRIMARY KEY,
    role_name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

-- Создание таблицы users
CREATE TABLE IF NOT EXISTS users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- Добавляем роль admin
INSERT INTO roles (role_name, description) 
VALUES ('admin', 'Администратор с полными правами доступа')
ON DUPLICATE KEY UPDATE role_name = 'admin';

-- Добавляем пользователя admin с ролью admin
INSERT INTO users (username, password, role_id) 
VALUES ('admin', '$2a$10$B.ITVXGQjhdW4QupKlwkfOrJz0QLKlhJt8pSLKBLqIN0pxsHoRSSK', (SELECT id FROM roles WHERE role_name = 'admin'))
ON DUPLICATE KEY UPDATE username = 'admin';
