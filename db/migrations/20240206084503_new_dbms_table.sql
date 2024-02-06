-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dbms(
    id serial primary key,
    name varchar(10) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dbms;
-- +goose StatementEnd
