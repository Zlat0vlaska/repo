-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS resumes_skills
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resume_id   UUID NOT NULL,
    skill_id    UUID NOT NULL,
    FOREIGN KEY (resume_id) REFERENCES resumes(id),
    FOREIGN KEY (skill_id) REFERENCES skills(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes_skills;
-- +goose StatementEnd