package qr

import (
	"github.com/ahargunyllib/thera-be/pkg/log"
	qrcode "github.com/yougg/go-qrcode"
)

type CustomQRCodeInterface interface {
	Encode(content string, size, margin int) ([]byte, error)
}

type CustomQRCodeStruct struct {
}

var QR = getQRCode()

func getQRCode() CustomQRCodeInterface {
	return &CustomQRCodeStruct{}
}

func (q *CustomQRCodeStruct) Encode(content string, size, margin int) ([]byte, error) {
	qrCode, err := qrcode.Encode(content, qrcode.Medium, size, size, margin)
	if err != nil {
		log.Panic(log.CustomLogInfo{
			"error": err.Error(),
		}, "[QRCODE][Encode] failed to generate qr")
		return nil, err
	}

	return qrCode, nil
}
