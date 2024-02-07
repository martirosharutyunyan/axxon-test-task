-- +goose Up
-- +goose StatementBegin
CREATE TYPE task_status AS ENUM ('DONE', 'IN_PROGRESS', 'ERROR', 'NEW');
CREATE TABLE tasks
(
    id         uuid        DEFAULT uuid_generate_v4() NOT NULL,
    created_at timestamp   DEFAULT CURRENT_TIMESTAMP  NOT NULL,
    updated_at timestamp   DEFAULT CURRENT_TIMESTAMP  NOT NULL,
    url        varchar                                NOT NULL,
    headers    text[],
    length     integer,
    status     task_status DEFAULT 'NEW'              NOT NULL,
    CONSTRAINT pk_tasks PRIMARY KEY (id),
    CONSTRAINT unq_url UNIQUE (url)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
