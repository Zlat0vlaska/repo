-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS experiences
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS experiences;
-- +goose StatementEnd
