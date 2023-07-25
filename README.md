

# FortiSafe

## ℹ️ Problem Statement
Build a robust containerized task management system to handle user authentication, authorization and access management.

## 📍 Key Features
- Secure user registration and authentication
- Account Deactivation and Deletion: Allow users to deactivate or delete their accounts, if applicable. Implement a mechanism to handle account deletion securely while considering data retention policies.
- Role-based and Group-based access management on resources(Tasks) with ability to create custom roles and groups (Need to make sure endpoints are secure)
- Protection against vulnerabilities like SQL injection attacks
- Support for bulk upload using CSV(Both users and tasks) making sure all the relationships are preserved accurately

## ⚙️ Tech Stack 
- **GoLang** - Used for developing efficient and fast server-side applications due to its compiled nature and strong concurrency support.
- **AWS RDS (PostgreSQL Instance)** - Utilized as a managed database service to provide scalable, reliable, and performant storage for the application.
- **Docker** - Employed for containerizing the application, ensuring consistency across different environments and facilitating easy deployment and scaling.
- **Nginx** - Used as a reverse proxy to efficiently handle client requests, load balance, and improve security by serving as a barrier between clients and the application server.

## ⚒️ Go Packages Used
- **uuid** -  Generates unique identifiers for entities.
- **jwt** - Creates secure JSON Web Tokens for authentication.
- **bcrypt** - Hashes and encrypts passwords securely.
- **gorm** - Simplifies database interactions with an ORM in Go. It also protects from **SQL Injection**.
- **gofiber** - Fast and efficient web framework for building APIs in Go.
- **godotenv** - Loads environment variables from a .env file.
- **postgres** - Robust and scalable relational database management system.

## 🔧 Getting Started
To get a local copy up and running follow these simple steps.

### 👉🏻 Prerequisites
In order to get a copy of the project and run it locally, you'll need to have Go (v1.15 or later) and Docker installed on your machine.

If you don't have Go installed, you can download it from the [official Go website](https://go.dev/doc/install). After installation, you can verify it by typing `go version` in your terminal. It should display the installed version of Go.

