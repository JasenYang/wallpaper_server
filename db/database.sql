create table user (
uid int primary key auto_increment,
name varchar(200),
password varchar(200),
pid varchar(200)
);

create table image (
pid int primary key auto_increment,
name varchar(200),
classify varchar(20),
filename varchar(200),
uid int
);