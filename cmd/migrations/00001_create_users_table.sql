-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users
(
    id INT GENERATED ALWAYS AS IDENTITY,
    email VARCHAR(128) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    surname VARCHAR(50),
    directory_type VARCHAR(50) CHECK (directory_type IN ('USER', 'APPLICANT'))
);

CREATE UNIQUE INDEX users_email_idx ON users(email);

INSERT INTO users (email, first_name, last_name, directory_type)
VALUES
('alice.smith@example.com', 'Алиса', 'Иванова', 'USER'),
('bob.jones@example.com', 'Борис', 'Смирнов', 'USER'),
('charlie.brown@example.com', 'Виктор', 'Кузнецов', 'USER'),
('daniel.smith@example.com', 'Даниил', 'Иванов', 'USER'),
('eve.johnson@example.com', 'Ева', 'Петрова', 'APPLICANT'),
('frank.johnson@example.com', 'Фёдор', 'Петров', 'APPLICANT'),
('grace.evans@example.com', 'Галина', 'Новикова', 'APPLICANT'),
('hannah.evans@example.com', 'Галина', 'Николаева', 'APPLICANT'),
('isabel.jones@example.com', 'Ирина', 'Смирнова', 'APPLICANT'),
('jack.sanders@example.com', 'Иван', 'Сидоров', 'APPLICANT');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd