-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banned_words_to_assign(
    id bigserial primary key,
    assign_id bigint not null,
    banned_word_id int not null,
    CONSTRAINT assign_id_fk FOREIGN KEY(assign_id) REFERENCES assign(id) ON DELETE CASCADE,
    CONSTRAINT banned_word_id_fk FOREIGN KEY(banned_word_id) REFERENCES banned_word(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banned_words_to_assign;
-- +goose StatementEnd
