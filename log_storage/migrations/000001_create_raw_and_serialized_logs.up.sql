CREATE TABLE IF NOT EXISTS raw_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log_source VARCHAR(255) NOT NULL,  -- Источник лога
    log_string TEXT NOT NULL,          -- Содержимое лога в виде строки
    created_at_system DATETIME NOT NULL,      -- Время создания лога в системе
    sensor_id BIGINT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS serialized_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log_source VARCHAR(255) NOT NULL,  -- Источник лога
    log_serialized JSON NOT NULL,       -- Сериализованные данные лога (JSON строка)
    created_at_system DATETIME NOT NULL,      -- Время создания лога
    sensor_id BIGINT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);