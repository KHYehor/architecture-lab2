-- Create tables.
DROP TABLE IF EXISTS States;
DROP TABLE IF EXISTS TabletsList;
CREATE TABLE TabletsList
(
  id   SERIAL PRIMARY KEY,
  "name" VARCHAR(50) NOT NULL
);

CREATE TABLE State
(
    id            SERIAL PRIMARY KEY,
    battery       VARCHAR(50) NOT NULL,
    devicetime    TIME with time zone NOT NULL,
    timestamp     TIME with time zone  NOT NULL,
    currentVideo  VARCHAR(50) NOT NULL,
    tablesid      INTEGER,
    FOREIGN KEY (tabletid) REFERENCES Tablets(id)
);

-- Insert demo data.

