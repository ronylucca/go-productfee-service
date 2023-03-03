# Go Microservice Project

This project is a Go microservice that receives transactions via RabbitMQ, calculates the final product price based on a fee, and persists data.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You will need to have the following installed:

- [Go](https://golang.org/)
- [RabbitMQ](https://www.rabbitmq.com/)
- [Grafana](https://www.grafana.com/)
- [Prometheus](https://www.prometheus.com/)

### Installing

1.  Clone the repository to your local machine: `git clone https://github.com/ronylucca/go-productfee-service`

2.  Install dependencies: `go get -u ./...`

3.  Start RabbitMQ: `rabbitmq-server`

4.  Run the project: `go run main.go`

## Usage

This microservice receives transactions from RabbitMQ, calculates the final product price based on a fee, and persists data in a database (not included). The service can be used as part of an ecommerce system or any other system that requires pricing calculations and data persistence.

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
