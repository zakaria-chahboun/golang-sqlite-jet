## Golang: new SQLite driver + Jet generator

In this example i didn't use ORMs, Instead i used a compiled go files from jet project: https://github.com/go-jet/jet

I also used this new SQLite driver with No CGO from https://gitlab.com/cznic/sqlite

## Installation

```bash
go mod tidy
```

To generate go files for your *schema/query*, run:

```
jet -source=sqlite -dsn="./db/database.db" -path=./gen
```

Then run the app:
```
go run .
```

Follow me on [twitter](https://twitter.com/zaki_chahboun) 
------
2023-01-11 10:53

