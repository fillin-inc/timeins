# `Time in seconds` package

[![Build Status](https://travis-ci.org/fillin-inc/timeins.svg?branch=master)](https://travis-ci.org/fillin-inc/timeins)

`timeins` package has been developed to return time up to seconds with the JSON API using Go.

If the `time.Time` type is included in the structure for response, when returning it to JSON, it returns the value up to millisecond as a string. If you specify the `timeins.Time` type in this package as a structure, you can return the time until seconds as a string.

## Example

``` golang
package main

import (
  "encoding/json"
  "fmt"
  "time"

  "github.com/fillin-inc/timeins"
)

type res struct {
  CreatedAt timeins.Time
}

func main() {
  r := res{
    CreatedAt: timeins.Time(time.Now()),
  }

  marshaled, _ := json.Marshal(r)
  fmt.Println(string(marshaled))
}
```

## Development

This project uses `.editorconfig` to maintain consistent coding styles across different editors and IDEs. Most modern editors support EditorConfig automatically or through plugins.

## License

MIT License. see [LICENSE](https://github.com/fillin-inc/timeins/blob/master/LICENSE) file.
