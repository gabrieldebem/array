## Package of Basic Arrays Functions

This package contains basic functions for arrays manipulation.

## Installation

```bash
 go get github.com/gabrieldebem/arrays
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/gabrieldebem/arrays"
)

func main() {
    arr := []int{1,2,3,4}

    fmt.Println(arrays.Filter[int](arr, func(i int) bool {
        return i % 2 == 0
    }))
    // [2,4]
    
    fmt.Println(arrays.Map[int, string](arr, func(i int) string {
        return fmt.Sprintf("%d", i)
    }))
    // ["1","2","3","4"]

    fmt.Println(arrays.Reduce[int](arr, 0, func(acc int, i int) int {
        return acc + i
    }))
    // 10
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.   