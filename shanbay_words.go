package shanbay_words

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	token string
	dir   string
)

/*
Dir 分析文件所在目录
Token 获取的扇贝access_token
*/
func Do(Dir string, Token string) {
	dir = Dir
	token = Token
	list, err := ioutil.ReadDir(dir)
	checkErr(err)
	words := make([]string, 0)
	for _, v := range list {
		if v.IsDir() {
			return
		}
		name := v.Name()
		content, err := ioutil.ReadFile(dir + "/" + name)
		checkErr(err)
		ss := getWord(string(content))
		words = append(words, ss...)
	}
	//去重复
	result := removeRe(&words)
	var text string
	var havelear string
	for _, word := range result {
		if learnWord(word) {
			havelear += fmt.Sprintln(word)
		} else {
			break
		}
		text += fmt.Sprintln(word)
	}
	bb := []byte(havelear)
	err = ioutil.WriteFile("learnwords.txt", bb, 777)
	aa := []byte(text)
	err = ioutil.WriteFile("words.txt", aa, 777)
	checkErr(err)
}
func removeRe(words *[]string) []string {
	var a = make(map[string]struct{}, 0)
	for _, v := range *words {
		a[v] = struct {
		}{}
	}
	result := make([]string, 0)
	for k, _ := range a {
		result = append(result, k)
	}
	return result
}

type Res struct {
	Data Word   `json:"data"`
	Msg  string `json:"msg"`
}
type Word struct {
	Id int `json:"id"`
}

func getWord(content string) []string {
	ref, err := regexp.Compile("[a-zA-Z]{3,}")
	checkErr(err)
	words := ref.FindAllString(content, -1)
	return words
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("出错了")
		fmt.Printf(err.Error())
	}
}

func requestAPi(url string) string {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	checkErr(err)
	body, err := ioutil.ReadAll(res.Body)
	result := string(body)
	checkErr(err)
	return result

}
func learnWord(word string) bool {
	fmt.Println("记录单词" + word)
	return addWord(searchWord(word))
}
func searchWord(word string) int {
	var result = requestAPi("https://api.shanbay.com/bdc/search/?word=" + word)
	var r = new(Res)
	json.Unmarshal([]byte(result), r)
	fmt.Println(r.Data.Id)
	return r.Data.Id
}
func addWord(id int) bool {
	if id == 0 {
		return true
	}
	client := http.DefaultClient
	data := strings.NewReader("id=" + strconv.Itoa(id))
	req, err := http.NewRequest("POST", "https://api.shanbay.com/bdc/learning/", data)
	checkErr(err)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	checkErr(err)
	body, err := ioutil.ReadAll(res.Body)
	result := string(body)
	var r = new(Res)
	json.Unmarshal([]byte(result), r)

	if r.Msg == "SUCCESS" {
		println(result)
		return true
	}
	return false

}
