[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/YCCXVJKc)
[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-718a45dd9cf7e7f842a935f5ebbe5719a5e09af4491e668f4dbf3b35d5cca122.svg)](https://classroom.github.com/online_ide?assignment_repo_id=11469441&assignment_repo_type=AssignmentRepo)

# Role & Group Based Access Control System

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
- **gorm** - Simplifies database interactions with an ORM in Go.
- **gofiber** - Fast and efficient web framework for building APIs in Go.
- **godotenv** - Loads environment variables from a .env file.
- **postgres** - Robust and scalable relational database management system.

## 📖 API References

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
| `Authorization` | `Bearer <your-auth-token>` |

#### 4. Deactivate User

```http
  PUT /api/user/deactivate
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

Form Data
| Key | Value     |
| :-------- | :------- |
| `users` | `.csv file` |

### Permission

#### 1. Create Permission

```http
  POST /api/permission/create
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

### Role

#### 1. Create Role

```http
  POST /api/role/create
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <your-auth-token>` |

Body
| Parameter | Type     |
| :-------- | :------- |
| `name` | `string` |
| `description` | `string` |
| `permissions` | `[<permission-id>, <permission-id>...]` |

#### 2. Add Permission

```http
  PUT /api/role/add
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

#### 4. Remove Permission

```http
  DELETE /api/role/permission/remove
```

Header
| Key | Value     |
| :-------- | :------- |
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

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
| `Authorization` | `Bearer <your-auth-token>` |

Form Data
| Key | Value     |
| :-------- | :------- |
| `resources` | `.csv file` |


