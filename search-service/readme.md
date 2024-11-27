# Dynamic Search API with Node.js, Express.js, and MySQL

This project is a RESTful API built with Node.js, Express.js, and MySQL that allows dynamic searching through a database of posts. The API supports querying by multiple fields (`title`, `content`, `author`, and `createdAt`) using query parameters.

## Features

- **Dynamic Search**: Query posts by `title`, `content`, `author`, and `createdAt`.
- **SQL Injection Prevention**: Uses parameterized queries to ensure security.
- **Flexible Matching**: Supports partial matching for text fields using the `LIKE` operator.
- **Pagination**: Easy to extend with pagination for large datasets.

## Technologies Used

### 1. **Node.js**
- A JavaScript runtime environment that allows building fast and scalable server-side applications.

### 2. **Express.js**
- A minimal and flexible Node.js web application framework for building APIs and web applications.

### 3. **MySQL**
- A widely-used relational database management system for storing structured data.

### 4. **Packages Used**
| Package    | Description |
|------------|-------------|
| `express`  | Provides tools for building robust APIs and handling HTTP requests and responses. |
| `mysql2`   | A fast and reliable MySQL client for Node.js that supports prepared statements. |
| `body-parser` | Middleware for parsing incoming request bodies in JSON or URL-encoded format (optional in this project). |

## Installation

Follow these steps to set up and run the project:

### Prerequisites
- Node.js and npm installed on your machine.
- MySQL installed and a database created.

### Steps
1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_folder>
   add .env
   npm i
   
