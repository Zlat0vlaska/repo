-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS candidates
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(30) NOT NULL,
    gender          VARCHAR(10) NOT NULL,
    birth_date      DATE        NOT NULL,
    registration    VARCHAR(70) NOT NULL,
    create_date     DATE        NOT NULL,
    hr_id           UUID      NOT NULL,
    FOREIGN KEY (hr_id) REFERENCES hrs (id)
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS candidates;
-- +goose StatementEnd