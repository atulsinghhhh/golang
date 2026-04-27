-- migrate:up
CREATE TABLE IF NOT EXISTS refresh_token (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    refresh_token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_refresh_user
        FOREIGN KEY (user_id) REFERENCES users(id)
);

-- migrate:down
DROP TABLE IF EXISTS refresh_token;