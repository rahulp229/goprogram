package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TeamData struct {
	AssestId string `json:"guid"`
	Teams    []ExtraInfo
}
type ExtraInfo struct {
	SportType  string `json:"sports_type"`
	LeagueName string `json:"league_name"`
}

func main() {

	extrainfo := getExtraInfo()
	extrainfo1 := getExtraInfo1()
	extrainfo.Teams = extrainfo.Teams[:1]

	fmt.Printf("extraInfo %+v:", extrainfo)
	fmt.Printf("extraInfo1 %+v:", extrainfo1)

	//	log.Printf("%v\n", data)
}

func getExtraInfo() TeamData {
	resp, _ := http.Get("https://cbd46b77.cdn.cms.movetv.com/cms/api/programs/0df55b72e61e42f794c68f0280e1f4f4")
	//	resp1, _ := http.Get("https://storage.googleapis.com/everledger-platform-uat/b85dcada-c414-477c-ba4e-e4173a12df24/1d6aff57-159f-4e2a-99ec-88ebfbbcf00c/8fd2de37-5664-47df-9fbc-8f4bf6ccfe71/pink-neon.jpg")
	r, _ := ioutil.ReadAll(resp.Body)

	var td TeamData
	err := json.Unmarshal(r, &td)
	if err != nil {
		log.Fatal(err)
	}

	return td
}

func getExtraInfo1() string {
	//resp, _ := http.Get("https://cbd46b77.cdn.cms.movetv.com/cms/api/programs/0df55b72e61e42f794c68f0280e1f4f4")
	f, _ := os.Open("https://storage.googleapis.com/everledger-platform-uat/b85dcada-c414-477c-ba4e-e4173a12df24/1d6aff57-159f-4e2a-99ec-88ebfbbcf00c/8fd2de37-5664-47df-9fbc-8f4bf6ccfe71/pink-neon.jpg?Expires=1639160627&GoogleAccessId=everledger-uat-bucket%40everledger-internal.iam.gserviceaccount.com&Signature=Mjaurv9yljsjbJ4YYDPgQSFHD8pSLuJLtZxG8oAPDIyZkLpbx8KlSNlJ4nYBP2WVNdGlyD5rLEmX63KH6zoN1imWqLxV2rfng5llJ8GQgQLX2IxfzaY7keZWj%2F%2FfBDVUoFmRXOuoQECDH6pEp2JkXANd02%2BR5qHZssMXyPXE%2FJAG3OvtEmlg7zOlvlouM%2FvnBUEkat6YBUqEGwxcvt2AWbEQDNIUZAaHuqHwOqHGtR085N%2FTnKWE8QHxAQ26yOcfci9mM6BRGLY21yNycAOZhZUkg5rb6IXNTOtV2qdgJjL5LHvFA75Djw6YtohtJcKhGaZukyQKO8bEGIlyNQ1Phw%3D%3D")
	stats, _ := f.Stat()  
	fmt.Println("extraInfo1 ", stats)
os.OpenFile()
	resp1, _ := http.Get("https://storage.googleapis.com/everledger-platform-uat/b85dcada-c414-477c-ba4e-e4173a12df24/1d6aff57-159f-4e2a-99ec-88ebfbbcf00c/8fd2de37-5664-47df-9fbc-8f4bf6ccfe71/pink-neon.jpg?Expires=1639160627&GoogleAccessId=everledger-uat-bucket%40everledger-internal.iam.gserviceaccount.com&Signature=Mjaurv9yljsjbJ4YYDPgQSFHD8pSLuJLtZxG8oAPDIyZkLpbx8KlSNlJ4nYBP2WVNdGlyD5rLEmX63KH6zoN1imWqLxV2rfng5llJ8GQgQLX2IxfzaY7keZWj%2F%2FfBDVUoFmRXOuoQECDH6pEp2JkXANd02%2BR5qHZssMXyPXE%2FJAG3OvtEmlg7zOlvlouM%2FvnBUEkat6YBUqEGwxcvt2AWbEQDNIUZAaHuqHwOqHGtR085N%2FTnKWE8QHxAQ26yOcfci9mM6BRGLY21yNycAOZhZUkg5rb6IXNTOtV2qdgJjL5LHvFA75Djw6YtohtJcKhGaZukyQKO8bEGIlyNQ1Phw%3D%3D")
	r, _ := ioutil.ReadAll(resp1.Body)

	var td string
	err := json.Unmarshal(r, &td)
	if err != nil {
		log.Fatal(err)
	}

	return td
}
