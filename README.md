Client - react
GraphQL_server - api gateway
Auth_service - сервис аунтентификации/авторизации (grpc)
Health_service - сервис данных здоровья клиента (grpc)
Note_service - сервис для записей клиента (grpc)
Crew_service - сервис для сотрудников(врачей), хранят записи (из Note_service - RabbitMQ для добавления записи к врачу, из Auth_service - grpc для получения списка врачей)