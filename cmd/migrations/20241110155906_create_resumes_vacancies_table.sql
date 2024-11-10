-- +goose Up
-- +goose StatementBegin
CREATE TABLE resume_vacancy (
  id UUID PRIMARY KEY,
  resume_id UUID REFERENCES resumes(id),
  vacancy_id UUID REFERENCES vacancies(id),
  status VARCHAR(255),
  resume_status VARCHAR(255),
  vacancy_status VARCHAR(255),
  notes VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resume_vacancy;
-- +goose StatementEnd
