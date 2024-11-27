CREATE TABLE IF NOT EXISTS raw_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log_source VARCHAR(255) NOT NULL,
    log_string TEXT NOT NULL,
    created_at_system DATETIME NOT NULL,
    sensor_id BIGINT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS serialized_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log_source VARCHAR(255) NOT NULL,
    log_serialized JSON NOT NULL, 
    created_at_system DATETIME NOT NULL,
    sensor_id BIGINT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
