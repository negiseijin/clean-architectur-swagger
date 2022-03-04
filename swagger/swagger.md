


# Clean Architecture Server
clean-architecture-swagger
  

## Informations

### Version

1.0.0

## Tags

  ### <span id="tag-hello"></span>hello

こんにちは

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  hello

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /get-hello | [get greeting](#get-greeting) | こんにちは取得 |
  


## Paths

### <span id="get-greeting"></span> こんにちは取得 (*getGreeting*)

```
GET /get-hello
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-greeting-200) | OK | returns a greeting |  | [schema](#get-greeting-200-schema) |

#### Responses


##### <span id="get-greeting-200"></span> 200 - returns a greeting
Status: OK

###### <span id="get-greeting-200-schema"></span> Schema
   
  


 [HelloList](#hello-list)

## Models

### <span id="hello"></span> hello


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| date | date (formatted string)| `strfmt.Date` |  | |  | `2022-02-24` |
| greet | string| `string` |  | |  | `こんにちは` |
| id | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="hello-list"></span> helloList


  

[][Hello](#hello)
