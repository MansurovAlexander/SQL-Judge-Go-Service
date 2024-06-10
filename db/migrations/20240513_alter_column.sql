-- up
ALTER TABLE submission ADD COLUMN submission_id bigint NOT NULL;

--down
ALTER TABLE submission DROP COLUMN submission_id;