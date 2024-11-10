-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS meetings
(
    id               UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name             VARCHAR(100)             NOT NULL,
    place            VARCHAR(100)             NOT NULL,
    comment          VARCHAR(1000),
    recipient_emails VARCHAR(128)[]
        CHECK (array_length(recipient_emails, 0) <= 5),
    applicant_email  VARCHAR(128)             NOT NULL,
    start_date       TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date         TIMESTAMP WITH TIME ZONE NOT NULL,
    is_full_day      BOOLEAN                  NOT NULL,
    is_online        BOOLEAN                  NOT NULL,
    author_email     VARCHAR(128)             NOT NULL,
    FOREIGN KEY (author_email) REFERENCES users (email) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE meetings;
-- +goose StatementEnd
