-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS submission(
    id bigserial primary key,
    student_id bigint not null,
    time int default null,
    memory int default null,
    script text not null,
    grade varchar(40) not null,
    assign_id bigint not null,
    CONSTRAINT submission_assign_id_fk FOREIGN KEY(assign_id) REFERENCES assign(id) ON DELETE CASCADE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS submission;
-- +goose StatementEnd
