# SlotMachine

This repository contains an API built with Go using the Fiber framework, integrated with MongoDB and Redis for a 3x3 Slot Machine game.

## Table of Contents

- [Features](#features)
- [Dependencies](#dependencies)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
  - [Running Locally](#running-locally)
  - [Docker](#docker)
  - [Kubernetes](#kubernetes)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- Player management (create, suspend, get details)
- Slot machine gameplay with RTP calculation
- MongoDB for player data storage
- Redis for caching and gameplay statistics
- Health checks for MongoDB, Redis, and application status
- Kubernetes support for deployment

## Dependencies

- Go (v1.21+)
- Fiber (v2.26.0)
- MongoDB (v4.0+)
- Redis (v5.0+)
- Docker (optional, for containerization)
- Kubernetes (optional, for orchestration)

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go: [Installation Guide](https://golang.org/doc/install)
- MongoDB: [Installation Guide](https://docs.mongodb.com/manual/installation/)
- Redis: [Installation Guide](https://redis.io/download)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Aswin-tech-07/SlotMachine.git
   cd SlotMachine