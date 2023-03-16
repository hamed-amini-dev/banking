# Accounts Service

## Build

#### build for windows

```
make build-windows
```

#### build for linux

```
make build-linux
```

#### build for mac

```
make build-mac
```

## Run tests

make test

## API requests

### Get Account Information

#### Route Url Schema= "http://localhost:8080/accounts/{account-id}"

```
curl -X "GET"  "http://localhost:8080/accounts/17f904c1-806f-4252-9103-74e7a5d3e340"
```

### Get List All Account

```
curl -X "GET"  "http://localhost:8080/accounts"
```

### Transfer Balance

```
curl -X POST -H "Content-type: application/json" \
-d '{
	"from":"17f904c1-806f-4252-9103-74e7a5d3e340",
	"to":"3d253e29-8785-464f-8fa0-9e4b57699db9",
	"balance":"40"
}' 'http://localhost:8080/accounts'



```
