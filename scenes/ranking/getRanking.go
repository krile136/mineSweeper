package ranking

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/krile136/mineSweeper/funcs/aes"
	"github.com/krile136/mineSweeper/store"
)

// 自分のハイスコアとランキング上位５人を取得する
func (r *Ranking) getRanking() {
	log.Print("start get ranking")

	var method string = "GET"
	var url string = fmt.Sprintf("http://localhost:8081/api/score/%d", store.Data.Env.UserId)
	var contentType string = "application/json"
	var Authorization string = "Bearer " + store.Data.Env.ApiToken

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		r.getRankingIndexCh <- err
		return
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", Authorization)

	// send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		r.getRankingIndexCh <- err
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		r.getRankingIndexCh <- errors.New(fmt.Sprintf("failed to get Ranking Data. Status code : %d", res.StatusCode))
		return
	}

	var resp map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resp)

	respData := resp["data"]
	if respData == nil {
		r.getRankingIndexCh <- errors.New(fmt.Sprintf("failed to get Ranking Data. response data is nil"))
	}

	// レスポンスデータは暗号化されているので復号化する
	if respDataString, ok := respData.(string); ok {
		fmt.Println("start decrpt")
		respDataArray := strings.Split(respDataString, "|")
		iv, _ := hex.DecodeString(respDataArray[0])
		encrypted, _ := base64.StdEncoding.DecodeString(respDataArray[1])
		key, _ := hex.DecodeString(store.Data.Env.AesKey)
		decrypted, _ := aes.Decrypt(encrypted, key, iv)
		decryptedString := string(decrypted)

		// 復号したレスポンスデータから中身をパースする
		err := json.Unmarshal([]byte(string(decryptedString)), &r.resp)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
