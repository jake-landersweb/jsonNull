# Go Json Null

## About

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

One thing to note is the inclusion of the `Json` datatype. This allows for the marshalling and unmarshalling of the json datatype stored in MySQL. If you treat the json datatype of a regular string, then when encoding on object that uses this type into json will result in a 'double' encoding, causing extra '"' and "\" to be present.

## Usage

```shell
$ go get github.com/jake-landersweb/jsonNull
```

Change:

```go
type User struct {
    Name sql.NullString
    Date sql.NullTime
}
```

To:

```go
type User struct {
    Name jsonNull.String
    Date jsonNull.DateTime
}
```

## Available Types

- `Bool`
- `Byte`
- `DateTime`
- `Float64`
- `Int16`
- `Int32`
- `Int64`
- `String`
- `Json`

## Credits

The inspiration and most implementation was pulled from this excellent [Stack Overflow comment](https://stackoverflow.com/a/33072822).