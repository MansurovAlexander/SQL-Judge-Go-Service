-- up
ALTER TABLE databases ADD COLUMN IF NOT EXISTS file_name VARCHAR(255) UNiQUE;
ALTER TABLE databases ALTER COLUMN name DROP NOT NULL;
ALTER TABLE databases ALTER COLUMN name SET DEFAULT 'Not restored';

ALTER TABLE assign ALTER COLUMN id DROP DEFAULT;

-- down
ALTER TABLE databases DROP COLUMN file_name;
ALTER TABLE databases ALTER COLUMN name SET NOT NULL;
ALTER TABLE databases ALTER COLUMN name DROP DEFAULT;

ALTER TABLE assign ALTER COLUMN id SET DEFAULT nextval('assign_id_seq'::regclass);