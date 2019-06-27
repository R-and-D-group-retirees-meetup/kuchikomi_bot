package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	// BroadCastURL は BOT にてブロードキャスト配信をおこなうための API URL
	BroadCastURL = "https://api.line.me/v2/bot/message/broadcast"
	// ScrapingURL はスクレイピング先 URL を指定
	ScrapingURL = "https://en-hyouban.com/company/10006535807/kuchikomi"
)

var (
	// AccessToken は LINE Developers にて提供されるアクセストークン
	AccessToken = os.Getenv("LINE_ACCESS_TOKEN")
)

// Message は送信するメッセージを保持する構造体
type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// PostMessages は送信するメッセージ群を保持する構造体
type PostMessages struct {
	Messages []Message `json:"messages"`
}

// postBroadCast は BOT メッセージのブロードキャストメッセージを送信する
func postBroadCast(postMessages PostMessages) {
	jsonByte, err := json.Marshal(postMessages)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	req, err := http.NewRequest(
		"POST",
		BroadCastURL,
		bytes.NewBuffer(jsonByte),
	)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer resp.Body.Close()
	return
}

// getScrapingURL はランダムな値に応じた URL を取得する
func getScrapingURL() string {
	pageNum := rand.Intn(10)
	pageNumStr := "/?pagenum="
	switch pageNum {
	case 0:
	case 1:
		// 0, 1 の場合はページ番号が存在しないため、空文字を指定
		pageNumStr = ""
	default:
		pageNumStr = pageNumStr + strconv.Itoa(pageNum)
	}
	return ScrapingURL + pageNumStr
}

func getKuchikomiMsg() string {
	var kuchikomiMessages []string
	doc, err := goquery.NewDocument(getScrapingURL())
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	doc.Find("#kuchikomiList>.kuchikomi>.box>.comment").Each(func(i int, s *goquery.Selection) {
		comment := strings.TrimSpace(s.Text())
		kuchikomiMessages = append(kuchikomiMessages, comment)
	})
	var trimMsg []rune
	for _, runeV := range []rune(kuchikomiMessages[rand.Intn(10)]) {
		// 不要な空白を取り除く
		if runeV == 9 {
			continue
		}
		trimMsg = append(trimMsg, runeV)
	}
	trimSlice := strings.Split(strings.Replace(string(trimMsg), "口コミ投稿日", "\n口コミ投稿日", -1), "\n")
	// 先頭と末尾の不要な文言を削除
	return strings.Join(trimSlice[1:len(trimSlice)-1], "\n")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var postMesasges PostMessages
	postMesasges.Messages = append(postMesasges.Messages, Message{Type: "text", Text: getKuchikomiMsg()})
	postBroadCast(postMesasges)
}
