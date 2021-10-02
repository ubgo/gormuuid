# gormuuid  [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/knesklab/util/blob/master/LICENSE)

A Postgres uuid[] field for gorm.io

## Installation
```
go get github.com/ubgo/gormuuid
```

## How to use

```go
type User struct {
  TagIDs gormuuid.UUIDArray  `gorm:"type:uuid[]"` // Column will take {} as default
}

var tagIds gormuuid.UUIDArray = []uuid.UUID{uuid.New(), uuid.New()}
db.Create(&User{
  TagIDs: tagIds,
})
```

## Select Query
```go
var users []User
tagId, _ := uuid.Parse("c4ba81a1-9e57-4e60-b811-2860136ab803")
db.Model(&User{}).Where("tag_ids && ?", pq.Array([]uuid.UUID{tagId})).Find(&users)
```


## Contribute

If you would like to contribute to the project, please fork it and send us a pull request.  Please add tests
for any new features or bug fixes.

## Stay in touch

* Author - [Aman Khanakia](https://twitter.com/mrkhanakia)
* Website - [https://khanakia.com](https://khanakia.com/)

## License

goutil is [MIT licensed](LICENSE).
