## Introduction
- 友人との通話時の話題を事前に募集し、それを通話中にランダムに表示したい

## Commands

### up
```
docker-compose up -d
```

### exec
```
# NOT bash
docker-compose exec db sh
```

```
docker-compose exec app go run main.go
```

### stop
```
docker-compose stop
```

### down
```
docker-compose down
```

### use app
```
# add user
curl -X POST http:/localhost:10000/users -d '{ "name": "Create test name", "email": "Create test email"}'
curl -X PUT http:/localhost:10000/users/0 -d '{ "name": "Put test name", "email": "Put test email"}'
```

get ACCESS_TOKEN
```
curl http:/localhost:10000/auth
```

use Bearer token
```
curl --request GET \
--url http:/localhost:10000/users \
--header 'authorization: Bearer ACCESS_TOKEN'
```
