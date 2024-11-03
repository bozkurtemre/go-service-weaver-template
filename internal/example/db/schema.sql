CREATE TABLE examples (
  id         BIGSERIAL   PRIMARY KEY,
  message    text        NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);