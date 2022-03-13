


# Clean Architecture Server
clean-architecture-swagger
  

## Informations

### Version

1.0.0

## Tags

  ### <span id="tag-todo"></span>todo

todo

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  todo

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /create-todo | [create todo](#create-todo) | todo作成 |
| GET | /get-todo/{todoId} | [get todo](#get-todo) | todo取得 |
| GET | /get-todo | [get todo list](#get-todo-list) | todo取得 |
| PUT | /update-todo/{todoId} | [update todo](#update-todo) | todo更新 |
| PUT | /update-todo | [update todo list](#update-todo-list) | todo更新 |
  


## Paths

### <span id="create-todo"></span> todo作成 (*createTodo*)

```
POST /create-todo
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [CreateTodo](#create-todo) | `models.CreateTodo` | | ✓ | | todo作成パラメーター |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-todo-201) | Created | todoを作成しました |  | [schema](#create-todo-201-schema) |

#### Responses


##### <span id="create-todo-201"></span> 201 - todoを作成しました
Status: Created

###### <span id="create-todo-201-schema"></span> Schema

### <span id="get-todo"></span> todo取得 (*getTodo*)

```
GET /get-todo/{todoId}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| todoId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | ID of todo to return |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-todo-200) | OK | todoを取得します |  | [schema](#get-todo-200-schema) |

#### Responses


##### <span id="get-todo-200"></span> 200 - todoを取得します
Status: OK

###### <span id="get-todo-200-schema"></span> Schema
   
  

[Todo](#todo)

### <span id="get-todo-list"></span> todo取得 (*getTodoList*)

```
GET /get-todo
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-todo-list-200) | OK | todoを取得します |  | [schema](#get-todo-list-200-schema) |

#### Responses


##### <span id="get-todo-list-200"></span> 200 - todoを取得します
Status: OK

###### <span id="get-todo-list-200-schema"></span> Schema
   
  


 [TodoList](#todo-list)

### <span id="update-todo"></span> todo更新 (*updateTodo*)

```
PUT /update-todo/{todoId}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [UpdateTodo](#update-todo) | `models.UpdateTodo` | | ✓ | | todo更新パラメーター |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-todo-200) | OK | todoを更新しました |  | [schema](#update-todo-200-schema) |

#### Responses


##### <span id="update-todo-200"></span> 200 - todoを更新しました
Status: OK

###### <span id="update-todo-200-schema"></span> Schema

### <span id="update-todo-list"></span> todo更新 (*updateTodoList*)

```
PUT /update-todo
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [UpdateTodoList](#update-todo-list) | `models.UpdateTodoList` | | ✓ | | todo更新パラメーター |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-todo-list-200) | OK | todoを更新しました |  | [schema](#update-todo-list-200-schema) |

#### Responses


##### <span id="update-todo-list-200"></span> 200 - todoを更新しました
Status: OK

###### <span id="update-todo-list-200-schema"></span> Schema

## Models

### <span id="create-todo"></span> createTodo


> TODO作成
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | | TODO名 | `TODO名` |



### <span id="todo"></span> todo


> TODO
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | date (formatted string)| `strfmt.Date` |  | | 作成日 | `2022-02-24` |
| done | boolean| `bool` |  | | 完了 | `false` |
| id | uuid (formatted string)| `strfmt.UUID` |  | | id | `96c4b7d9-1849-4f6d-b616-06076af3da6e` |
| name | string| `string` |  | | TODO名 | `TODO名` |
| updatedAt | date (formatted string)| `strfmt.Date` |  | | 更新日 | `2022-02-24` |



### <span id="todo-list"></span> todoList


> TODOリスト
  



[][Todo](#todo)

### <span id="update-todo"></span> updateTodo


> TODO更新
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| done | boolean| `bool` |  | | 完了 | `false` |
| name | string| `string` |  | | TODO名 | `TODO名` |



### <span id="update-todo-list"></span> updateTodoList


> TODO更新リスト
  



[][UpdateTodo](#update-todo)
