
// recriando uma brincadeira antiga
package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        switch {
        case i%3 == 0 && i%5 == 0:
            fmt.Println("Pin Pan")
        case i%3 == 0:
            fmt.Println("Pan")
        case i%5 == 0:
            fmt.Println("Pin")
        default:
            fmt.Println(i)
        }
    }
}
