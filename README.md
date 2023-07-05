# FileDrop

FileDrop is a versatile file sharing website that allows users to upload and share files using unique IDs for easy access and collaboration.

## Table of Contents

- [Introduction](#introduction)
- [Key Features](#key-features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
- [Deployment](#deployment)
- [Contributing](#contributing)

## Introduction

FileDrop is a feature-rich file sharing platform that provides users with a seamless experience for uploading and sharing files. The project combines modern frontend technologies with secure backend storage to offer an intuitive and efficient file sharing solution.

## Key Features

- **File Sharing**: Upload and share files of any type with ease.
- **Unique ID Generation**: Every file uploaded to FileDrop is assigned a unique ID, simplifying file access and collaboration.
- **Secure Storage**: FileDrop ensures the confidentiality and privacy of uploaded files through secure storage mechanisms.
- **Efficient Searching**: Quickly locate specific files using their unique ID.
- **User-Friendly Interface**: FileDrop offers a user-friendly and intuitive interface, enhancing the overall experience for users.

## Technologies

- **Frontend**: Vue.js
- **Backend**: Go (Golang)
- **Database**: PostgreSQL

## Getting Started

To start using FileDrop locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/kalviumcommunity/FileDrop.git
   cd FileDrop
   ```

2. Install frontend dependencies:

   ```bash
   cd frontend
   npm install
   ```

3. Configure the backend:

- Create a `.env` file in the root directory (or set environment variables) with the following settings:

  ```bash
  PORT=8080
  DATABASE_URL=postgresql://user:password@localhost:5432/filedrop
  ```

4. Run the frontend:

   ```bash
   # From the root directory
   cd frontend
   npm run dev
   ```

5. Run the backend:

   ```bash
   # From the root directory
   go run main.go
   ```

6. Open your web browser and visit http://localhost:8080 to access the FileDrop application.

## Deployment

To deploy FileDrop to a production environment, follow these general steps:

1. Build the frontend for production:

   ```bash
   # From the root directory
   cd frontend
   npm run build
   ```

2. Build the backend for production:

   ```bash
   # From the root directory
   go build -o filedrop
   ```

3. Set up a PostgreSQL instance on AWS RDS or any other preferred cloud provider.

4. Set the necessary environment variables for the backend in the production environment.

5. Deploy the built backend and the frontend build to your AWS EC2 instance or any other preferred cloud server.

6. Configure your AWS security groups and network settings to allow incoming traffic on port 8080.

7. Open your web browser and visit your AWS EC2 instance's public IP or domain name to access the FileDrop application.

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please don't hesitate to submit bug reports, feature requests, and pull requests through the GitHub repository.

### Contribution Guidelines

To contribute to FileDrop, please follow these guidelines:

1. Fork the repository to your GitHub account.
2. Clone the project to your local machine.
3. Create a new branch with a descriptive name (e.g., my-new-feature).
4. Make changes to the code and commit them to your branch.
5. Push your changes to your fork (`git push origin my-new-feature`).
6. Open a pull request in the original repository, describing the changes you've made.

Please ensure that you adhere to the standard coding conventions and best practices.
