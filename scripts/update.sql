UPDATE queue 
SET STATUS = $1, WORKER = $2
WHERE ID = $3
AND STATUS = $4;