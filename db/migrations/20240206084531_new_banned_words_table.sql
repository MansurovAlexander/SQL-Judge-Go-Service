-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banned_words(
    id serial primary key,
    banned_word varchar(20) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banned_words;
-- +goose StatementEnd
