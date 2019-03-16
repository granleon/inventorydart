DROP DATABASE IF EXISTS qc CASCADE;
CREATE USER IF NOT EXISTS leader;
CREATE DATABASE IF NOT EXISTS qc;
GRANT ALL ON DATABASE qc TO leader;

CREATE TABLE IF NOT EXISTS qc.partnumbers (
  id SERIAL PRIMARY KEY,
  partnumber STRING,
  chemistry STRING,
  createdAt TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS qc.inventory (
  id SERIAL PRIMARY KEY,
  partnumberid SERIAL NOT NULL REFERENCES qc.partnumbers (id),
  lotnumber STRING,
  createdAt TIMESTAMPTZ
);

INSERT INTO qc.partnumbers (partnumber, chemistry, createdAt) VALUES 
  ('444444', 'ApoA', '2016-03-26 10:10:10-05:00'),
  ('444445', 'ApoB', '2016-03-28');

INSERT INTO qc.inventory (lotnumber, createdAt, partnumberid) VALUES
  ( 'M808080', '2016-03-21 10:10:10-05:00', (SELECT id FROM qc.partnumbers AS pn WHERE pn.partnumber = '444444' )),
  ( 'M808081', '2016-03-22', (SELECT id FROM qc.partnumbers AS pn WHERE pn.partnumber = '444444' )),
  ( 'M808081', '2016-03-22 10:10:10-05:00', (SELECT id FROM qc.partnumbers AS pn WHERE pn.partnumber = '444444' )),
  ( 'M808082', '2016-03-28', (SELECT id FROM qc.partnumbers AS pn WHERE pn.partnumber = '444445' ));