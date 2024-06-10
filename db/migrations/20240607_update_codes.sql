--up
DELETE FROM status;
INSERT INTO status(id, name) VALUES
(0, 'Bad word'),
(1, 'Admission word'),
(2, 'Correct answer'),
(3, 'Wrong answer'),
(4, 'Time limit exceeded'),
(5, 'Memory limit exceeded'),
(6, 'Unknown error'),
(7, 'Not checked'),
(8, 'Checked');

--down
DELETE FROM status;
INSERT INTO status(id, name) VALUES
(0, 'Bad word'),
(1, 'Accepted'),
(2, 'Time limit exceeded'),
(3, 'Memory limit exceeded'),
(4, 'Wrong answer'),
(5, 'Unknown error'),
(6, 'Not checked'),
(7, 'Checked');