-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS databases(
    id bigserial primary key,
    name varchar(30) not null,
    description varchar(50) default null,
    creation_script text not null,
    dbms_id int not null,
    CONSTRAINT dbms_id_fk FOREIGN KEY(dbms_id) REFERENCES dbms(id) ON DELETE CASCADE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS databases;
-- +goose StatementEnd
