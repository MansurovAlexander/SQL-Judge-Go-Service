-- up
ALTER TABLE banned_words_to_assign DROP COLUMN banned_word_id;
ALTER TABLE banned_words_to_assign ADD COLUMN banned_word varchar(50) NOT NULL;

-- down
ALTER TABLE banned_words_to_assign DROP COLUMN banned_word;
ALTER TABLE banned_words_to_assign ADD COLUMN banned_word_id integer NOT NULL;