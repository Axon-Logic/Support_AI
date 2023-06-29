# Environments (Don't forget to open ports for GRPC and GIN in the Dockerfile)
```
GRPC_SERVER_PORT=XXXX
GIN_SERVER_PORT=XXXX
```

# GET PROVIDER

```http
GET /getProvider
```
## Description
This API is used to get the hostname and port of the Master provider.

## Request
```
```

## Responses
```json
{
    "HostName": "Host Name",
    "GrpcPort": "Port",
    "GinPort": "Port"
}
```

GetProvider returns the following status codes:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK: The request was successful(However, the provider may not be set)` |
| 500 | `INTERNAL SERVER ERROR` |



# GRPC


## Messages
```proto
message Provider {
    string HostName = 1;
	string GrpcPort = 2;
    string GinPort = 3;
}

message Empty {}
```
## Services

```proto
service Main {
    rpc RegisterProvider (stream Provider) returns (stream Empty) {}
    rpc GetMasterProvider (Empty) returns (Provider) {}
}
```

## Description
RegisterProvider is used to register a provider to the master provider. The first connection will be the master, If the connection is lost, the provider will be removed from the list of providers and the master goes to the second connection.
GetMasterProvider is used to get the master provider. 