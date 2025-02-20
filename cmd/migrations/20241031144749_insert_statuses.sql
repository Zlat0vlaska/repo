-- +goose Up
-- +goose StatementBegin
INSERT INTO statuses (name, description)
VALUES
    -- Статусы для вакансий
    ('в архиве', 'Вакансия закрыта и перемещена в архив'),
    ('активная', 'Вакансия доступна для отклика кандидатов'),

    -- Статусы для резюме
    ('новое', 'Резюме только что добавлено и ожидает обработки'),
    ('на рассмотрении', 'Резюме рассматривается HR-отделом'),
    ('одобрено', 'Резюме одобрено и готово к следующему этапу'),
    ('отклонено', 'Резюме отклонено и не будет продвигаться дальше'),
    ('в архиве', 'Резюме перемещено в архив и больше не актуально');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM statuses
WHERE name IN ('в архиве', 'активная', 'новое', 'на рассмотрении', 'одобрено', 'отклонено');
-- +goose StatementEnd
