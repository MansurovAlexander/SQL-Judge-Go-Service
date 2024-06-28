--up
INSERT INTO status(id, name) VALUES
(9, 'Empty output');

--down
DELETE FROM status where id = 9;