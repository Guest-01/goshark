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

// 하지만 지금처럼 Raw Socket을 만든다면 시스템콜을 이용한 Socket 직접 생성
socket, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
```

### 일반적인 x86 시스템과 네트워크 규약간에 엔디안 차이로 인해 패킷에 Byte를 담을 때는 변환이 필요하다

Intel x86 시스템에서는 리틀 엔디안을 사용하는 반면, 네트워크에서는 빅 엔디안을 쓰는 것이 규칙이다. 따라서 같은 값도(예를 들어 `syscall.ETH_P_ALL`) 시스템 내에서 사용될 때는(위의 예제처럼 소켓을 생성한다던지) 그대로 써도 되지만, 패킷에 담을 때는 빅 엔디안으로 변환해주어야한다. 그리고 이런 변환 과정을 Host to Network Short(`htons`)라고 부른다.

```go
// 아래 정보(Protocol)는 이더넷 헤더에 담기기 때문에 네트워크 규약에 맞춰 빅엔디안으로 변경 필요
socketAddr := syscall.SockaddrLinklayer{
    Ifindex:  iface.Index,
    Protocol: htons(syscall.ETH_P_ALL),
}
```

## TODO
- 커맨드 인자로 인터페이스 지정하기
- 목적지 IP 등으로 조건부 캡처하기 (=필터링)
- 함수 분리 및 테스트 작성
