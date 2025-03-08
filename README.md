# goshark
> simple packet sniffer in go

Go를 학습하기 위한 프로젝트 2번째. (1번째는 [port-scanner-go](https://github.com/Guest-01/port-scanner-go))

외부 패키지를 사용하지 않고, Wireshark와 비슷하게 패킷을 캡쳐해보는 프로젝트입니다.

## 알게된 사실

### Go에는 system call을 제공하는 패키지가 있어서 직접 시스템콜을 호출할 수 있다.

그래서 `socket`을 만들 때 원래 일반 TCP/UDP 소켓을 만든다면, 추상화된 `net` 패키지 내의 함수를 이용하였겠지만 지금 요구사항에서는 더 상위 레이어(Layer 2)에 바인딩한 소켓을 만들어야 하기 때문에 `syscall` 패키지를 활용해 직접 소켓을 만든다.

```go
// 일반적인 TCP/UDP 소켓으로 서버를 만든다면...
listener, err := net.Listen("tcp", ":8080")

// 시스템콜을 이용한 Socket 직접 생성
socket, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_ALL)))
```

(작성 중)

## TODO
- 커맨드 인자로 인터페이스 지정하기
- 목적지 IP 등으로 조건부 캡처하기 (=필터링)
- 함수 분리 및 테스트 작성
