# Hexagonal Architecture 

When developing software, engineers usually face two main types of complexity:

1. Business Complexity: comes from the product and usually can't be avoided. Certain products are more complex than other and this can't be changed. 
1. Technical Complexity: added by engineers in order to solve the problem they are trying to solve. It comes from the framework, architecture or even database and sometimes can be avoided or at least mitigated. 

Being able to separate these two will help you write more maintainable and scalable software, since you can change the technical complexity without touching the business complexity. In other words it allows you to change the framework, architecture or database without changing the business logic.