For Docker, you can download it from the [official Docker website](https://www.docker.com/products/docker-desktop/). After installation, you can verify it by typing `docker --version` in your terminal. It should display the installed version of Docker.

Make sure you also have a working Docker Compose. Docker Desktop installs Docker Compose by default on Mac and Windows, but you might need to add it separately in some Linux distributions. You can check its availability by typing `docker-compose --version` in your terminal.

### 👉🏻 Get Local Copy
1. Clone the Repository
```bash
git clone https://github.com/prasoonsoni/FortiSafe
```
2. Change the directory
```bash
cd FortiSafe
```
### 👉🏻 Create Environment Variables
1. Change the name of `.env.example` to `.env`
2. Add the following variables to `.env` file
```env
DB_HOST = <your-db-host>
DB_NAME = <your-db-name>
DB_USER = <your-db-user>
DB_PASSWORD = <your-db-password>
DB_PORT= <your-db-port>
JWT_SECRET = <your-jwt-secret>
ADMIN_EMAIL = <your-admin-email>
ADMIN_PASSWORD = <your-admin-password>
```


### 👉🏻 Running the Project
#### 1. Using Docker
In order to test our service we first need to build and run docker-compose. Docker-compose will automate the build and the run of our two Dockerfile.
To run this commands you must be in the repository’s root.
1. Build the Image
```bash
docker-compose build
```
2. Start the service
```bash
docker-compose up -d
```
Now we have and built the image and service is started for both **go** and **nginx** (used for reverse-proxy).
The Nginx reverse proxy will send all request from `localhost/fortisafe/` to Golang service on port `3000`.

Backend is accessible at `http://localhost/fortisafe/`

#### 2. Without Docker
1. Download the required packages
```bash
go mod download
```
2. Run the `main.go`
```bash
go run main.go
```
> Note - When running without Docker we don't have access to reverse proxy (nginx) service.

Backend is accessible at `http://localhost:3000/`


## 📂 Complete Project Folder Structure
```
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
│
├── controllers
│   ├── groupController.go
│   ├── permissionController.go
│   ├── resourceController.go
│   ├── roleController.go
│   └── userController.go
│
├── db
│    ├── db.go
│    └── migrate.go
│
├── middlewares
│   ├── authenticateAdmin.go
│   └── authenticateUser.go
│
├── models
│   ├── account_status_logs.go
│   ├── body.go
│   ├── group.go
│   ├── permission.go
│   ├── resource.go
│   ├── response.go
│   ├── role.go
│   ├── role_permission.go
│   └── user.go
│
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
│
└── routes
    ├── groupRoutes.go
    ├── permissionRoutes.go
    ├── resourceRoutes.go
    ├── roleRoutes.go
    └── userRoutes.go
```
## 🔐 Pre Configured Permissions
> Note - These are the basic permissions considered while creating this project.
1. **create**: This permission allows a user to create new resources or data in the system. 
2. **read**: This permission gives a user the ability to read and retrieve existing resources or data.
3. **update**: This permission grants a user the ability to modify or update existing resources or data.
4. **delete**: This permission enables a user to remove existing resources or data from the system.

## 🔦 Basic Workflow
![architecture](https://github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/assets/75159757/e92121cc-99ac-4f83-b5a1-dbabc59b0d32)


## 📖 API References
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/28558819-fbc27156-acd1-40fb-911f-053538bf7dda?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D28558819-fbc27156-acd1-40fb-911f-053538bf7dda%26entityType%3Dcollection%26workspaceId%3D7daa153e-aea8-4ce7-a519-f33bbddc43eb)
[![Postman API Docs](https://img.shields.io/badge/Postman%20API%20Docs-FF6C37?style=for-the-badge&logo=Postman&logoColor=white)](https://documenter.getpostman.com/view/28558819/2s946mZ9Ld)

### User

#### 1. Create User

```http
  POST /api/user/create
```

Body
| Parameter | Type     |
| :-------- | :------- |
| `name` | `string` |
| `email` | `string` |
| `password` | `string` |
| `role_id` | `string` |
| `group_id` | `string` |

#### 2. Login User

```http
  POST /api/user/login
```

Body
| Parameter | Type     |
| :-------- | :------- |
| `email` | `string` |
| `password` | `string` |

#### 3. Get User

```http
  GET /api/user/get
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <user-auth-token>` |

#### 4. Deactivate User

```http
  PUT /api/user/deactivate
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <user-auth-token>` |

#### 5. Activate User

```http
  PUT /api/user/activate
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <your-auth-token>` |

#### 6. Delete User

```http
  DELETE /api/user/delete
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <your-auth-token>` |

#### 7. Bulk Create User

```http
  POST /api/user/create/bulk
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Form Data
| Key | Value     |
| :-------- | :------- |
| `users` | `.csv file` |

#### 8. Login Admin

```http
  POST /api/admin/login
```

Body
| Parameter | Type     |
| :-------- | :------- |
| `email` | `string` |
| `password` | `string` |

### Permission

#### 1. Create Permission

```http
  POST /api/permission/create
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `name` | `string` |
| `description` | `string` |

#### 2. Get All Permissions

```http
  GET /api/permission/all
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

### Role

#### 1. Create Role

```http
  POST /api/role/create
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `name` | `string` |
| `description` | `string` |
| `permissions` | `[<permission-id>, <permission-id>...]` |

#### 2. Add Permission

```http
  PUT /api/role/permission/add
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `role_id` | `string` |
| `permissions` | `[<permission-id>, <permission-id>...]` |

#### 3. Get All Roles

```http
  GET /api/role/get/all
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

#### 4. Remove Permission

```http
  DELETE /api/role/permission/remove
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `role_id` | `string` |
| `permission_id` | `string` |

#### 5. Assign Role

```http
  PUT /api/role/assign
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `user_id` | `string` |
| `role_id` | `string` |

#### 6. Unassign Role

```http
  PUT /api/role/unassign?user_id=<user-id>
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Query Params
| Parameter | Type     |
| :-------- | :------- |
| `user_id` | `string` |

### Resource

#### 1. Create Resource

```http
  POST /api/resource/create
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <user-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `name` | `string` |
| `description` | `string` |

#### 2. Get Resource

```http
  GET /api/resource/get/:resource_id
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <user-auth-token>` |

Path Variables
| Parameter | Type     |
| :-------- | :------- |
| `resource_id` | `string` |

#### 3. Update Resource

```http
  PUT /api/resource/update/:resource_id
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <user-auth-token>` |

Path Variables
| Parameter | Type     |
| :-------- | :------- |
| `resource_id` | `string` |

#### 4. Delete Resource

```http
  DELETE /api/resource/delete/:resource_id
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <user-auth-token>` |

Path Variables
| Parameter | Type     |
| :-------- | :------- |
| `resource_id` | `string` |

#### 5. Add Associated Role

```http
  PUT /api/resource/role/add
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `resource_id` | `string` |
| `roles` | `[<role-id>, <role-id>...]` |

#### 6. Remove Associated Role

```http
  DELETE /api/resource/role/remove
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `resource_id` | `string` |
| `role_id` | `string` |

#### 7. Bulk Create Resource

```http
  POST /api/user/create/bulk
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <user-auth-token>` |

Form Data
| Key | Value     |
| :-------- | :------- |
| `resources` | `.csv file` |

#### 8. Add Associated Group

```http
  PUT /api/resource/group/add
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `resource_id` | `string` |
| `groups` | `[<group-id>, <group-id>...]` |

#### 9. Remove Associated Group

```http
  DELETE /api/resource/group/remove
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `resource_id` | `string` |
| `group_id` | `string` |

### Group

#### 1. Create Group

```http
  POST /api/group/create
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `name` | `string` |
| `description` | `string` |
| `permissions` | `[<permission-id>, <permission-id>...]` |

#### 2. Add Permission

```http
  PUT /api/group/permission/add
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `group_id` | `string` |
| `permissions` | `[<permission-id>, <permission-id>...]` |

#### 3. Remove Permission

```http
  DELETE /api/group/permission/remove
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `group_id` | `string` |
| `permission_id` | `string` |

#### 4. Assign Group

```http
  PUT /api/group/assign
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `user_id` | `string` |
| `group_id` | `string` |

#### 5. Unassign Group

```http
  PUT /api/group/unassign?user_id=<user-id>
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <admin-auth-token>` |

Query Params
| Parameter | Type     |
| :-------- | :------- |
| `user_id` | `string` |

## 📷 Screenshots
1. Building Docker Image
![Building Image](https://github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/assets/75159757/7ab76087-010f-4c33-aeeb-f2a6c4e36c52)

2. Running Docker Image
![Running Docker Image](https://github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/assets/75159757/f2aaaf80-717f-49f0-b05d-0cad254c8a11)

3. Accessing host using reverse proxy
![Accessing host using reverse proxy](https://github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/assets/75159757/3f62c39b-f500-4eb6-8651-371651c0a0f3)






