-- Создание таблицы roles
CREATE TABLE roles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    role_name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

-- Создание таблицы users
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- Создание таблицы api_keys
CREATE TABLE api_keys (
    id INT AUTO_INCREMENT PRIMARY KEY,
    api_key VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    revoked BOOLEAN DEFAULT FALSE
);

-- Таблица сенсоров
CREATE TABLE sensors (
    id INT AUTO_INCREMENT PRIMARY KEY,  
    sensor_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    hostname VARCHAR(255) NOT NULL,  
    os_version VARCHAR(255) NOT NULL,          
    sensor_type VARCHAR(50) NOT NULL,             
    agent_version VARCHAR(50),          
    created_at TIMESTAMP DEFAULT NOW() 
);

-- Добавляем роль admin
INSERT INTO roles (role_name, description) 
VALUES ('admin', 'Администратор с полными правами доступа')
ON DUPLICATE KEY UPDATE role_name = 'admin';

-- Добавляем пользователя admin с ролью admin
INSERT INTO users (username, password, role_id) 
VALUES ('admin', '$2a$10$B.ITVXGQjhdW4QupKlwkfOrJz0QLKlhJt8pSLKBLqIN0pxsHoRSSK', (SELECT id FROM roles WHERE role_name = 'admin'))
ON DUPLICATE KEY UPDATE username = 'admin';
