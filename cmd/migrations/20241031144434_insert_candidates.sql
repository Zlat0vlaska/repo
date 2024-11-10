-- +goose Up
-- +goose StatementBegin
INSERT INTO candidates (name, gender, birth_date, registration, create_date, hr_id)
VALUES
    ('Алексей Соколов', 'мужской', '1985-07-08', 'Казань', CURRENT_DATE, (SELECT id FROM hrs WHERE name = 'Анна Иванова')),
    ('Ольга Сергеева', 'женский', '1991-09-14', 'Самара', CURRENT_DATE, (SELECT id FROM hrs WHERE name = 'Анна Иванова')),

    ('Николай Зуев', 'мужской', '1987-02-21', 'Воронеж', CURRENT_DATE, (SELECT id FROM hrs WHERE name = 'Сергей Петров')),
    ('Елена Фомина', 'женский', '1993-06-30', 'Ростов-на-Дону', CURRENT_DATE, (SELECT id FROM hrs WHERE name = 'Сергей Петров')),

    ('Василий Кравцов', 'мужской', '1995-12-05', 'Уфа', CURRENT_DATE, (SELECT id FROM hrs WHERE name = 'Мария Сидорова')),
    ('Ирина Щербакова', 'женский', '1990-10-18', 'Красноярск', CURRENT_DATE, (SELECT id FROM hrs WHERE name = 'Мария Сидорова'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM candidates
WHERE name IN (
               'Алексей Соколов',
               'Ольга Сергеева',
               'Николай Зуев',
               'Елена Фомина',
               'Василий Кравцов',
               'Ирина Щербакова'
    );
-- +goose StatementEnd
