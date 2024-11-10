-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS communication (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) NOT NULL,
    vacancy_id UUID NOT NULL REFERENCES vacancies(id),
    resume_id UUID NOT NULL REFERENCES resumes(id)
);

CREATE TABLE IF NOT EXISTS communication_status (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status_name VARCHAR(30) NOT NULL,
    communication_id UUID REFERENCES communication (id)
);

INSERT INTO communication_status (status_name)
VALUES ('Откликнут'), ('В работе'), ('Отклонен');

CREATE TABLE IF NOT EXISTS change_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    communication_status_id UUID REFERENCES communication_status (communication_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS change_history;
DROP TABLE IF EXISTS communication_status;
DROP TABLE IF EXISTS communication;

-- +goose StatementEnd