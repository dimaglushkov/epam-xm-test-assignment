## XM Test assignment

### Quick start
To start the app you need to have `docker` and `docker-compose` installed.
1. download the repo by running `git clone git@github.com:dimaglushkov/epam-xm-test-assignment.git`
2. build and start containers by running `docker-compose up --build`
3. use postman or any similar tool (or scripts from [`/bin`](/bin)) to send HTTP requests to the app

### Configuration
Configuration is done by [`.env`](/.env) file


### API
```
Method   Address           Requires auth
GET      /companies/:id    -
POST     /companies/       +
PATCH    /companies/:id    +
DELETE   /companies/:id    +
```

> To access protected endpoints (`create`, `update`, `delete`), you need to set proper
> `Authorization` HTTP header. Visit 
> [bin/create.sh](https://github.com/dimaglushkov/epam-xm-test-assignment/tree/main/bin/create.sh)
> to see an example.


### Project structure
The app is represented by several containers:
1. `db` - represents postgresql db instance
2. `zookeeper` & `kafka` - responsible for events-related communication
3. `app` - the app itself, contains the main logic of the application
4. `dummy-consumer` - simple go app responsible for consuming and logging every event sent by `app`


### Code structure
The code structure I've ended up with is highly inspired by the hexagonal 
(a.k.a. ports and adapters architecture). The main objective of this approach is separation of concerns.
The whole app is represented by several layers that are connected through 
interface. This approach allows to apply dependency inversion principle and provides,
separates core application logic from any externals, and offers extreme flexibility in adding/removing
external dependencies and tools.


### Ways to improve this solution
Due to time constraints, I have not implemented several features that I believe are necessary 
for this app to be truly production-ready:
1. Proper app monitoring (using tools such as Prometheus and Grafana)
2. Better test coverage
3. Advanced logging tools (e.g., logrus, zap, in place of the default log module)
4. Improved documentation
5. Proper Kafka & Postgres configuration

Since the task was to implement build a microservice to handle companies, I focused on the service, 
instead of infrastructure 