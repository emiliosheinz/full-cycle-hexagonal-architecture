# Hexagonal Architecture 

When developing software, engineers usually face two main types of complexity:

1. Business Complexity: comes from the product and usually can't be avoided. Certain products are more complex than other and this can't be changed. 
1. Technical Complexity: added by engineers in order to solve the problem they are trying to solve. It comes from the framework, architecture or even database that was chosen and sometimes can be avoided or at least mitigated. 

Being able to separate these two will help you write more maintainable and scalable software, since you can change the technical complexity without touching the business complexity. In other words it allows you to change the framework, architecture or database without changing the business logic.

Hexagonal Architecture is a way to achieve this separation. It was created by Alistair Cockburn in 2005 and it's also known as Ports and Adapters. The main idea is to separate the core business logic from the external world, such as databases, frameworks, etc through adapters.

<img src="./docs/images/hexagonal-architecture.gif" alt="Hexagonal Architecture" width="100%"/>

The main concept that allows this separation to happen is the Dependency Inversion Principle. High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.

## Running Locally

1. Clone this repo
2. Run `docker-compose up -d`
3. Run `docker exec -it appproduct sh`
4. Run `sqlite3 sqlite.db` to access the database
5. Run `create table products(id string, name string, price float, status string);` to create the products table 

### Now you can use the cli to interact with the app

1. Create a new product
    ```bash
    go run main.go cli -a=create -p=<product_price> -n="<product_name>"
    ```

1. Enable a product
    ```bash
    go run main.go cli -a=enable -i=<product_id>
    ```

1. Disable a product
    ```bash
    go run main.go cli -a=disable -i=<product_id>
    ```

1. Get product details 
    ```bash
    go run main.go cli -i=<product_id>
    ```

## Or run the HTTP server and use the API

1. Run the server
    ```bash
    go run main.go http
    ```

1. Create a new product
    ```bash
    curl -X POST http://localhost:9000/product -d '{"name": "<product_name>", "price": <product_price>}'
    ```
1. Get product details
    ```bash
    curl -X GET http://localhost:9000/product/<product_id>
    ```

1. Enable a product
    ```bash
    curl -X POST http://localhost:9000/product/<product_id>/enable
    ```

1. Disable a product
    ```bash
    curl -X POST http://localhost:9000/product/<product_id>/disable
    ```

