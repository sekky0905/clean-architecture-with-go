# clean-architecture-with-go

The sample codes of Clean Architecture with Go.

## How to start
1 clone this repository.

```
git clone git@github.com:SekiguchiKai/clean-architecture-with-go.git
```

2 initialize the project by Makefile.

```
cd $GOPATH/src/github.com/SekiguchiKai/clean-architecture-with-go/server
make init
```

3 start application.

```
cd $GOPATH/src/github.com/SekiguchiKai/clean-architecture-with-go/server
make all
```

### Access Point

#### LIST
```
http://localhost:8080/v1/langs?limit=${num}
```

#### GET
```
http://localhost:8080/v1/langs/${id}
```

#### POST
```
http://localhost:8080/v1/langs
```

#### PUT
```
http://localhost:8080/v1/langs/${id}
```

#### DELETE
```
http://localhost:8080/v1/langs/${id}
```

## Caution
At this moment, you should configure DB and make .env file by yourself .

## Reference
https://blog.tai2.net/the_clean_architecture.html

https://nrslib.com/clean-architecture/

https://www.slideshare.net/pospome/go-80591000

https://qiita.com/hirotakan/items/698c1f5773a3cca6193e

https://postd.cc/golang-clean-archithecture/

https://qiita.com/kondei/items/41c28674c1bfd4156186
