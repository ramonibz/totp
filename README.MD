### TOTP Algorithm implementation
This simple program is an implementation of the [TOTP](https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm) (Time based One Time Password) algorithm

`go get github.com/ramonibz/totp/totp`

Basic usage example:
```go
import "github.com/ramonibz/totp/totp"

func main() {
    totp.GetTotpCode(seed)
}

```
