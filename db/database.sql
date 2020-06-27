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

create table model (
pid int primary key auto_increment,
name varchar(200),
classify varchar(20),
model_path varchar(200),
images_path varchar(200),
uid int
);
