-- +goose Up
-- +goose StatementBegin

 
CREATE TABLE IF NOT EXISTS resumes
(
    id                          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    profession                  VARCHAR(50) NOT NULL,
    schedule                    VARCHAR(20) NOT NULL,
    citizenship                 VARCHAR(20) NOT NULL,
    business_trips_readiness    VARCHAR(20) NOT NULL,
    permission                  VARCHAR(20) NOT NULL,
    salary                      FLOAT       NOT NULL,
    relocation_readiness        VARCHAR(20) NOT NULL,
    candidate_id                UUID      NOT NULL,
    status_id                   UUID      NOT NULL,
    grade_id                    UUID      NOT NULL,
    country_id                  UUID      NOT NULL,
    FOREIGN KEY (candidate_id)  REFERENCES candidates(id),
    FOREIGN KEY (status_id)     REFERENCES statuses(id),
    FOREIGN KEY (grade_id)      REFERENCES grades(id),
    FOREIGN KEY (country_id)    REFERENCES countries(id)
);
-- +goose StatementEnd
 
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes;
-- +goose StatementEnd
