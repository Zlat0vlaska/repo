-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS hrs
(
    id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name   VARCHAR(30) NOT NULL,
    email  VARCHAR(30) NOT NULL,
    number VARCHAR(12) NOT NULL
);

INSERT INTO hrs (name, email, number)
VALUES ('Анна Иванова', 'anna.ivanova@example.com', '79991234567'),
       ('Сергей Петров', 'sergey.petrov@example.com', '79991234568'),
       ('Мария Сидорова', 'maria.sidorova@example.com', '79991234569');

CREATE TABLE IF NOT EXISTS vacancies
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(50) NOT NULL,
    company     VARCHAR(20) NOT NULL,
    hr_id       UUID REFERENCES hrs (id),
    state       VARCHAR(20) NOT NULL,
    job_tittle  VARCHAR(30) NOT NULL,
    salary      INTEGER     NOT NULL,
    description VARCHAR(50) NOT NULL,
    date_create DATE        NOT NULL DEFAULT CURRENT_DATE,
    is_favorite BOOLEAN     NOT NULL DEFAULT false,
    country_id  UUID REFERENCES countries (id)
);

INSERT INTO vacancies (name, company, hr_id, country_id, state, job_tittle, salary, description)
VALUES
    ('Backend Developer', 'TechCorp', (SELECT id FROM hrs WHERE name = 'Анна Иванова'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Senior Golang Developer', 150000, 'Разработка серверной части на Go'),
    ('Frontend Developer', 'WebSolutions', (SELECT id FROM hrs WHERE name = 'Сергей Петров'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Middle React Developer', 120000, 'Разработка интерфейсов на ReactJS'),
    ('Backend Developer', 'TechCorp', (SELECT id FROM hrs WHERE name = 'Анна Иванова'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Junior Golang Developer', 100000, 'Работа с серверной частью на Go'),
    ('DevOps Engineer', 'CloudTech', (SELECT id FROM hrs WHERE name = 'Мария Сидорова'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Senior DevOps Engineer', 140000, 'Поддержка CI/CD процессов'),
    ('Frontend Developer', 'WebSolutions', (SELECT id FROM hrs WHERE name = 'Сергей Петров'), (SELECT id FROM countries WHERE name = 'Россия'), 'closed', 'Junior React Developer', 90000, 'Создание интерфейсов на ReactJS'),
    ('System Analyst', 'DataAnalytica', (SELECT id FROM hrs WHERE name = 'Анна Иванова'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Middle System Analyst', 130000, 'Анализ системных требований'),
    ('DevOps Engineer', 'CloudTech', (SELECT id FROM hrs WHERE name = 'Мария Сидорова'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Junior DevOps Engineer', 110000, 'Внедрение CI/CD процессов'),
    ('Database Administrator', 'DataSecure', (SELECT id FROM hrs WHERE name = 'Анна Иванова'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Junior DB Administrator', 90000, 'Управление базами данных'),
    ('Project Manager', 'InnovateGroup', (SELECT id FROM hrs WHERE name = 'Сергей Петров'), (SELECT id FROM countries WHERE name = 'Россия'), 'closed', 'Senior Project Manager', 150000, 'Управление проектами'),
    ('System Analyst', 'DataAnalytica', (SELECT id FROM hrs WHERE name = 'Анна Иванова'), (SELECT id FROM countries WHERE name = 'Россия'), 'open', 'Junior System Analyst', 100000, 'Поддержка анализа системных требований'),
    ('Backend Developer', 'TechCorp', (SELECT id FROM hrs WHERE name = 'Анна Иванова'), (SELECT id FROM countries WHERE name = 'Россия'), 'archived', 'Middle Golang Developer', 130000, 'Поддержка и доработка серверной части на Go'),
    ('DevOps Engineer', 'CloudTech', (SELECT id FROM hrs WHERE name = 'Мария Сидорова'), (SELECT id FROM countries WHERE name = 'Россия'), 'archived', 'Middle DevOps Engineer', 125000, 'Поддержка и внедрение CI/CD процессов');

--

CREATE TABLE IF NOT EXISTS vacancies_cities
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vacancy_id UUID REFERENCES vacancies (id),
    city_id    UUID REFERENCES cities (id)
);

INSERT INTO vacancies_cities (vacancy_id, city_id)
VALUES
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior Golang Developer'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle React Developer'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior Golang Developer'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior DevOps Engineer'), (SELECT id FROM cities WHERE name = 'Санкт-Петербург')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior React Developer'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle System Analyst'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior DevOps Engineer'), (SELECT id FROM cities WHERE name = 'Санкт-Петербург')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior DB Administrator'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Senior Project Manager'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Junior System Analyst'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle Golang Developer'), (SELECT id FROM cities WHERE name = 'Москва')),
    ((SELECT id FROM vacancies WHERE job_tittle = 'Middle DevOps Engineer'), (SELECT id FROM cities WHERE name = 'Санкт-Петербург'));


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vacancies_cities;
DROP TABLE IF EXISTS vacancies;
DROP TABLE IF EXISTS hrs;
-- +goose StatementEnd
