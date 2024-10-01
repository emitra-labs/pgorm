# pgorm

This library combines PostgreSQL and GORM for simplified database interactions.

## Quickstart

Make sure you've set these required environment variables:

```
PGORM_URL="postgres://username:password@localhost:5432/database_name?sslmode=disable"
```

Here's a simple code to demonstrates how to use it:

```go
package main

import (
    "fmt"

    "github.com/emitra-labs/pgorm"
)

func main() {
    // Open database connection
    pgorm.Open()

    var result int64

    err := pgorm.DB.Raw("SELECT 1 + 1").Scan(&result).Error
    if err != nil {
        panic(err)
    }

    fmt.Println("result:", result)

    // Close the connection
    pgorm.Close()
}
```