# Go Json Null

**A quick and dirty package for handling SQL and JSON typed values inside of Go.**

One of the problems with working with SQL inside of Go is the fact that columns can be null. While this is not a huge problem and can be easily worked around with using sql types such as `sql.NullString` or `sql.NullInt32`, it becomes more difficult to use these raw types when serializing into json.

For example, if my data struct looks as follows:

```go
type User struct {
    Name sql.NullString
    Date sql.NullTime
}
```

And I were to encode this to json with the following:

```go
u := User { Name: "Jake", Date: time.Now() }
encoded, err := json.Marshal(&u)
```

Then the resulting json would look like this:

```json
{
    "Name": {
        "String": "Jake",
        "Valid": true
    },
    "Date": {
        "Time": "0001-01-01T00:00:00Z",
        "Valid": true
    }
}
```

While this is not necessarily the worst thing ever, I like my json to be as flat as possible. This package solves that by giving types wrapped around the base `sql.Null` types that implement custom json marshal functions. 

## Credits

The code was pulled from this excellent [Stack Overflow comment](https://stackoverflow.com/a/33072822), I just decided to package it up for use in my projects.