CREATE TABLE IF NOT EXISTS raw_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log_source VARCHAR(255) NOT NULL,
    log_string TEXT NOT NULL,
    created_at_system DATETIME NOT NULL,
    sensor_id VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    read_flag BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS serialized_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log_source VARCHAR(255) NOT NULL,
    log_serialized BLOB NOT NULL, 
    created_at_system DATETIME NOT NULL,
    sensor_id VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    read_flag BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS security_events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log_id INT NOT NULL, 
    event_type VARCHAR(255) NOT NULL,
    event_description TEXT,
    detected_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    read_flag BOOLEAN DEFAULT FALSE --Используется для оповещений
    FOREIGN KEY(log_id) REFERENCES serialized_logs(id)
        ON DELETE CASCADE
)
