## Package of Basic Arrays Functions

This package contains basic functions for array manipulation.

## Installation

```bash
 go get github.com/gabrieldebem/array
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/gabrieldebem/array"
)

func main() {
    arr := []int{1,2,3,4}

    fmt.Println(array.Filter[int](arr, func(i int) bool {
        return i % 2 == 0
    }))
    // [2,4]
    
    fmt.Println(array.Map[int, string](arr, func(i int) string {
        return fmt.Sprintf("%d", i)
    }))
    // ["1","2","3","4"]

    fmt.Println(array.Reduce[int](arr, 0, func(acc int, i int) int {
        return acc + i
    }))
    // 10
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.   


## Look all this badges   \o/

<img src="https://shields.io/github/go-mod/go-version/gabrieldebem/array"> <img src="https://shields.io/github/license/gabrieldebem/array"> <img src="https://shields.io/github/v/release/gabrieldebem/array"> <img src="https://github.com/gabrieldebem/array/actions/workflows/ci.yml/badge.svg">
