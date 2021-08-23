package finger

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/twmb/murmur3"
)

func Mmh3Hash32(raw []byte) string {
	var h32 hash.Hash32 = murmur3.New32()
	_, err := h32.Write([]byte(raw))
	if err == nil {
		return fmt.Sprintf("%d", int32(h32.Sum32()))
	} else {
		//log.Println("favicon Mmh3Hash32 error:", err)
		return "0"
	}
}

func StandBase64(braw []byte) []byte {
	bckd := base64.StdEncoding.EncodeToString(braw)
	var buffer bytes.Buffer
	for i := 0; i < len(bckd); i++ {
		ch := bckd[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()

}

func favicohash(host string) string {
	timeout := time.Duration(8 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Timeout:   timeout,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse /* 不进入重定向 */
		},
	}
	resp, err := client.Get(host)
	if err != nil {
		//log.Println("favicon client error:", err)
		return "0"
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//log.Println("favicon file read error: ", err)
			return "0"
		}
		return Mmh3Hash32(StandBase64(body))
	} else {
		return "0"
	}
}
