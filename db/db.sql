	DROP TABLE IF EXISTS machines;
	DROP TABLE IF EXISTS discks;

	CREATE TABLE machines
	(
		Id INT PRIMARY KEY,
		MachineName VARCHAR(50) NOT NULL,
    	CpuCount INT,
    	TotalDiskSpace INT
	);

	CREATE TABLE discks
	(
		Id INT PRIMARY KEY,
		TotalDiskSpace INT
	);

	INSERT INTO machines (Id, MachineName, CpuCount, TotalDiskSpace) VALUES (1, 'server-1', 4, 31457280);
	INSERT INTO machines (Id, MachineName, CpuCount, TotalDiskSpace) VALUES (2, 'server-2', 4, 42567391);
	INSERT INTO machines (Id, MachineName, CpuCount, TotalDiskSpace) VALUES (3, 'server-3', 4, 53678402);

  	INSERT INTO discks (Id, TotalDiskSpace) VALUES (1, 2054617);
    INSERT INTO discks (Id, TotalDiskSpace) VALUES (2, 3145628);
	INSERT INTO discks (Id, TotalDiskSpace) VALUES (3, 4256739);