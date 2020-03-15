CREATE TABLE files (
   id BIGSERIAL PRIMARY KEY,
   -- don't remove any data from db
   removed BOOLEAN DEFAULT FALSE,
   fileName TEXT NOT NULL,
   content TEXT NOT NULL
);