-- migrate:up

ALTER TABLE refresh_token ADD COLUMN expires_at TIMESTAMP NOT NULL AFTER refresh_token;


-- migrate:down

ALTER TABLE refresh_token DROP COLUMN expires_at;
