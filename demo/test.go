package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/smantel-ch/openvpn-go"
)

func main() {

	// fork from: https://github.com/smantel-ch/openvpn-go

	username := "mirror"
	password := "mirror"
	config := "client\nproto tcp\nremote hub.gede100.com 443\ndev tun\nresolv-retry infinite\nnobind\npersist-key\npersist-tun\nremote-cert-tls server\nverify-x509-name server_YfEnafJGEN3AFYjD name\nauth SHA256\nauth-user-pass\ncipher AES-128-GCM\ntls-client\ntls-version-min 1.2\ntls-cipher TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256\nverb 3\n\n\n\n\n## Custom configuration ##\n\n## end ##\n\n<ca>\n-----BEGIN CERTIFICATE-----\nMIIB3zCCAYWgAwIBAgIUGwhXwLjDwHJVt2jT4JoV/sTcwi4wCgYIKoZIzj0EAwIw\nIDEeMBwGA1UEAwwVb3Zwbl9kbEVNQ3ZGeWxDaWNoRVMxMCAXDTI1MTIxMjE4NDgy\nOVoYDzIxMjUxMTE4MTg0ODI5WjAgMR4wHAYDVQQDDBVvdnBuX2RsRU1DdkZ5bENp\nY2hFUzEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARsfyYvsDJQ9OJa4MaNrtcR\n5ilaUbeMFrk5eIIxrldD+BiuqAFoV9IBUicCkun8KyQ+HMj575t8FEXBVXfsuFNc\no4GaMIGXMAwGA1UdEwQFMAMBAf8wHQYDVR0OBBYEFOnrOqtShB58zWNnp8HkT/w4\nYMJAMFsGA1UdIwRUMFKAFOnrOqtShB58zWNnp8HkT/w4YMJAoSSkIjAgMR4wHAYD\nVQQDDBVvdnBuX2RsRU1DdkZ5bENpY2hFUzGCFBsIV8C4w8ByVbdo0+CaFf7E3MIu\nMAsGA1UdDwQEAwIBBjAKBggqhkjOPQQDAgNIADBFAiAgh7UWRrJUCcj8PPOB38WB\nQj1ivecbMIfpJR26BZYkpQIhALC9hUtLWzOKHnmXG76H518Ek+vzqYioYUIROw7j\neNaX\n-----END CERTIFICATE-----\n</ca>\n<cert>\n-----BEGIN CERTIFICATE-----\nMIIB5jCCAYygAwIBAgIRAPulJ6xHJdWx0aJ2ZN7+LOswCgYIKoZIzj0EAwIwIDEe\nMBwGA1UEAwwVb3Zwbl9kbEVNQ3ZGeWxDaWNoRVMxMB4XDTI1MTIxMjE4NTAyOFoX\nDTI2MTIxMjE4NTAyOFowGjEYMBYGA1UEAwwP6ZWc5YOP5pyN5Yqh5ZmoMFkwEwYH\nKoZIzj0CAQYIKoZIzj0DAQcDQgAEPgcogcD/CkXExJOsmFHpdY7mWP2ZgPsaLJ1N\nFam25Y5/uPIThff1ds7g0TfTmiVL3mQ7u47Ai991j5Wn/gJ0QqOBrDCBqTAJBgNV\nHRMEAjAAMB0GA1UdDgQWBBRM2K1QyO3d+DNmBgZRZpMtwnXNlzBbBgNVHSMEVDBS\ngBTp6zqrUoQefM1jZ6fB5E/8OGDCQKEkpCIwIDEeMBwGA1UEAwwVb3Zwbl9kbEVN\nQ3ZGeWxDaWNoRVMxghQbCFfAuMPAclW3aNPgmhX+xNzCLjATBgNVHSUEDDAKBggr\nBgEFBQcDAjALBgNVHQ8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhANN1jXmIzLMu\nZ6J4rit8CUM2tupgUXjOTilE0OstD8zQAiBCtHiBtZnSxginPnhE2uHG3n/q2tDS\nDRnTmV7Feu7lRA==\n-----END CERTIFICATE-----\n</cert>\n<key>\n-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg9eZn2XPrCUkcDaL6\n21oxZwxzwqpQ0hBwGalwykSOt/ehRANCAAQ+ByiBwP8KRcTEk6yYUel1juZY/ZmA\n+xosnU0Vqbbljn+48hOF9/V2zuDRN9OaJUveZDu7jsCL33WPlaf+AnRC\n-----END PRIVATE KEY-----\n</key>\n<tls-crypt>\n#\n# 2048 bit OpenVPN static key\n#\n-----BEGIN OpenVPN Static key V1-----\n87485beb1e7fb67c5442c8a43bbba0c3\n3a0a4e78996236cc235932ae89111a0e\n8c3ed939417074f106c5b75b57fccd95\nd2cce840f6329f197407999675e28a7e\n2742fcfece297d2cad7c92eae895a951\n6d49589ac3c51089dd374faa901645eb\n82741c2248c6598e0c367473a0251b8d\ncccf40d9700d82d1db252f92531be842\ne7e5a10c5ac4e415c8a59c21f648cc7c\n2d44cacf54a3e44329d56304a267c721\n05426a5bc37e22fc51b1d6e6c452b7a2\n42b321b97e612c29b0e82c7a1a6d2acf\ncfc4d2771f386f844d7e135e859abebf\n6a5d94cee9fbd62a0554170bcf392597\nc4d0cd1b5407814474e3401a67375b36\nf221146111a6fb060ef00ffd5948e119\n-----END OpenVPN Static key V1-----\n</tls-crypt>\n"

	openvpn.SetLogger(&openvpn.SimpleLogger{})

	client, err := openvpn.NewVPNClient()
	if err != nil {
		log.Fatalf("Failed to initialize VPN client: %v", err)
	}

	client.SetConfig([]byte(config))
	client.SetCredentials(username, password)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()

	fmt.Println("Connecting to VPN...")
	if err := client.Connect(ctx); err != nil {
		if errors.Is(err, openvpn.ErrConnectionFailed) {
			log.Printf("VPN connection failed: %v", err)
		} else if errors.Is(err, openvpn.ErrTimeout) {
			log.Printf("VPN connection timed out.")
		} else {
			log.Fatalf("Unexpected VPN error: %v", err)
		}
	}

	fmt.Printf("Connected! Status: %s\n", client.Status())
	fmt.Println("Sleeping for 5 seconds...")
	time.Sleep(5 * time.Second)

	fmt.Println("Disconnecting...")
	waitCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client.DisconnectAndWait(waitCtx)
	fmt.Printf("Final Status: %s\n", client.Status())
}
