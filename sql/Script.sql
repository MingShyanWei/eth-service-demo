-- 0.00134025
Explain
SELECT * FROM `blocks` ORDER BY Num DESC LIMIT 10

Explain
SELECT `tx_hash` FROM `transactions` WHERE num = 20728139

Explain
SELECT * FROM `blocks` ORDER BY Num DESC LIMIT 10


Explain
SELECT * FROM `transactions` WHERE tx_hash = '0x00022e52e02a3de454789b20cbe69333f4646d52f8e2a7bdd9ab100dc44623b9'


Explain
SELECT * FROM `transaction_logs` WHERE tx_hash = '0x00022e52e02a3de454789b20cbe69333f4646d52f8e2a7bdd9ab100dc44623b9'


select @@profiling 

set profiling=1;  

show profiles