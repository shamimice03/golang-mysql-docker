```
> docker compose up -d
```

```
> mysql -h localhost --protocol=TCP -u root -p

> mysql > create database recordings;

> mysql > use recordings;

```

```
> docker ps
> docker exec -i 3bfe1ee0029b mysql -u root -ptest <dbname> < /path/to/file/create-tables.sql
> docker exec -i 3bfe1ee0029b mysql -u root -ptest recordings < create-tables.sql
```



- https://github.com/joho/godotenv