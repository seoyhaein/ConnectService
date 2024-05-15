### 확인사항
- gogoproto(https://github.com/cosmos/gogoproto) 사용 안함.  
- gogoproto 사용했을때는 ~.pb.go 파일 하나만 사용하면 되었지만 표준방식으로 사용하면 두개를 만들어야 한다.  
- 일단 성능적으로 낫다고 하지만, 안정적으로 standard proto 를 사용하기로 함.  
~~- pb 파일 생성은 window/linux 둘다 작성 함.~~ (Makefile 만들었음.)  

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
- API 해설 문서 만들기  
~~- 코드 정리(notion 참고, 진행 중)~~    
- Makefile 추가했는데 보강해야 하고, 리눅스에서 테스트 해야한다.  
- 테스트 코드 제작  