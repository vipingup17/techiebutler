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

##### Employee Create Api

    Request

    curl --location --request POST 'http://127.0.0.1:8080/employee/create' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "employee 3",
        "position": "position 3",
        "salary": 4500000.00
    }'

    Response

    {
        "status": "SUCCESS"
    }

##### Employee Get Api

    Request

    curl --location --request GET 'http://127.0.0.1:8080/employee/get/1' \
    --data-raw ''

    Response

    {
        "name": "employee 1",
        "position": "position 1",
        "salary": 3500000
    }

##### Employee Update Api    

    Request

    curl --location --request PUT 'http://127.0.0.1:8080/employee/update/2' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "employee 2010",
        "salary": 3500000
    }'

    Response

    {
        "name": "employee 2010",
        "salary": 3500000
    }

##### Employee Delete Api     

    Request

    curl --location --request DELETE 'http://127.0.0.1:8080/employee/delete/1'

    Response

    {
        "status": "SUCCESS"
    }

##### List Employees with pagination support

    Request

    curl --location --request GET 'http://127.0.0.1:8080/employee/list-all-employees?page=1&limit=10'

    Response

    {
        "employees_list": [
            {
                "name": "employee 2",
                "position": "position 2",
                "salary": 4500000
            },
            {
                "name": "employee 3",
                "position": "position 3",
                "salary": 4500000
            }
        ]
    }