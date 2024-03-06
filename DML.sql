SELECT bill.id, bill.billdate, bill.entrydate,bill.finishdate,bill.employeeid,bill.customerid
	,billd.id,billd.billid,billd.productid,billd.quantity FROM transaction AS bill
	JOIN transaction_details AS billd ON billd.billid = bill.id;