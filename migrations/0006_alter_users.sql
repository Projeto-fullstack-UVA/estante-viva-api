ALTER TABLE users DROP CONSTRAINT IF EXISTS users_role_check;
ALTER TABLE users ADD CONSTRAINT users_role_check CHECK (role IN ('student', 'teacher', 'admin'));

ALTER TABLE users DROP COLUMN IF EXISTS campus;
ALTER TABLE users ADD COLUMN institution_id BIGINT REFERENCES institutions(id) ON DELETE SET NULL;
