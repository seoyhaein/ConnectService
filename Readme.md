### 확인사항
- gogoproto(https://github.com/cosmos/gogoproto) 사용 안함.  
- 일단 성능적으로 낫다고 하지만, 안정적으로 standard proto 를 사용하기로 함.  
- pb 파일 생성은 window/linux 둘다 작성 함.

### 설치사항
- protoc 설치
- protoc 의 설치는 https://github.com/protocolbuffers/protobuf/releases/tag/v26.1 이 링크에서 다운
- protoc-gen-go 설치
- https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.33.0 이 링크에서 설치
- 아래 grpc 설치
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
- 일단 go install 을 하면 gopath 기준으로 gopath/bin 에 설치가 된다. 나는 여기서 설치된것을 가져와서 해당 디렉토리에 넣어 두었다.
- 물론, 별도로 위에처럼 링크를 찾아서 해도 된다. https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc

### TODO
~~- 최신 버전으로 업그레이드~~
- 서비스 등록  
- 서버 초기 동작 확인  