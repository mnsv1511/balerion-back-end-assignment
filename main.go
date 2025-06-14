package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	ThaiDigitWords = map[string]string{
		"0": "ศูนย์",
		"1": "หนึ่ง",
		"2": "สอง",
		"3": "สาม",
		"4": "สี่",
		"5": "ห้า",
		"6": "หก",
		"7": "เจ็ด",
		"8": "แปด",
		"9": "เก้า",
	}
	ThaiPlaceValueWords = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)
		thaiResp := ""

		// Split data Baht and Satang
		inputList := strings.Split(input.String(), ".")

		for i, data := range inputList {
			// Data digit baht
			if i == 0 {
				digitBahtList := strings.Split(data, "")
				// numberOfMillion := (len(digitBahtList) - 1 - i) / 6
				for placeBaht, digitBaht := range digitBahtList {
					digitThai := ThaiDigitWords[digitBaht]
					place := (len(digitBahtList) - 1 - placeBaht) % 6
					placeThai := ThaiPlaceValueWords[place]

					// Condition special thai digit
					if data != "0" && digitBaht == "0" {
						digitThai = ""
						placeThai = ""
					} else if digitBaht == "1" && place == 0 && placeBaht != 0 {
						digitThai = "เอ็ด"
					} else if digitBaht == "1" && place == 1 {
						digitThai = ""
					} else if digitBaht == "2" && place == 1 {
						digitThai = "ยี่"
					}
					if len(digitBahtList)-placeBaht >= 6 && place == 0 {
						placeThai = placeThai + "ล้าน"
					}

					// Concat data digit and place string
					thaiResp = thaiResp + (digitThai + placeThai)
				}

				// Condition baht
				if data == "0" && len(inputList) == 2 {
					thaiResp = ""
					continue // 0 baht xx satang
				} else if len(inputList) == 1 {
					thaiResp = thaiResp + ("บาทถ้วน") // xx baht not have satang

				} else if len(inputList) == 2 {
					thaiResp = thaiResp + ("บาท") // xx baht have satang
				}

			}

			// Data digit satang should be only 01 - 99
			if i == 1 {
				if len(data) == 1 {
					data = data + "0"
				}
				digitSatangList := strings.Split(data, "")
				for placeSatang, digitSatang := range digitSatangList {
					digitThai := ThaiDigitWords[digitSatang]
					place := (len(digitSatangList) - 1 - placeSatang) % 6
					placeThai := ThaiPlaceValueWords[place]

					// Condition special thai digit
					if data != "0" && digitSatang == "0" {
						digitThai = ""
						placeThai = ""
					} else if digitSatang == "1" && place == 0 && placeSatang != 0 && thaiResp != "" {
						digitThai = "เอ็ด"
					} else if digitSatang == "1" && place == 1 {
						digitThai = ""
					} else if digitSatang == "2" && place == 1 {
						digitThai = "ยี่"
					}

					// Concat data digit and place string
					thaiResp = thaiResp + (digitThai + placeThai)
				}
				thaiResp = thaiResp + "สตางค์"
			}
		}

		fmt.Println(thaiResp)
	}
}
