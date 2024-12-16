ALTER TABLE users 
ADD COLUMN username VARCHAR(255) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT unique_username UNIQUE (username);