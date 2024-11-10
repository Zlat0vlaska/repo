-- +goose Up
-- +goose StatementBegin
INSERT INTO resumes (profession, schedule, citizenship, business_trips_readiness, permission, salary, relocation_readiness, candidate_id, status_id, grade_id, country_id)
VALUES
    ('Менеджер проектов', 'удаленная работа', 'Россия', 'не готов', 'разрешение на работу', 90000.0, 'не готов',
     (SELECT id FROM candidates WHERE name = 'Алексей Соколов'), (SELECT id FROM statuses WHERE name = 'на рассмотрении'), (SELECT id FROM grades WHERE name = 'junior'), (SELECT id FROM countries WHERE name = 'Россия')),

    ('Аналитик данных', 'гибкий график', 'Россия', 'готовность', 'разрешение на работу', 100000.0, 'готовность',
     (SELECT id FROM candidates WHERE name = 'Ольга Сергеева'), (SELECT id FROM statuses WHERE name = 'одобрено'), (SELECT id FROM grades WHERE name = 'middle'), (SELECT id FROM countries WHERE name = 'Россия')),

    ('Финансовый аналитик', 'удаленная работа', 'Беларусь', 'готовность', 'разрешение на работу', 110000.0, 'не готов',
     (SELECT id FROM candidates WHERE name = 'Николай Зуев'), (SELECT id FROM statuses WHERE name = 'отклонено'), (SELECT id FROM grades WHERE name = 'senior'), (SELECT id FROM countries WHERE name = 'Россия')),

    ('Дизайнер', 'гибкий график', 'Россия', 'готовность', 'разрешение на работу', 70000.0, 'готовность',
     (SELECT id FROM candidates WHERE name = 'Елена Фомина'), (SELECT id FROM statuses WHERE name = 'одобрено'), (SELECT id FROM grades WHERE name = 'senior'), (SELECT id FROM countries WHERE name = 'Россия')),

    ('Программист', 'полный день', 'Казахстан', 'не готов', 'разрешение на работу', 130000.0, 'готовность',
     (SELECT id FROM candidates WHERE name = 'Василий Кравцов'), (SELECT id FROM statuses WHERE name = 'новое'), (SELECT id FROM grades WHERE name = 'senior'), (SELECT id FROM countries WHERE name = 'Россия')),

    ('Аналитик', 'полный день', 'Россия', 'готовность', 'разрешение на работу', 95000.0, 'не готов',
     (SELECT id FROM candidates WHERE name = 'Ирина Щербакова'), (SELECT id FROM statuses WHERE name = 'на рассмотрении'), (SELECT id FROM grades WHERE name = 'senior'), (SELECT id FROM countries WHERE name = 'Россия'));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM resumes
WHERE candidate_id IN (
    SELECT id FROM candidates WHERE name IN (
                                             'Анна Кузнецова',
                                             'Алексей Соколов',
                                             'Ольга Сергеева',
                                             'Иван Смирнов',
                                             'Николай Зуев',
                                             'Елена Фомина',
                                             'Василий Кравцов',
                                             'Ирина Щербакова'
        )
);
-- +goose StatementEnd
