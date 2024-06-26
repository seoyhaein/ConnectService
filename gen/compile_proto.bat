@echo off
REM 첫 번째 인수로 프로토콜 버퍼 파일 이름을 받습니다.
set PROTO_FILE=%1

REM protoc.exe 경로
set PROTOC_PATH="./protoc.exe"
REM 프로토콜 버퍼 폴더의 경로
set PROTO_PATH="../protos"
REM 생성된 파일을 저장할 디렉토리의 경로
set OUT_DIR="../protos"

REM 실행 Protocol Buffers 코드 생성
%PROTOC_PATH% --proto_path=%PROTO_PATH% --proto_path=%PROTO_PATH%/include --go_out=%OUT_DIR% --go_opt=paths=source_relative %PROTO_PATH%/%PROTO_FILE%

REM 실행 gRPC 코드 생성
%PROTOC_PATH% --proto_path=%PROTO_PATH% --proto_path=%PROTO_PATH%/include --go-grpc_out=%OUT_DIR% --go-grpc_opt=paths=source_relative %PROTO_PATH%/%PROTO_FILE%
