--up
ALTER TABLE submission DROP CONSTRAINT submission_assign_id_fk;
ALTER TABLE submission ADD CONSTRAINT submission_assign_id_fk FOREIGN KEY (assign_id) REFERENCE assign(assign_id) ON DELETE CASCADE;