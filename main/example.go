package main

import (
	"fmt"
	"github.com/ramonibz/totp/totp"
)

const RANDOM_SEED	=	"YWCHKEKW6FVBLCLT2RJB4NHE5NZD6J4MNOJZL7GFELN3YTPJX537RZBIRA4MCBWE"

func main() {
		fmt.Println(totp.GetTotpCode(RANDOM_SEED))
}
