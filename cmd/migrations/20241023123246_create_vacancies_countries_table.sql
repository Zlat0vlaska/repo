-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS countries
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL
);

INSERT INTO countries (name) VALUES ('Россия');

CREATE TABLE IF NOT EXISTS regions
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    country_id UUID REFERENCES countries (id),
    name       VARCHAR(50) NOT NULL
);

INSERT INTO regions (country_id, name)
VALUES ((SELECT id FROM countries WHERE name = 'Россия'), 'Москва'),
       ((SELECT id FROM countries WHERE name = 'Россия'), 'Санкт-Петербург');

CREATE TABLE IF NOT EXISTS cities
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    region_id  UUID REFERENCES regions (id),
    name       VARCHAR(50) NOT NULL
);

INSERT INTO cities (region_id, name)
VALUES ((SELECT id FROM regions WHERE name = 'Москва'), 'Москва'),
       ((SELECT id FROM regions WHERE name = 'Санкт-Петербург'), 'Санкт-Петербург');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS regions;
DROP TABLE IF EXISTS countries;
-- +goose StatementEnd
