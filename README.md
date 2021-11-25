


# Membuat DB Mysql Docker
* Buat dan Start Container 
```
    mkdir dbdata
    docker run -itd --name=dbrestapiserver2 -p 3307:3306 -v $(pwd)/dbdata:/var/lib/mysql --env="MYSQL_ROOT_PASSWORD=root212" mysql/mysql-server
    docker exec -it dbrestapiserver2 bash
    mysql -u root -proot212

    CREATE USER 'absor'@'%' IDENTIFIED BY 'absor';
    GRANT ALL PRIVILEGES ON *.* TO 'absor'@'%' WITH GRANT OPTION;
    FLUSH PRIVILEGES;

    docker container start a0f48df8d56b
    docker inspect a0f48df8d56b
    docker container stop a0f48df8d56b
    docker rm a0f48df8d56b

    langkah start:
    docker container start a0f48df8d56b
    docker exec -it dbrestapiserver2 bash
    mysql -u root -proot212
    use dbrestapi;
```

* DB 
```
    create database dbrestapi;
    use dbrestapi;

    * User App
    CREATE USER 'restapi'@'%' IDENTIFIED BY 'P6Bfy4PNmk';
    GRANT ALL PRIVILEGES ON *.* TO 'restapi'@'%';  
    FLUSH privileges;

    CREATE table category 
    (
	id integer primary key auto_increment,
	name varchar(200) not null
    ) engine = InnoDB;

    SELECT * FROM category ;


```
* db test
```
    create database dbrestapi_test;
    use dbrestapi_test;

    CREATE table category 
    (
	id integer primary key auto_increment,
	name varchar(200) not null
    ) engine = InnoDB;

    SELECT * FROM category ;


```