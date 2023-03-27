package login

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/krile136/mineSweeper/funcs/aes"
	"github.com/krile136/mineSweeper/store"
)

type RequestBody struct {
	UserId       int    `json:"user_id"`
	OneTimeToken string `json:"one_time_token"`
}

type ResponseBody struct {
	ApiToken string `json:"api_token"`
}

func (l *Login) login() {

	fmt.Println("start")
	requestBody := RequestBody{
		UserId:       store.Data.Env.UserId,
		OneTimeToken: store.Data.Env.OneTimeToken,
	}

	fmt.Printf("UserId: %d\n", store.Data.Env.UserId)
	fmt.Printf("OneTimeToken: %s\n", store.Data.Env.OneTimeToken)

	// json文字列を生成する
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		l.getApiTokenCh <- err
		return
	}
	buf := bytes.NewBuffer(jsonString)
	method := "POST"
	url := "http://localhost:8081/api/authenticate"
	contentType := "application/json"

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		l.getApiTokenCh <- err
		return
	}
	req.Header.Add("Content-Type", contentType)

	fmt.Println("start post")
	// send post request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		l.getApiTokenCh <- err
		return
	}
	defer res.Body.Close()

	fmt.Printf("status code: %d\n", res.StatusCode)
	if res.StatusCode != 200 {
		l.getApiTokenCh <- errors.New(fmt.Sprintf("failed to get API Token. Status code : %d", res.StatusCode))
		return
	}

	fmt.Println("response ok")
	var resp map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resp)

	respData := resp["data"]
	fmt.Printf("data: %s\n", respData)

	if respData == nil {
		l.getApiTokenCh <- errors.New(fmt.Sprintf("failed to get API Token. response data is nil"))
		return
	}

	if respDataString, ok := respData.(string); ok {
		fmt.Println("start decrpt")
		// レスポンスデータは暗号化されているので復号化する
		respDataArray := strings.Split(respDataString, "|")
		iv, _ := hex.DecodeString(respDataArray[0])
		encrypted, _ := base64.StdEncoding.DecodeString(respDataArray[1])
		key, _ := hex.DecodeString(store.Data.Env.AesKey)
		decrypted, _ := aes.Decrypt(encrypted, key, iv)
		decryptedString := string(decrypted)

		// 復号したレスポンスからデータをパースする
		responseBody := ResponseBody{}
		err := json.Unmarshal([]byte(string(decryptedString)), &responseBody)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(responseBody.ApiToken)
		store.Data.Env.ApiToken = responseBody.ApiToken

		return
	}

	l.getApiTokenCh <- errors.New(fmt.Sprintf("failed to get API Token. Status code : %d", res.StatusCode))
	return
}
