package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"
)

const DEFAULT_TIME_SECONDS = 30

func GetTotpCode(seed string) string {

	seed = strings.TrimSpace(seed)
	if n := len(seed) % 8; n != 0 {
		seed = seed + strings.Repeat("=", 8-n)
	}

	seed = strings.ToUpper(seed)

	decodedKey, err := base32.StdEncoding.DecodeString(seed)
	if err != nil {
		panic(err)
	}

	t := time.Now().UTC().Unix()
	counter := uint64(math.Floor(float64(t)/ float64(DEFAULT_TIME_SECONDS)))
	buf := make([]byte, 8)
	mac := hmac.New(sha1.New, decodedKey)
	binary.BigEndian.PutUint64(buf, counter)

	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))
	mod := int32(value % int64(math.Pow10(6)))


	f := fmt.Sprintf("%%0%dd", 6)
	return fmt.Sprintf(f, mod)
}
