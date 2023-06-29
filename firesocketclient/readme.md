## Environments

Don't forget to change port in the Dockerfile
```javascript
HUB_GRPC_URL="0.0.0.0:XXXX"
WEBSOCKET_SERVER_URL="0.0.0.0:XXXX"
```

| Parameter              | Type     | Description                                                             |
| :--------------------- | :------- | :---------------------------------------------------------------------- |
| `HUB_GRPC_URL`         | `string` | **Required**. Grpc server URL of your HUB service                       |
| `WEBSOCKET_SERVER_URL` | `string` | **Required**. The url on which the Client's Websocket server is running |


## WebSocket
<big><pre>
[**wss://nurislom.space/ws**](wss://nurislom.space/ws)
</pre></big>

## Subscribe or Unsubscribe
Use the following JSON object to subscribe or unsubscribe from a sender:

```json
{
"Sender":"Sender unique ID",
"Subscribe":1
}

```
| Parameter   | Type     | Values | Description                                                                        |
| :---------- | :------- | :----- | :--------------------------------------------------------------------------------- |
| `Sender`    | `string` |        | **Required**. The ID of the Sender to subscribe or unsubscribe from.               |
| `Subscribe` | `int`    | 0 - 1  | **Required**. Set to `1` to subscribe to the group, or `0` to unsubscribe from it. |

## Post Message
Use the following JSON object to post a message from sender to server:
```json
{
"Sender":"123",
"UserName":"nbuzurxonov",
"Message":"Buxoro - Toshkent ertaga",
"Caption":"",
"MessageType":"text",
"MessageId":"",
"ClientName":"telegram"
}
```

| Parameter     | Type     | Description                                                                           |
| :------------ | :------- | :------------------------------------------------------------------------------------ |
| `Sender`      | `string` | **Required**. The ID of the Sender to subscribe or unsubscribe from.                  |
| `UserName`    | `string` | **Not Required**. The UserName of the Sender.                                         |
| `Message`     | `string` | **Required**. The message text(If it is not text send base64).                        |
| `Caption`     | `string` | **Not Required**. A Caption of the message.                                           |
| `MessageType` | `string` | **Required**. The type of the message.(text,photo,voice,etc.)                         |
| `MessageId`   | `string` | **Not Required**. The Unique Id of the message(If not set Id will generate by server) |
| `ClientName`  | `string` | **Required**. The name of the client.(telegram,android,ios,web,etc.) |


## Response From Server

```json
{
    "Sender": "123",
    "Message": "Buxoro - Toshkent ertaga",
    "MessageType": "text",
    "ClientName": "telegram",
    "MessageDate": 1686845200,
    "MessageCount": 1,
    "IsClient": true
}
```

```json
{
"Sender": "123",
"Message": "https://eticket.railway.uz/uz/home?sd-value=16-06-2023\u0026sd-value2=\u0026sf-name=Buxoro\u0026sf-code=2900800\u0026st-name=Toshkent\u0026st-code=2900001",
"MessageType": "text",
"Action": "utteractionapicall",
"Score": 1,
"Dep": "Buxoro",
"Arr": "Toshkent",
"Date": "16.06.2023",
"Language": "uzlt",
"ClientName": "telegram",
"Intent": "buying_ticket_uzlt",
"DepCode": "2900800",
"ArrCode": "2900001",
"MessageDate": 1686845200
}
```


| Parameter      | Type     | Values                                                                                                 | Description                                              |
| :------------- | :------- | :----------------------------------------------------------------------------------------------------- | :------------------------------------------------------- |
| `Sender`       | `string` |                                                                                                        | The ID of the Sender.                                    |
| `Message`      | `string` |                                                                                                        | The message text.                                        |
| `MessageType`  | `string` | text,photo,voice                                                                                       | The type of the message.                                 |
| `Action`       | `string` | https://gitlab.axonlogic.uz/railway/rnd/uzrailways_rasa_support_bot/-/blob/pre-prod/utils/phrases.json | The action of the message.                               |
| `Score`        | `int`    | (0,1]                                                                                                  | The score of the Action.(May be removed in the future)   |
| `Dep`          | `string` |                                                                                                        | The departure city.                                      |
| `Arr`          | `string` |                                                                                                        | The arrival city.                                        |
| `Date`         | `string` |                                                                                                        | The date of the ride.                                    |
| `Language`     | `string` | uzlt,uzcyr,ru,kr                                                                                       | The language of the message.                             |
| `ClientName`   | `string` | telegram,android,ios,web                                                                               | The name of the client.                                  |
| `Intent`       | `string` |                                                                                                        | The intent of the message.(May be removed in the future) |
| `DepCode`      | `string` |                                                                                                        | The departure city code.                                 |
| `ArrCode`      | `string` |                                                                                                        | The arrival city code.                                   |
| `MessageDate`  | `int`    |                                                                                                        | The date of the message(UNIX).                           |
| `MessageCount` | `int`    |                                                                                                        | The count of the message.(For support website)           |
| `IsClient`     | `bool`   | true,false                                                                                             | Is the message from the client.(If false may not be)     |



