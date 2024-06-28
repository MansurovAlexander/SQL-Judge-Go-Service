--up

ALTER TABLE assign ADD COLUMN status_id int DEFAULT 0;
ALTER TABLE assign ADD CONSTRAINT assign_status_id_fk FOREIGN KEY (status_id) REFERENCES status(id) ON DELETE RESTRICT;

--down
ALTER TABLE assign DROP CONSTRAINT assign_status_id_fk;
ALTER TABLE assign DROP COLUMN status_id;