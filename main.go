package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	conf "github.com/seoyhaein/ConnectService/config"
	"github.com/seoyhaein/ConnectService/v1rpc"
)

var (
	Version = "0.0.0"
	err     error
)

func main() {
	// 기본 설정 초기화
	newConfig := conf.DefaultConfig()
	// 플래그 세트 생성
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// 플래그 정의
	fs.StringVar(&newConfig.Address, "u", newConfig.Address, "server address")
	fs.BoolVar(&newConfig.Silent, "silent", newConfig.Silent, "Log nothing to stdout/stderr")
	fs.StringVar(&newConfig.ConfigFilePath, "path", newConfig.ConfigFilePath, "config file path")

	// 플래그 파싱
	fs.Parse(os.Args[1:])
	// TODO 향후 조정 하자.
	if len(fs.Args()) < 2 {
		// 설정 파일에서 값 로드
		if err := newConfig.LoadFromFile(); err != nil {
			log.Fatalf("Failed to load config from file: %v\n", err)
		}
	}

	if !newConfig.Silent {
		fmt.Println("사용자 입력 파라미터", newConfig.Address)
		fmt.Println("config 파일로부터 읽은 데이터", newConfig.Filename)
		// TODO 버전 표시는 추후 명령어로 10/30
		fmt.Println("현재 Version", Version)
	}

	if err = v1rpc.Server(); err != nil {
		log.Fatalln(err)
	}
}
