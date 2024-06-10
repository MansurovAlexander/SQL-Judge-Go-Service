--up
ALTER TABLE assign ADD COLUMN subtask_id numeric(3) NOT NULL DEFAULT -1;
ALTER TABLE assign ADD COLUMN assign_id bigint NOT NULL DEFAULT -1;

--down
ALTER TABLE assign DROP COLUMN subtask_id;
ALTER TABLE assign DROP COLUMN assign_id;