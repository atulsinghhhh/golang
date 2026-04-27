-- migrate:up
CREATE TABLE IF NOT EXISTS post_likes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_postlikes_post
        FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_postlikes_user
        FOREIGN KEY (user_id) REFERENCES users(id),

    UNIQUE (post_id, user_id)
);

-- migrate:down
DROP TABLE IF EXISTS post_likes;