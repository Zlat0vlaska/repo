-- +goose Up
-- +goose StatementBegin
CREATE TABLE educations
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description VARCHAR(30) NOT NULL
);
CREATE TABLE resumes_educations
(
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resume_id    UUID NOT NULL,
    education_id UUID REFERENCES educations (id),
    FOREIGN KEY (resume_id) REFERENCES resumes(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes_educations;
DROP TABLE IF EXISTS educations;
-- +goose StatementEnd
