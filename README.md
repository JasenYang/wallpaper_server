Android:
Google launched android in 2018 and only supports Android 9.0 and above with SDK >= 28
We used the androidx instead of android support library, so the min sdk should be more than 28 (we use 29), and Android should be more than 9.0

Important:
In the code:
/Wallpaper/app/src/main/java/cs/hku/wallpaper/constants/Constants.java
You should change variable IP to your local network IP address. (You can get local ip address with command ifconfig [for linux] or ipconfig [for windows])

or

If you do not want to run a service in your local, I also deploy a server on my remote machine. Just set IP = "http://112.74.183.5:6789/"


---

Service:
We use go language with gin framework to build the service, after enter into the code root folder, you should use command 'go mod download' to get all the library those are used.

There are somethings you should change

1. handler/uploadImage.go   line 13
    change PATH to your local file path which must be exist, this PATH will be used to store the images which are uploaded.

2. db/mysql.go   line12 ~ 17
    set your mysql server username and password

3. Set a mysql database and tables
    enter into the mysql client and execute below command

```
create database wallpaper;
use wallpaper;
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
```
