CREATE DATABASE IF NOT EXISTS monitor;
USE monitor;

CREATE TABLE IF NOT EXISTS tree_verify (
    Tree_ID VARCHAR(100) PRIMARY KEY,
    Timestamp TIMESTAMP,
    Tree_Head VARCHAR(100),
    Tree_Size INT,
    Root_Hash VARCHAR(100),
    size_of_update INT
);

-- check if the database and table were created successfully
SELECT 
  CASE 
    WHEN EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = 'monitor') 
      AND EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = 'tree_verify' AND table_schema = 'monitor') 
    THEN 'Database monitor and tree_verify table created successfully.' 
    ELSE 'Error creating database and table.' 
  END AS result;