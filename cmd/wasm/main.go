package main

import (
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"strconv"
	"strings"
	"syscall/js"

	"bytes"
	"encoding/base64"
	"encoding/hex"

	"github.com/skip2/go-qrcode"
)

func generateQRCode(data string, size int, backgroundColor color.Color, foregroundColor color.Color, recoveryLevel qrcode.RecoveryLevel, disableBorder bool) (string, error) {

	// Create a new QR code with the specified data
	qr, err := qrcode.New(data, recoveryLevel)
	if err != nil {
		return "", err
	}

	qr.DisableBorder = disableBorder

	// Set custom foreground and background colors
	qr.BackgroundColor = backgroundColor
	qr.ForegroundColor = foregroundColor

	// Generate QR code image
	qrImage := qr.Image(size)

	// Encode the image to PNG format
	var buf bytes.Buffer
	err = png.Encode(&buf, qrImage)
	if err != nil {
		return "", err
	}

	// Encode the buffer as a base64 string
	base64QRCode := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64QRCode, nil
}

func qrCodeB64Wrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 6 {
			err := "Invalid no of arguments passed"
			fmt.Printf("validation error %s\n", err)
			return err
		}
		inputText := args[0].String()

		sizeText := args[1].String()
		size, err := strconv.Atoi(sizeText)
		if err != nil {
			fmt.Printf("unable to convert to size %s\n", err)
			return err.Error()
		}

		bgColorText := args[2].String()
		bgColor, err := parseHexColor(bgColorText)
		if err != nil {
			fmt.Printf("unable to convert color %s\n", err)
		}
		fgColorText := args[3].String()
		fgColor, err := parseHexColor(fgColorText)
		if err != nil {
			fmt.Printf("unable to convert color %s\n", err)
		}

		recoveryLevelText := args[4].String()
		recoveryLevel, err := parseRecoveryLevel(recoveryLevelText)
		if err != nil {
			fmt.Printf("unable to convert recovery level %s\n", err)
		}

		disableBorder := args[5].Bool()

		// fmt.Printf("size %d bg %s fg %s level %s border %t \n", size, bgColorText, fgColorText, recoveryLevelText, disableBorder)

		b64, err := generateQRCode(inputText, size, bgColor, fgColor, recoveryLevel, disableBorder)
		if err != nil {
			fmt.Printf("unable to convert to qrcode %s\n", err)
			return err.Error()
		}
		return b64
	})
	return jsonFunc
}

var ErrInvalidHexString = errors.New("invalid hexadecimal string")

func parseHexColor(hexString string) (color.Color, error) {
	// Remove '#' prefix if present
	if len(hexString) > 0 && hexString[0] == '#' {
		hexString = hexString[1:]
	}

	// Parse hexadecimal string to RGB values
	rgb, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	// If the length of the decoded string is not 3 or 6, return an error
	if len(rgb) != 3 && len(rgb) != 6 {
		return nil, ErrInvalidHexString
	}

	// Extend short hex strings (e.g., "abc" => "aabbcc")
	if len(rgb) == 3 {
		rgb = []byte{rgb[0], rgb[0], rgb[1], rgb[1], rgb[2], rgb[2]}
	}

	// Create and return color.RGBA instance
	return color.RGBA{R: rgb[0], G: rgb[1], B: rgb[2], A: 255}, nil
}

func parseRecoveryLevel(levelStr string) (qrcode.RecoveryLevel, error) {
	switch strings.ToLower(levelStr) {
	case "low":
		return qrcode.Low, nil
	case "medium":
		return qrcode.Medium, nil
	case "high":
		return qrcode.High, nil
	case "highest":
		return qrcode.Highest, nil
	default:
		return qrcode.Low, errors.New("invalid recovery level")
	}
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("qrCode64", qrCodeB64Wrapper())
	<-make(chan struct{})
}
