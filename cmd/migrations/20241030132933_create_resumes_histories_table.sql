-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS resumes_histories  
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(40) NOT NULL,
    description VARCHAR(40) NOT NULL,
    create_date DATE NOT NULL,
    id_resume   UUID NOT NULL,
    FOREIGN KEY (id_resume) REFERENCES resumes (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes_histories;
-- +goose StatementEnd
