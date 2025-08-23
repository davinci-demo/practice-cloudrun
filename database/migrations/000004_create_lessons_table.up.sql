
CREATE TABLE lessons (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    rawdata JSONB NOT NULL
);

-- define views that extract from json
CREATE VIEW lessons_v AS
SELECT
    id, created,
    rawdata ->> 'lessonid' AS lessonid,
    rawdata ->> 'title' AS title,
    rawdata ->> 'content' AS content,
    rawdata ->> 'resourceurl' AS resourceurl

FROM lessons;
