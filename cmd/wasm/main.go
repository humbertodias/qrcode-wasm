package main

import (
	"fmt"
	"syscall/js"

	"bytes"
	"encoding/base64"

	qrcode "github.com/skip2/go-qrcode"
)

func generateQRCode(input string) (string, error) {
	// Generate QR code for the input string
	qr, err := qrcode.New(input, qrcode.Medium)
	if err != nil {
		return "", err
	}

	// Encode the QR code as PNG to a buffer
	var buf bytes.Buffer
	if err := qr.Write(256, &buf); err != nil {
		return "", err
	}

	// Convert the buffer to a base64 encoded string
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Return the base64 encoded string
	return base64String, err
}

func qrCodeB64Wrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputText := args[0].String()
		b64, err := generateQRCode(inputText)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return b64
	})
	return jsonFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("qrCode64", qrCodeB64Wrapper())
	<-make(chan struct{})
}
