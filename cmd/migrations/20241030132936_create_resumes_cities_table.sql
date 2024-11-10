-- +goose Up
-- +goose StatementBegin

CREATE TABLE resumes_cities
(
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resume_id    UUID NOT NULL,
    city_id      UUID NOT NULL,
    FOREIGN KEY (resume_id) REFERENCES resumes(id),
    FOREIGN KEY (city_id) REFERENCES cities(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes_cities;
-- +goose StatementEnd
