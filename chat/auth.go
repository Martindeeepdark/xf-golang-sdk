package chat

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

// getAuthorizationURL 获取授权URL
// see https://www.xfyun.cn/doc/spark/general_url_authentication.html.
func getAuthorizationURL(hostURL, apiKey, apiSecret string) (*url.URL, error) {
	urlAddr, err := url.Parse(hostURL)
	if err != nil {
		return nil, err
	}

	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 MST")

	// Get signature_origin
	signatureOrigin := fmt.Sprintf("host: %s\ndate: %s\nGET %s HTTP/1.1", urlAddr.Host, date, urlAddr.Path)

	// Calculate signature
	key := []byte(apiSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signatureOrigin))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Calculate authorization_origin
	authorizationOrigin := fmt.Sprintf("api_key=%q,algorithm=%q,headers=%q,signature=%q",
		apiKey, "hmac-sha256", "host date request-line", signature)

	// Encode authorization_origin
	authorization := base64.StdEncoding.EncodeToString([]byte(authorizationOrigin))

	// Construct authorization URL
	q := urlAddr.Query()
	q.Set("authorization", authorization)
	q.Set("date", date)
	q.Set("host", urlAddr.Host)
	urlAddr.RawQuery = q.Encode()

	return urlAddr, nil
}
