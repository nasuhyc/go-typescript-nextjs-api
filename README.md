# Project Title

In this project, Go (Fiber, Gorm) and PostgreSQL are used on the server side, React, Typescript and NextJS are used on the client side.

## Run it on your computer

Clone the project

```bash
git clone https://github.com/nasuhyc/go-typescript-nextjs-api.git
```

Go to the frontend directory.

```bash
  cd frontend
```

Install required packages

```bash
  npm install
```

Run the server

```bash
  npm run dev
```

Important

```
** Please edit your env information(.env)
** Please create your database in potsgres.

```

Go to the backend directory.

```bash
  cd backend
```

Run the server

```bash
  air || go run main.go
```

## API Usage

#### Create

```http
  POST  /api/user/create
```

| Parametre    | Tip      | Açıklama              |
| :----------- | :------- | :-------------------- |
| `first_name` | `string` | Nasuh                 |
| `last_name`  | `string` | Yücel                 |
| `age`        | `int`    | 26                    |
| `email`      | `string` | nasuhyc@gmail.com     |
| `file`       | `string` | .png , jpg, jpeg, gif |

#### Get All

```http
  GET  /api/user/getAll
```

#### Get By ID

```http
  GET  /api/user/get/${id}
```

#### Delete

```http
  GET  /api/user/delete/${id}
```

#### Update

```http
  PUT  /api/user/update/${id}
```

| Parametre    | Tip      | Açıklama              |
| :----------- | :------- | :-------------------- |
| `first_name` | `string` | Nasuh                 |
| `last_name`  | `string` | Yücel                 |
| `age`        | `int`    | 26                    |
| `email`      | `string` | nasuhyc@gmail.com     |
| `file`       | `string` | .png , jpg, jpeg, gif |
