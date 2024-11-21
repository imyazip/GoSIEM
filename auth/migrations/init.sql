CREATE DATABASE IF NOT EXISTS auth_db;

CREATE TABLE IF NOT EXISTS auth_db.api_keys (
    id INT AUTO_INCREMENT PRIMARY KEY,
    key_value VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO auth_db.api_keys (key_value) VALUES ('your-api-key');