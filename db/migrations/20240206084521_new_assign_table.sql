-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS assign(
    id bigserial primary key,
    time_limit int default null,
    memory_limit int default null,
    correct_script text not null,
    db_id bigint not null,
    CONSTRAINT db_id_fk FOREIGN KEY(db_id) REFERENCES databases(id) ON DELETE CASCADE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS assign;
-- +goose StatementEnd
