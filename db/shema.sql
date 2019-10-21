-- Create tables.
DROP TABLE IF EXISTS TabletsList;
DROP TABLE IF EXISTS TabletsStates;
CREATE TABLE TabletsList
(
  id   SERIAL PRIMARY KEY,
  "name" VARCHAR(50) NOT NULL
);

CREATE TABLE TabletsState
(
    id            SERIAL PRIMARY KEY,
    battery       VARCHAR(50) NOT NULL,
    devicetime    timestamp,
    serverTime    timestamp,
    currentVideo  VARCHAR(50) NOT NULL,
    tabletid      INTEGER REFERENCES TabletsList(id)
);

-- Insert demo data.
INSERT INTO TabletsList (name) VALUES ('class1-tablet2');
INSERT INTO TabletsState (battery, currentVideo, deviceTime, serverTime, tabletid) VALUES ('100%', 'video1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, (SELECT id FROM TabletsList WHERE name = 'class1-tablet2'));
