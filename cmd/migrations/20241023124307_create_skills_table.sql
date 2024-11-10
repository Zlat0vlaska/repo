-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS skills
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) NOT NULL
);

INSERT INTO skills (name)
VALUES ('jira'),
       ('uml'),
       ('scrum'),
       ('html'),
       ('jsp'),
       ('mongodb'),
       ('redis'),
       ('bootstrap'),
       ('junit'),
       ('oracle'),
       ('symfony'),
       ('postman'),
       ('test'),
       ('vue'),
       ('soap'),
       ('autocad'),
       ('ajax'),
       ('string'),
       ('postgresql'),
       ('mvc'),
       ('dbeaver'),
       ('rest'),
       ('json'),
       ('typescript'),
       ('nats'),
       ('rabbitmq'),
       ('python'),
       ('ruby'),
       ('docker'),
       ('openapi'),
       ('javascript'),
       ('git'),
       ('agile'),
       ('webpack'),
       ('php'),
       ('css3'),
       ('html5'),
       ('laravel'),
       ('ui'),
       ('sql'),
       ('css'),
       ('jpa'),
       ('веб-программирование'),
       ('jenkins'),
       ('swagger'),
       ('figma'),
       ('golang'),
       ('test2'),
       ('kafka'),
       ('jdbc'),
       ('pug'),
       ('c#'),
       ('jquery'),
       ('redux'),
       ('reactjs'),
       ('gulp'),
       ('vpn'),
       ('linux'),
       ('java'),
       ('c++'),
       ('oop'),
       ('yii'),
       ('confluence'),
       ('xml');

CREATE TABLE IF NOT EXISTS vacancies_skills
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vacancy_id UUID REFERENCES vacancies (id),
    skill_id   UUID REFERENCES skills (id)
);

INSERT INTO vacancies_skills (vacancy_id, skill_id)
VALUES
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior Golang Developer'), (SELECT id FROM skills WHERE name = 'golang')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior Golang Developer'), (SELECT id FROM skills WHERE name = 'agile')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle React Developer'), (SELECT id FROM skills WHERE name = 'reactjs')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle React Developer'), (SELECT id FROM skills WHERE name = 'css')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior Golang Developer'), (SELECT id FROM skills WHERE name = 'golang')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior Golang Developer'), (SELECT id FROM skills WHERE name = 'docker')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior DevOps Engineer'), (SELECT id FROM skills WHERE name = 'docker')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior DevOps Engineer'), (SELECT id FROM skills WHERE name = 'agile')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior React Developer'), (SELECT id FROM skills WHERE name = 'reactjs')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior React Developer'), (SELECT id FROM skills WHERE name = 'css')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle System Analyst'), (SELECT id FROM skills WHERE name = 'autocad')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle System Analyst'), (SELECT id FROM skills WHERE name = 'uml')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior DevOps Engineer'), (SELECT id FROM skills WHERE name = 'docker')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior DevOps Engineer'), (SELECT id FROM skills WHERE name = 'postgresql')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior DB Administrator'), (SELECT id FROM skills WHERE name = 'postgresql')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior DB Administrator'), (SELECT id FROM skills WHERE name = 'docker')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior Project Manager'), (SELECT id FROM skills WHERE name = 'agile')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior Project Manager'), (SELECT id FROM skills WHERE name = 'css')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior System Analyst'), (SELECT id FROM skills WHERE name = 'autocad')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior System Analyst'), (SELECT id FROM skills WHERE name = 'postgresql')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle Golang Developer'), (SELECT id FROM skills WHERE name = 'golang')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle Golang Developer'), (SELECT id FROM skills WHERE name = 'docker')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle DevOps Engineer'), (SELECT id FROM skills WHERE name = 'docker')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle DevOps Engineer'), (SELECT id FROM skills WHERE name = 'postgresql'));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vacancies_skills;
DROP TABLE IF EXISTS skills;
-- +goose StatementEnd
