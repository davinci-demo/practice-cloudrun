-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="UTC";

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    rawdata JSONB NOT NULL
);

DROP TABLE IF EXISTS lessons;
CREATE TABLE lessons (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    rawdata JSONB NOT NULL
);

DROP TABLE IF EXISTS courses;
CREATE TABLE courses (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    rawdata JSONB NOT NULL
);

-- define views that extract from json
CREATE VIEW users_v AS
SELECT
    id, created,
    rawdata ->> 'email' AS email

FROM users;

CREATE VIEW courses_v AS
SELECT
    id, created,
    rawdata ->> 'courseid' AS courseid,
    rawdata ->> 'title' AS title,
    rawdata ->> 'description' AS description,
    rawdata ->> 'image' AS image,
    rawdata ->> 'subject' AS subject,
    rawdata ->> 'instructor' AS instructor,
    rawdata ->> 'updated' AS updated,
    rawdata ->> 'published' AS published

FROM courses;



