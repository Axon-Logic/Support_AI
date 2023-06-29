## Environments

Don't forget to change port in the Dockerfile
```javascript
HUB_GRPC_URL="0.0.0.0:XXXX"
WEBSOCKET_SERVER_URL="0.0.0.0:XXXX"
GIN_SERVER_PORT="XXXX"
HISTORY_SERVICE_URL="http://0.0.0.0:XXXX"

```

| Parameter              | Type     | Description                                                             |
| :--------------------- | :------- | :---------------------------------------------------------------------- |
| `HUB_GRPC_URL`         | `string` | **Required**. Grpc server URL of your HUB service                       |
| `WEBSOCKET_SERVER_URL` | `string` | **Required**. The url on which the Client's Websocket server is running |
| `GIN_SERVER_PORT`      | `string` | **Required**. The port on which the Client's API server is running      |
| `HISTORY_SERVICE_URL`  | `string` | **Required**. The url on which the History service is running           |


## WebSocket getChats Url

<big><pre>
[**wss://nurislom.space/getChats**](wss://nurislom.space/getChats)
</pre></big>

## Response From Server

```json
{
    "client_id": "123",
    "message": "Assalomu alaykum",
    "date": 1686849700,
    "message_count": 1,
    "count": true,
    "user_name": "123",
    "first_name": "123",
    "last_name": "123",
    "owner": false,
    "type": "text",
    "message_id": "1331dasdqe",
    "caption":"Bir nimala",
    "reply_message_id":"13212-qwel"
}
```

```json
{
    "client_id": "123",
    "message": "Assalomu Alaykum, Hurmatli mijoz.",
    "date": 1686849700,
    "message_count": 0,
    "count": false,
    "user_name": "123",
    "first_name": "123",
    "last_name": "123",
    "owner": true,
    "type": "text",
    "message_id": "1331dasdqe",
    "caption":"Bir nimala",
    "reply_message_id":"13212-qwel"
}
```
| Parameter          | Type     | Description                                                                       |
| :----------------- | :------- | :-------------------------------------------------------------------------------- |
| `client_id`        | `string` | **Required**. Client's id                                                         |
| `message`          | `string` | **Required**. Message text(If it is not text it will be equal to type of message) |
| `date`             | `int`    | **Required**. Message's create date(UNIX)                                         |
| `message_count`    | `int`    | **Required**. Messages count from current client                                  |
| `count`            | `bool`   | **Required**. Is Message count equal to 0                                         |
| `user_name`        | `string` | **Required**. User's username(client_id)                                          |
| `first_name`       | `string` | **Required**. User's first name(client_id)                                        |
| `last_name`        | `string` | **Required**. User's last name(client_id)                                         |
| `owner`            | `bool`   | **Required**. Is Support owner of Message                                         |
| `type`             | `string` | **Required**. Message's type(text,photo,voice)                                    |
| `message_id`       | `string` | **Required**. Message's id                                                        |
| `caption`          | `string` | **Required**. Message's Caption                                                   |
| `reply_message_id` | `string` | **Required**. Message's reply message_id                                          |



## WebSocket getMessages

<big><pre>
[**wss://nurislom.space/getChats**](wss://nurislom.space/getMessages)
</pre></big>

## Description
Send client's id to server for subscribe client's messages, If send again but other client's id, it will unsubscribe previous client's id and subscribe new client's id

## Request to Server
```json
{
    "client_id":"123"
}
```
| Parameter   | Type     | Description               |
| :---------- | :------- | :------------------------ |
| `client_id` | `string` | **Required**. Client's id |


## Response From Server

```json
{
    "client_id": "123",
    "message": "Assalomu alaykum",
    "date": 1686849700,
    "message_count": 1,
    "count": true,
    "user_name": "123",
    "first_name": "123",
    "last_name": "123",
    "owner": false,
    "type": "text",
    "message_id": "1331dasdqe",
    "caption":"Bir nimala",
    "reply_message_id":"13212-qwel"
}
```

```json
{
    "client_id": "123",
    "message": "Assalomu Alaykum, Hurmatli mijoz.",
    "date": 1686849700,
    "message_count": 0,
    "count": false,
    "user_name": "123",
    "first_name": "123",
    "last_name": "123",
    "owner": true,
    "type": "text",
    "message_id": "1331dasdqe",
    "caption":"Bir nimala",
    "reply_message_id":"13212-qwel"
}
```
| Parameter          | Type     | Description                                                     |
| :----------------- | :------- | :-------------------------------------------------------------- |
| `client_id`        | `string` | **Required**. Client's id                                       |
| `message`          | `string` | **Required**. Message text(If it is not text it will be base64) |
| `date`             | `int`    | **Required**. Message's create date(UNIX)                       |
| `message_count`    | `int`    | **Required**. Messages count from current client                |
| `count`            | `bool`   | **Required**. Is Message count equal to 0                       |
| `user_name`        | `string` | **Required**. User's username(client_id)                        |
| `first_name`       | `string` | **Required**. User's first name(client_id)                      |
| `last_name`        | `string` | **Required**. User's last name(client_id)                       |
| `owner`            | `bool`   | **Required**. Is Support owner of Message                       |
| `type`             | `string` | **Required**. Message's type(text,photo,voice)                  |
| `message_id`       | `string` | **Required**. Message's id                                      |
| `caption`          | `string` | **Required**. Message's Caption                                 |
| `reply_message_id` | `string` | **Required**. Message's reply message_id                        |

## API update message

<big><pre>
POST [**https://nurislom.space/firesocketsupportapi/temp/update/{client_id}**](https://nurislom.space/firesocketsupportapi/temp/update/client_id)
</pre></big>

## Description
Send client's id to update message_count to 0

**HTTP Codes**:

| Code | Message | Description                                     |
| ---- | ------- | ----------------------------------------------- |
| 200  | OK      | The request was received successfully(May not). |


## API Send new message

<big><pre>
POST [**https://nurislom.space/firesocketsupportapi/temp/newMessage/{client_id}**](https://nurislom.space/firesocketsupportapi/temp/newMessage/client_id)
</pre></big>

## Description
Send new message from support to client

## Request
```json
{
"Sender":"123",
"Message":"Assalomu alaykum",
"MessageType":"text",
"ClientName":"telegram",
"Caption":"Bir nimala",
"MessageId":"lkjweq-123123jn"
}
```

| Parameter     | Type     | Description                                               |
| :------------ | :------- | :-------------------------------------------------------- |
| `Sender`      | `string` | **Required**. Client's id                                 |
| `Message`     | `string` | **Required**. Message text(If it is not text send base64) |
| `MessageType` | `string` | **Required**. Message's type(text,photo,voice)            |
| `ClientName`  | `string` | **Required**. Client's platform name                      |
| `Caption`     | `string` | **Not Required**. Message's Caption                       |
| `MessageId`   | `string` | **Not Required**. Message's Id                            |


**HTTP Codes**:

| Code | Message | Description                                     |
| ---- | ------- | ----------------------------------------------- |
| 200  | OK      | The request was received successfully(May not). |

