# CodeBank

### How to run
```
$ docker-compose up -d
```

### How enter in the image
```
$ docker exec -it appbank bash
```

### How to generate protocol bufers stub
```
$ docker exec -it appbank bash
$ make gen
```

### How to test with evans
```
$ docker exec -it appbank bash
$ evans -r repl -p=50051
```
