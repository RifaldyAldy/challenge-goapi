CREATE TABLE mst_customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phoneNumber VARCHAR(13) NOT NULL,
    address VARCHAR(255)
);

CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phonenumber VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE mst_product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price int NOT NULL,
    unit VARCHAR(10) NOT NULL
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    billDate DATE,
    entryDate DATE,
    finishDate DATE,
    employeeId int NOT NULL,
    customerId int NOT NULL,
    FOREIGN KEY (employeeId) REFERENCES employee(id),
    FOREIGN KEY (customerId) REFERENCES mst_customer(id)
);

CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    billId int NOT NULL,
    productId int NOT NULL,
    quantity int NOT NULL,
    FOREIGN KEY (billId) REFERENCES transaction(id),
    FOREIGN KEY (productId) REFERENCES mst_product(id)
);

