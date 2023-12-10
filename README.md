# NOTICE

__Source code has been moved to [Sourcehut](https://git.sr.ht/~nullevoid/mdb-go) and this repository has been archived.__

# MDB

MDB is a simple JSON file database written in Go.

MDB is meant to be used for toy projects as a simple alternative to loading
in large database systems. It is a hobby project at best.

## Usage

```go

import (
    "fmt"

    "github.com/MorganPeterson/mdb"
)

func main() {
	db := NewDatabase()
	
	err := Load(db, "db.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	id, err := Put(db, "hello 1")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = Update(db, id, "hello 1 updated")
	if err != nil {
		fmt.Println(err.Error())
	}

	dId, err := Put(db, "hello 2")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = Delete(db, dId)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = Commit(db, "db.json")
}
```

Resulting JSON file:

```json
{
  "data": {
    "2144D9F4-B57B-44E3-90CB-4F26BB696E77": {
      "created": "2023-11-24 16:11:05.880023432 +0000 UTC",
      "edited": "2023-11-24 16:11:05.880027792 +0000 UTC",
      "doc": "hello 1 updated"
    }
  }
}
```
