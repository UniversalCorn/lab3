Create tables.
	DROP TABLE IF EXISTS machines;
	DROP TABLE IF EXISTS discks;

	CREATE TABLE machines
	(
		id   SERIAL PRIMARY KEY,
		name VARCHAR(50),
    cpuCount SMALLINT,
    totalDiskSpace BIGINT,
	);

	CREATE TABLE discks
	(
		id   SERIAL PRIMARY KEY,
		totalDiskSpace BIGINT
	);

	-- Insert demo data.
	INSERT INTO machines (name) VALUES ('server-1');
  INSERT INTO machines (name) VALUES ('server-2');
  INSERT INTO machines (name) VALUES ('server-3');

	INSERT INTO discks (id) VALUES (1);
  INSERT INTO discks (id) VALUES (2);
  INSERT INTO discks (id) VALUES (3);