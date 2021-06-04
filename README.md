## About The Project
This API using go following a clean architecture is intended to be use a template for further projects

### Built With

* [Go]()
* [gin]()
* [go funk]()
* [postgres]()
* [docker]()


<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/DanielFCubides/nx-go-api.git
   ```

2. run database

```shell
docker run --name api-mysql -e MYSQL_ROOT_PASSWORD=r00t -e MYSQL_DATABASE=api -e MYSQL_USER=us3r -e MYSQL_PASSWORD=p455 -p 3301:3306 -d mysql:8
```

3. Install dependencies
   ```sh
   go build -o api
   ```

<!-- USAGE EXAMPLES -->

## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos
work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_


## Contact

Daniel Fernando Cubides - [@github](https://github.com/DanielFCubides/) - dfcubidesc@gmail.com
Project Link: [repo address](https://github.com/DanielFCubides/nx-go-api)
