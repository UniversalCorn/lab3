--Create tables.
	DROP TABLE IF EXISTS machines;
	DROP TABLE IF EXISTS discks;

	CREATE TABLE machines
	(
		Id INT AUTO_INCREMENT PRIMARY KEY,
		MachineName VARCHAR(50) NOT NULL,
    CpuCount INT,
    TotalDiskSpace INT,
	);

	CREATE TABLE discks
	(
		Id INT AUTO_INCREMENT PRIMARY KEY,
		TotalDiskSpace INT
	);

	-- Insert demo data.
	INSERT INTO machines (MachineName, CpuCount, TotalDiskSpace)
  VALUES
  ('server-1', 4, 31457280),
  ('server-2', 4, 42567391),
  ('server-3', 4, 53678402);

  INSERT INTO discks (TotalDiskSpace)
  VALUES
  (2054617),
  (3145628),
  (4256739);