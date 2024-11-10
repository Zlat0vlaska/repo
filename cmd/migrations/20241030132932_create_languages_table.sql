-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS languages 
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS resumes_languages  
(
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resume_id      UUID REFERENCES resumes (id),
    languages_id   UUID REFERENCES languages (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes_languages;
DROP TABLE IF EXISTS languages;
-- +goose StatementEnd
