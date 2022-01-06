create table products (
	id int auto_increment primary key,
    `description` varchar(45),
    price float
);

create table customers (
	id int auto_increment primary key,
    last_name varchar(45),
    first_name varchar(45),
    `condition` varchar(45)
);
create table invoices (
	id int auto_increment primary key,
    `datetime` datetime,
    idcustomer int,
    total float,
    foreign key (idcustomer) references customers(id)
);
create table sales (
	id int auto_increment primary key,
    idinvoice int,
    idproduct int,
    quantity float,
    foreign key (idinvoice) references invoices(id),
    foreign key (idproduct) references products(id)
);