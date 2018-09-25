# clean-architecture-with-go

The sample codes of Clean Architecture with Go.

## How to start
1 clone this repository.

```
cd ${Your_Working_Directory}
git clone git@github.com:SekiguchiKai/clean-architecture-with-go.git
```

2 initialize the project by Makefile.

```
cd clean-architecture-with-go/server
make init
```

3 start application.

```
docker-compose up -d
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

## Reference

エリック・エヴァンス(著)、 今関 剛 (監修)、 和智 右桂 (翻訳) (2011/4/9)『エリック・エヴァンスのドメイン駆動設計 (IT Architects’Archive ソフトウェア開発の実践)』 翔泳社

Robert C.Martin (著)、 角 征典 (翻訳)、 高木 正弘 (翻訳)　(2018/7/27)『Clean Architecture 達人に学ぶソフトウェアの構造と設計』 KADOKAWA

アラン・シャロウェイ (著)、 ジェームズ・R・トロット (著)、 村上 雅章 (翻訳) (2014/3/11)『オブジェクト指向のこころ (SOFTWARE PATTERNS SERIES)』 丸善出版

結城 浩 (2004/6/19)『増補改訂版Java言語で学ぶデザインパターン入門』 ソフトバンククリエイティブ

InfoQ.com、徳武 聡(翻訳) (2009年6月7日) 『Domain Driven Design（ドメイン駆動設計） Quickly 日本語版』 InfoQ.com Domain Driven Design（ドメイン駆動設計） Quickly 日本語版

https://blog.tai2.net/the_clean_architecture.html

https://nrslib.com/clean-architecture/

https://www.slideshare.net/pospome/go-80591000

https://qiita.com/hirotakan/items/698c1f5773a3cca6193e

https://postd.cc/golang-clean-archithecture/

https://qiita.com/kondei/items/41c28674c1bfd4156186

https://golang.org/pkg/database/sql

