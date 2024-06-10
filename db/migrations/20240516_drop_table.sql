-- up
DROP TABLE banned_word;

-- down
CREATE TABLE banned_word (
    id SERIAL PRIMARY KEY,
    banned_word varchar(50) NOT NULL
)