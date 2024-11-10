-- +goose Up
-- +goose StatementBegin
CREATE TABLE documents
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name BIGINT NOT NULL,
    path BIGINT NOT NULL
);
 
CREATE TABLE resumes_documents
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resume_id   UUID NOT NULL,
    document_id UUID REFERENCES documents (id),
    FOREIGN KEY (resume_id) REFERENCES resumes(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resumes_documents;
DROP TABLE IF EXISTS documents;
-- +goose StatementEnd
