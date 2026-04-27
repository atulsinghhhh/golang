-- migrate:up
CREATE TABLE IF NOT EXISTS comment_likes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    comment_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_commentlikes_comment
        FOREIGN KEY (comment_id) REFERENCES comments(id),
    CONSTRAINT fk_commentlikes_user
        FOREIGN KEY (user_id) REFERENCES users(id),

    UNIQUE (comment_id, user_id)
);

-- migrate:down
DROP TABLE IF EXISTS comment_likes;