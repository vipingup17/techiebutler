# TechieButler Assignment

### How to run the project

1. Clone the repo on your local machine
2. Make sure you have MySQL running on local machine on port 3306
3. Create tables in database by executing the commands present in Create table queries section
4. Make sure you have Golang installed on your local machine and execute the following command to enable go module support
    * GO111MODULE=on
5. Run go build .
6. Run go run .
7. The project should run on the default port which is 8080

### Create table queries 

    CREATE TABLE employee (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        position VARCHAR(255) NOT NULL,
        salary DECIMAL(15, 2) NOT NULL
    );

### Apis Curls and responses

##### Employee Creation Api

