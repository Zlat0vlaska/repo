-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS grades_resumes
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    grade_id  UUID NOT NULL,
    resume_id UUID NOT NULL,
    FOREIGN KEY (resume_id) REFERENCES resumes(id),
    FOREIGN KEY (grade_id) REFERENCES grades(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS grades_resumes;
-- +goose StatementEnd
