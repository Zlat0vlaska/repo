-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS grades
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(20) NOT NULL,
    description VARCHAR(50) NOT NULL
);

INSERT INTO grades(name, description)
VALUES ('junior', 'Начинающий специалист'),
       ('middle', 'Специалист среднего уровня'),
       ('senior', 'Старший специалист');

CREATE TABLE IF NOT EXISTS grades_vacancies
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    grade_id   UUID REFERENCES grades (id),
    vacancy_id UUID REFERENCES vacancies (id)
);

INSERT INTO grades_vacancies (grade_id, vacancy_id)
VALUES ((SELECT id FROM grades WHERE name = 'senior'), (SELECT id FROM vacancies WHERE job_tittle = 'Senior Golang Developer')),
       ((SELECT id FROM grades WHERE name = 'middle'), (SELECT id FROM vacancies WHERE job_tittle = 'Middle React Developer')),
       ((SELECT id FROM grades WHERE name = 'junior'), (SELECT id FROM vacancies WHERE job_tittle = 'Junior Golang Developer')),
       ((SELECT id FROM grades WHERE name = 'senior'), (SELECT id FROM vacancies WHERE job_tittle = 'Senior DevOps Engineer')),
       ((SELECT id FROM grades WHERE name = 'junior'), (SELECT id FROM vacancies WHERE job_tittle = 'Junior React Developer')),
       ((SELECT id FROM grades WHERE name = 'middle'), (SELECT id FROM vacancies WHERE job_tittle = 'Middle System Analyst')),
       ((SELECT id FROM grades WHERE name = 'junior'), (SELECT id FROM vacancies WHERE job_tittle = 'Junior DevOps Engineer')),
       ((SELECT id FROM grades WHERE name = 'junior'), (SELECT id FROM vacancies WHERE job_tittle = 'Junior DB Administrator')),
       ((SELECT id FROM grades WHERE name = 'senior'), (SELECT id FROM vacancies WHERE job_tittle = 'Senior Project Manager')),
       ((SELECT id FROM grades WHERE name = 'junior'), (SELECT id FROM vacancies WHERE job_tittle = 'Junior System Analyst')),
       ((SELECT id FROM grades WHERE name = 'middle'), (SELECT id FROM vacancies WHERE job_tittle = 'Middle Golang Developer')),
       ((SELECT id FROM grades WHERE name = 'middle'), (SELECT id FROM vacancies WHERE job_tittle = 'Middle DevOps Engineer'));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS grades_vacancies;
DROP TABLE IF EXISTS grades;
-- +goose StatementEnd
