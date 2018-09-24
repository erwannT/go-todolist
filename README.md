# Mysql

Build docker image for mysql with initialized table
```
docker build -t mysql-todo -f dockerfile_mysql
```

Run container
```
docker run -p 3306:3306 mysql-todo
```