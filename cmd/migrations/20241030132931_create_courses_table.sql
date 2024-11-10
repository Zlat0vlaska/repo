-- +goose Up
-- +goose StatementBegin
CREATE TABLE courses
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description VARCHAR(30) NOT NULL
);
CREATE TABLE resumes_courses
(
    id UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    resume_id  UUID NOT NULL,
    courses_id UUID NOT NULL,
    FOREIGN KEY (resume_id) REFERENCES resumes(id),
    FOREIGN KEY (courses_id) REFERENCES courses(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes_courses;
DROP TABLE IF EXISTS courses;
-- +goose StatementEnd
