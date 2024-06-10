-- up
ALTER TABLE submission ADD COLUMN subtask_id numeric(3) NOT NULL;

-- down
ALTER TABLE submission DROP COLUMN subtask_id;