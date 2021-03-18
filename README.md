# Key-Value Storage

## Installation
```
git clone https://github.com/maratishimbaev/short-url-service.git
```

## Usage
Run server:
```
docker-compose up
```

## API
Method | Path | Parameters | Description
------ | ---- | ---------- | -----------
POST | / | body | Create or update key
DELETE | /{key:[a-zA-Z]+} | - | Delete key
GET | /{key[a-zA-Z]+} | - | Get key value
GET | / | - | Get list of keys

Create key body example:
```json
{
  "key": "a",
  "value": "1"
}
```
