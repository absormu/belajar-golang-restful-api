


# Membuat DB Mysql Docker
* Buat dan Start Container 

mkdir dbdata
docker run -itd --name=dbrestapiserver2 -p 3307:3306 -v $(pwd)/dbdata:/var/lib/mysql --env="MYSQL_ROOT_PASSWORD=root212" mysql/mysql-server
docker exec -it dbrestapiserver2 bash
mysql -u root -proot212

CREATE USER 'absor'@'%' IDENTIFIED BY 'absor';
GRANT ALL PRIVILEGES ON *.* TO 'absor'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

docker inspect a0f48df8d56b
docker container stop dc2d6bd00b5e
docker rm dc2d6bd00b5e