SELECT bill.id, bill.billdate, bill.entrydate,bill.finishdate,bill.employeeid,bill.customerid
	,billd.id,billd.billid,billd.productid,billd.quantity FROM transaction AS bill
	JOIN transaction_details AS billd ON billd.billid = bill.id;

-- customer create
INSERT INTO mst_customer (name,phonenumber,address) VALUES ($1,$2,$3) RETURNING id;

-- customer get by id
SELECT id,name,phonenumber,address FROM mst_customer WHERE id = $1;

-- customer get All
SELECT id,name,phonenumber,address FROM mst_customer;

-- customer Delete by id
DELETE FROM mst_customer WHERE id = $1;

-- customer update by id
UPDATE mst_customer SET name=$1, phonenumber=$2, address=$3 WHERE id = $4;

-- product create
INSERT INTO mst_product (name,price,unit) VALUES ($1,$2,$3) RETURNING id;

-- product get by id
SELECT id,name,price,unit FROM mst_product WHERE id = $1;

-- product get All
SELECT id,name,price,unit FROM mst_product ORDER BY id ASC;
-- product get by productName query
SELECT id,name,price,unit FROM mst_product WHERE name ILIKE '%' || $1 || '%' ORDER BY id ASC;

-- product update 
UPDATE mst_product SET name=$1, price=$2, unit=$3 WHERE id = $4;

-- product delete 
DELETE FROM mst_product WHERE id = $1 RETURNING id;

-- transaksi create menggunakan TX
INSERT INTO transaction (billdate,entrydate,finishdate,employeeid,customerid) VALUES ($1,$2,$3,$4,$5) RETURNING id;
INSERT INTO transaction_details (productid,quantity,billid) VALUES ($1,$2,$3);

-- transaksi details get by id 
SELECT id,billdate,entrydate,finishdate,employeeid,customerid FROM transaction WHERE id = $1;
SELECT trans.id,trans.productid,trans.quantity,trans.billid,product.name,product.price,product.unit FROM transaction_details AS trans JOIN mst_product AS product ON trans.productid = product.id WHERE trans.billid = $1;
SELECT id,name,phonenumber,address FROM employee WHERE id = $1;
SELECT id,name,phonenumber,address FROM mst_customer WHERE id = $1;

-- transaksi list get All
SELECT id,billdate,entrydate,finishdate,employeeid,customerid FROM transaction;
SELECT name,phonenumber,address FROM mst_customer WHERE id = $1;
SELECT name,phonenumber,address FROM employee WHERE id = $1;
SELECT trans.id,trans.productid,trans.quantity,trans.billid,product.name,product.price,product.unit FROM transaction_details AS trans JOIN mst_product AS product ON trans.productid = product.id WHERE trans.billid = $1;

-- employee get All
SELECT id,name,phonenumber,address FROM employee ORDER BY id ASC;

-- employee create
INSERT INTO employee (name,phonenumber,address) VALUES ($1,$2,$3) RETURNING id