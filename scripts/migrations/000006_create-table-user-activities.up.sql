CREATE TABLE IF NOT EXISTS user_activities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    post_id INT NOT NULL,
    user_id BIGINT NOT NULL,
    is_liked BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,

    CONSTRAINT fk_user_activities_post_id FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_user_activities_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);