package gameover

import (
	"bytes"
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

type RequestParams struct {
	UserId int `json:"user_id"`
	Score  int `json:"score"`
}

type RequestBody struct {
	Data string `json:"data"`
}


func (g *Gameover) postScore() {
	fmt.Println("start post score")
	log.Print("start post score")

	// リクエストパラメーターを作成する
	requestParams := RequestParams{
		UserId: store.Data.Env.UserId,
		Score:  store.Data.CurrentScore,
	}
	jsonString, err := json.Marshal(requestParams)
	if err != nil {
		g.postScoreCh <- err
		return
	}

	// リクエストパラメータを暗号化
	requestBody := RequestBody{
		Data: aes.Encrypt(string(jsonString)),
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		g.postScoreCh <- err
		return
	}

	// POSTの準備をする
	method := "POST"
	url := "http://localhost:8081/api/score"
	buf := bytes.NewBuffer(data)
	contentType := "application/json"
	Authorization := "Bearer " + store.Data.Env.ApiToken

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		g.postScoreCh <- err
		return
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", Authorization)

	// POSTリクエスト実行
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		g.postScoreCh <- err
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		g.postScoreCh <- errors.New(fmt.Sprintf("failed to post score. Status code : %d", res.StatusCode))
		return
	}

	// レスポンスからdataを取得
	fmt.Println("response ok")
	var resp map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resp)

	respData := resp["data"]
	fmt.Printf("data: %s\n", respData)

	if respData == nil {
		g.postScoreCh <- errors.New(fmt.Sprintf("failed to post score. response data is nil"))
		return
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
		err := json.Unmarshal([]byte(string(decryptedString)), &g.resp)
		if err != nil {
			fmt.Println(err)
		}

		return
	}
}
