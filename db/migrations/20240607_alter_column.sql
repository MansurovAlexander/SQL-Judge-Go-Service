--up
ALTER TABLE submission ALTER COLUMN submission_id DROP NOT NULL;

--down
ALTER TABLE submission ALTER COLUMN submission_id SET NOT NULL;