# validgate

[![codecov](https://codecov.io/github/sshaparenko/validgate/graph/badge.svg?token=VPCRA71BD0)](https://codecov.io/github/sshaparenko/validgate)

Validgate is a service that validates the credit card number according to the Luhn Algorithm among with card expiration month and year.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)

## Requirements

- Go version 1.22.2 or higher
- Docker version 27.3.1 or higher

## Installation

Clone this repository

```bash
git clone https://github.com/sshaparenko/validgate.git
```

Build the Docker image

```bash
docker build -t validgate .
```

Run image

```bash
docker run validgate
```

You should see, that service is running on port 8080:

```
2024/10/27 17:08:19 Starting service on port: :8080
```

In order to get host, you should get the ID of container

```bash
docker ps -a | grep "validgate" | awk '{print $1}'
```

To get the host IP address run the next command

```bash
docker inspect 05a5942fd870 | grep "\"IPAddress\"" | head -n 1
```

Now, you can send requests to the `http://<IPAddress>/api/v1`

## Usage

API base URL is `/api/v1`

POST `/validate`:

```
curl -X POST http://<IPAddress>:8080/api/v1/validate \
     -H "Content-Type: application/json" \
     -H "Accept */*" \
     -d '{"card_number": "4111111111111111", "exp_month": 10, "exp_year": 2024}'
```
