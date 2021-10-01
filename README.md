# gormuuid  [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/knesklab/util/blob/master/LICENSE)

A Postgres uuid[] field for gorm.io

## How to use

```go
type User struct {
	TagIDs         gormuuid.UUIDArray  `gorm:"type:uuid[]"` // Column will take {} as default
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