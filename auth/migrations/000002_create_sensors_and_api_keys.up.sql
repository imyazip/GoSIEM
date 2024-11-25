-- Создание таблицы api_keys
CREATE TABLE IF NOT EXISTS api_keys(
    id INT AUTO_INCREMENT PRIMARY KEY,
    api_key VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    revoked BOOLEAN DEFAULT FALSE
);

-- Таблица сенсоров
CREATE TABLE IF NOT EXISTS sensors(
    id INT AUTO_INCREMENT PRIMARY KEY,  
    sensor_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    hostname VARCHAR(255) NOT NULL,  
    os_version VARCHAR(255) NOT NULL,          
    sensor_type VARCHAR(50) NOT NULL,             
    agent_version VARCHAR(50),          
    created_at TIMESTAMP DEFAULT NOW() 
);

INSERT INTO api_keys (api_key) 
VALUES ('api_key')
ON DUPLICATE KEY UPDATE api_key = 'api_key';


