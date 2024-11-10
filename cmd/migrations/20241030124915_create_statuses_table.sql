-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS statuses
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(20)     NOT NULL,
    description VARCHAR(200)    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS statuses;
-- +goose StatementEnd
