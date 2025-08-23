
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    rawdata JSONB NOT NULL
);

-- define views that extract from json
CREATE VIEW users_v AS
SELECT
    id, created,
    rawdata ->> 'email' AS email,
    rawdata ->> 'rbacrole' AS rbacrole

FROM users;
