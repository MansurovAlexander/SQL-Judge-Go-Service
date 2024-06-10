-- up
CREATE TABLE status (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL
);
INSERT INTO status(id, name) VALUES (0, 'Bad words'), (1, 'Accepted'), (2, 'Time limit exceeded'), (3, 'Memory limit exceeded'), (4, 'Wrong answer'), (5, 'Unknown error'), (6, 'Not checked'), (7, 'Checked');

ALTER TABLE submission ADD COLUMN status_id integer NOT NULL REFERENCES status(id) ON DELETE CASCADE;
ALTER TABLE submission DROP COLUMN grade;

-- down

ALTER TABLE submission DROP COLUMN status_id;
ALTER TABLE submission ADD COLUMN grade integer;

DROP TABLE status;