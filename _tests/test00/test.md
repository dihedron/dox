This is the some sample text with **bold** and _italics_.

Unfortunately, Markdown does not support underline.

This is a picture of Homer Simpson.

![homer](images/homer.jpg?width=0.75&label=fig:homer&caption=A%20sleeping%20Homer%20Simpson "Homer Simpson")

You can even set a reference to your code via an internal ref (TODO).

Now let's show a piece of golang code:

```go
package main 

import "fmt"

func main() {
    fmt.Println("hallo, world!")
}
```

and compare it with C:

```c
#include <stdio.h>

int main() {
    printf("hello, world");
    return 0;
}
```
