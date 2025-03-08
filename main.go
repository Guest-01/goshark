package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {
	// Raw Socket을 생성하기 위해 syscall 패키지를 이용하여 소켓을 생성합니다.
	socket, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_ALL)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Socket creation failed: %s\n", err)
		os.Exit(1)
	}
	defer syscall.Close(socket)

	// 캡쳐할 인터페이스를 지정하고 인덱스를 찾습니다.
	ifaceName := "eth0"
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Interface lookup failed: %s\n", err)
		os.Exit(1)
	}

	// 소켓을 인터페이스에 바인딩하기 위해 소켓 주소 구조체를 정의합니다.
	socketAddr := syscall.SockaddrLinklayer{
		Ifindex:  iface.Index,
		Protocol: htons(syscall.ETH_P_ALL),
	}

	// 소켓을 인터페이스에 바인딩합니다.
	if err := syscall.Bind(socket, &socketAddr); err != nil {
		fmt.Fprintf(os.Stderr, "Socket bind failed: %s\n", err)
		os.Exit(1)
	}

	// 패킷을 수신하기 위한 버퍼를 생성하고 for 루프를 돌며 패킷을 수신합니다.
	buffer := make([]byte, 65535)
	fmt.Println("Listening for packets...")
	for range 5 {
		n, _, err := syscall.Recvfrom(socket, buffer, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Recvfrom failed: %s\n", err)
			continue
		}
		fmt.Printf("\n--- Captured Packet (%d bytes) ---\n", n)
		fmt.Printf("% x\n", buffer[:n])
	}
}

func htons(i uint16) uint16 {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return binary.BigEndian.Uint16(b)
}
