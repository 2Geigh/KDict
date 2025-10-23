package main

import (
	"KDict/src/config"
	"KDict/src/routes"
	"log"
	"net/http"
)

func main() {

	// Load configuration from .env
	config.LoadConfig()

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Initialize routes
	routes.RegisterDictRoutes()

	// Start server
	log.Fatal(http.ListenAndServe(":3000", nil))

}

// For item in dictSearch.Results
// {
// 	한국어 기초사전 개발 지원(Open API) - 사전 검색
// 	5
// 	[
// 		{
// 			72461
// 			한자
// 			0
// 			漢字
// 			한ː짜
// 			중급
// 			명사
// 			https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=72461
// 			{
// 				1
// 				중국에서 만들어 오늘날에도 쓰고 있는 중국 고유의 문자.
// 			}
// 		}

// 		{
// 			85621
// 			한자리
// 			0
// 			한자리
// 			고급
// 			명사
// 			https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=85621
// 			{
// 				2
// 				중요하거나 높은 직위. 또는 어느 한 직위.
// 			}
// 		}

// 		{93420 한자어 0 漢字語 한ː짜어 고급 명사 https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=93420 {1 한자에 기초하여 만들어진 말.}} {85622 한자리하다 0  한자리하다  동사 https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=85622 {1 중요하거나 높은 직위에 오르다.}} {88977 한자음 0 漢字音 한ː짜음  명사 https://krdict.korean.go.kr/kor/dicSearch/SearchView?ParaWordNo=88977 {1 한자의 발음이나 소리.}}]}
