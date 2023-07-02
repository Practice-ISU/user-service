# user-service

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![GRPC](https://img.shields.io/badge/grpc-2da7b0?style=for-the-badge&logoColor=white)
![License MIT](https://img.shields.io/badge/MIT-aa0000?style=for-the-badge&logoColor=white)

In this repository you will find the code for the user management service that is used in Gatekeeper. This service provides methods for registration, login and verification of users.

The following environment variables are required for use:

* SERVICE_NAME: The name of the service that will be used to identify your service.
* DISCOVERY_ADDR: The address of the discovery service where your service will register or where it will be able to find other services.
* PORT_PING: The port on which your service will listen for requests to check its availability (ping).
* PORT_USERS: The port on which your service will listen for requests related to user functionality.
* SERVICE_IP: The IP address of your service, which can be used by other services to communicate directly with your service.

Make sure that these environment variables are set and configured correctly in your configuration.
