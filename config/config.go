package config

import (
	"encoding/json"
	"flag"
	"os"
)

// 10/30 grpc config 파일 적용 추후
type Config struct {
	Address        string // grpc 서버 주소
	Silent         bool
	ConfigFilePath string

	// fetching data from config file
	Filename string
	// 버전 관련 Makefile 관련 내용 다루기
	// Version string
}

func DefaultConfig() *Config {
	return &Config{
		Address:        ":50052",
		Silent:         true,
		ConfigFilePath: "./config/config.json",
	}
}

func (c *Config) LoadFromFile() error {
	file, err := os.ReadFile(c.ConfigFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}

/*
	아래 메서드에서 리시버( (c *Config) )가 포인터냐 포인터가 아니냐에 따라 메서드를 호출하는 방식이 달라진다.
	지금은 포인터로 리시버를 만들어 주었기 때문에 server.go 에서 아래와 같이 작성해줄 수 있다.

	c.RegisterFlags(fs)

	만약 리시버를 value 로 만들어 줄 경우( (c Config) ) server.go 에서는 아래와 같이 작성해줄 것이다.
    (*c).RegisterFlags(fs)
*/

func (c *Config) RegisterConfig(fs *flag.FlagSet) (*Config, error) {

	var conf = c

	fs.StringVar(&c.Address, "u", c.Address, "server address")
	fs.BoolVar(&c.Silent, "silent", c.Silent, "Log nothing to stdout/stderr")
	fs.StringVar(&c.ConfigFilePath, "path", c.ConfigFilePath, "config file path")

	// 최소 디폴트 값의 url 주소를 가지고 온다.
	// TODO 향후 grpc 및 기타 tcp 서버 관련 정보를 configfile 에 담는다.
	file, err := os.ReadFile(c.ConfigFilePath)

	if err != nil {
		return nil, err
	}

	// '=' 기호로 교체되었다.
	err = json.Unmarshal(file, conf)

	if err != nil {
		return nil, err
	}

	/*
		리턴을 표현할 때

		// 명시적으로 str2 를 리턴값의 이름을 지어줌.
		func FunctionA(str1 string) (str2 string)  {

		str2 = str1
		return
		}

		이렇게 두는 것이 return str2 라고 명시하는 것보다 더 효과적이라는 test 결과를 예전에 웹 상에서 본적이 있다.
		참고 삼아 한번 적어 보았다.
	*/
	return conf, nil
}
