# VipinNotes CLI

VipinNotes CLI is a powerful command-line tool for managing notes with administrative functionalities. It supports user management, authentication, and more.

---

## Table of Contents

- [Installation](#installation)
  - [Local Setup](#local-setup)
  - [Docker Setup](#docker-setup)
- [Usage](#usage)
  - [Commands](#commands)
  - [Example Usages](#example-usages)
- [Contributing](#contributing)

---

## Installation

### Local Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/vipinnotes-cli.git
   cd vipinnotes-cli
2. Build the binary:
    ```bash 
    go build -o vipinnotes vipinnotes.go
    ```
3. Move the binary to a location in your PATH:
    ``` bash
    mv vipinnotes /usr/local/bin
    ```
4. Verify installation:
    ``` bash
    vipinnotes --help
    ```

### Docker Setup
1. Clone the repository:

    ```bash
    git clone https://github.com/your-repo/vipinnotes-cli.git
    cd vipinnotes-cli
    ```

2. Build the Docker image:
    ```bash
    docker build -t vipinnotes .
    ```

3. Run the container interactively:
    ```bash
    docker run -it vipinnotes /bin/bash
    ```
4. Build the binary:
    ```bash 
    go build -o vipinnotes vipinnotes.go
    ```
5. Move the binary to a location in your PATH:
    ``` bash
    mv vipinnotes /usr/local/bin
    ```
6. Verify installation:
    ``` bash
    vipinnotes --help
    ```
# Usage

### Commands
 -  login: Log in as an admin user and save credentials locally.
 -  logout: Log out of the admin account.
 -  search: Search functionality for VipinNotes. Includes:
    -  name: <name> Search users by name, returning all users whose names - contain the string.
      - email: <email> Search users by email, returning all users whose - emails contain the string.
      - update: <email> Update the name of a user by their email.
      - users: <n> Fetch the last n users.