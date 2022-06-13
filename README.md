# Laptop store
Traning project. The main purpuse of the project is to implement microservice architecture.

### db_dump
Here you can find the dump of database. If you want to test the program you can resotre it.

### cmd
/server - to start grpc server within laptop service
/clinet - a simple client connecting to laptop service server and call create laptop method

### config
Here you can find all configurations, sush as port, host, user, and so on.

### pb & proto
- /proto - all proto files
- /pb - generated code from protos

### migrations
Database migrations. You can make changes and use command in Makefile to migrate up.

### sample
Here you can find methods to generate random data for laptop object.

### service 
Here you can finc business logic of project.

### storage 
Entities and Repos are implemented to manipulate with database.
