BEGIN;
ALTER TABLE projects ALTER COLUMN version SET DEFAULT 1;
COMMIT;