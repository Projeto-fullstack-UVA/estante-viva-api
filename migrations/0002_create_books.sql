CREATE TABLE IF NOT EXISTS books (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title TEXT NOT NULL,
  author TEXT NOT NULL,
  release_date DATE NOT NULL,
  edition TEXT,
  status TEXT NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'lent')),
  created_at TIMESTAMPTZ NOT NULL
);
