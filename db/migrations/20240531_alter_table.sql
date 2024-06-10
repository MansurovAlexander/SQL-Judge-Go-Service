--up
ALTER TABLE banned_words_to_assign DROP COLUMN banned_word;
ALTER TABLE banned_words_to_assign ADD COLUMN banned_words text DEFAULT '';
ALTER TABLE banned_words_to_assign ADD COLUMN admission_words text DEFAULT '';
ALTER TABLE banned_words_to_assign ADD COLUMN subtask_id integer NOT NULL;
--down
ALTER TABLE banned_words_to_assign RENAME COLUMN banned_words TO banned_word;
ALTER TABLE banned_words_to_assign ALTER COLUMN banned_word TYPE varchar(255);
ALTER TABLE banned_words_to_assign DROP COLUMN admission_words;
ALTER TABLE banned_words_to_assign DROP COLUMN subtask_id