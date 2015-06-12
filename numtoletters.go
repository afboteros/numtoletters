package numtoletters

import (
	"strconv"

	"github.com/magiconair/properties"
)

// Locales
const (
	LocaleES = "es."
	LocaleEN = "en."
)

// NumberToWord converts a number to words representing the name
func NumberToWord(locale string, number int) string {
	// Result variable to append the resulting string on each step
	var result string
	if number < 0 {
		result = result + getProperty(locale, "minus") + " " + NumberToWord(locale, number * -1)
	} else if len(strconv.Itoa(number)) >= 7 && len(strconv.Itoa(number)) <= 9 { // Million case
		million := number / 1000000
		resto := number - (million * 1000000)
		switch million {
		case 1:
			result = result + getProperty(locale, "million") + " " + NumberToWord(locale, resto)
		default:
			result = result + NumberToWord(locale, million) + " " + getProperty(locale, "millions") + " " + NumberToWord(locale, resto)
		}
	} else if len(strconv.Itoa(number)) >= 4 && len(strconv.Itoa(number)) <= 6 { // Thousand Case
		thousand := number / 1000
		resto := number - (thousand * 1000)
		if resto == 0 {
			result = result + " "
		}
		switch thousand {
		case 1:
			result = result + getProperty(LocaleES, "thousand") + " " + NumberToWord(locale, resto)
		default:
			result = result + NumberToWord(locale, thousand) + " " + getProperty(locale, "thousand") + " " + NumberToWord(locale, resto)
		}
	} else if len(strconv.Itoa(number)) == 3 { // Hundred Case
		hundred := number / 100
		resto := number - (hundred * 100)
		if resto == 0 {
			result = result + getProperty(locale, "hundred")
		} else {
			switch hundred {
			case 1:
				if locale == LocaleES {
					result = result + getProperty(locale, "1hundred") + " " + NumberToWord(locale, resto)
				} else if locale == LocaleEN {
					result = result + getProperty(locale, "1hundred") + " AND " + NumberToWord(locale, resto)
				}
			case 5:
				if locale == LocaleES {
					result = result + getProperty(locale, "5hundred") + " " + NumberToWord(locale, resto)
				} else if locale == LocaleEN {
					result = result + getProperty(locale, "5hundred") + " AND " + NumberToWord(locale, resto)
				}
			case 7:
				if locale == LocaleES {
					result = result + getProperty(locale, "7hundred") + " " + NumberToWord(locale, resto)
				} else if locale == LocaleEN {
					result = result + getProperty(locale, "7hundred") + " AND " + NumberToWord(locale, resto)
				}
			case 9:
				if locale == LocaleES {
					result = result + getProperty(locale, "9hundred") + " " + NumberToWord(locale, resto)
				} else if locale == LocaleEN {
					result = result + getProperty(locale, "9hundred") + " AND " + NumberToWord(locale, resto)
				}
			default:
				if locale == LocaleES {
					result = result + NumberToWord(locale, hundred) + getProperty(locale, "hundreds") + " " + NumberToWord(locale, resto)
				} else if locale == LocaleEN {
					result = result + NumberToWord(locale, hundred) + " " + getProperty(locale, "hundreds") + " " + NumberToWord(locale, resto)
				}
			}
		}
	} else if len(strconv.Itoa(number)) == 2 { // Ten Case
		ten := number / 10
		resto := number - (ten * 10)
		if resto == 0 {
			switch ten {
			case 1:
				result = result + getProperty(locale, "10")
			case 2:
				result = result + getProperty(locale, "20")
			case 3:
				result = result + getProperty(locale, "30")
			case 4:
				result = result + getProperty(locale, "40")
			case 5:
				result = result + getProperty(locale, "50")
			case 6:
				result = result + getProperty(locale, "60")
			case 7:
				result = result + getProperty(locale, "70")
			case 8:
				result = result + getProperty(locale, "80")
			case 9:
				result = result + getProperty(locale, "90")
			}
		} else {
			switch ten {
			case 1:
				switch resto {
				case 1:
					result = result + getProperty(locale, "11")
				case 2:
					result = result + getProperty(locale, "12")
				case 3:
					result = result + getProperty(locale, "13")
				case 4:
					result = result + getProperty(locale, "14")
				case 5:
					result = result + getProperty(locale, "15")
				default:
					result = result + getProperty(locale, "1X") + " " + NumberToWord(locale, resto)
				}
			case 2:
				if locale == LocaleES {
					result = result + getProperty(locale, "2X") + NumberToWord(locale, resto)
				} else if locale == LocaleEN {
					result = result + getProperty(locale, "2X") + " " + NumberToWord(locale, resto)
				}
			case 3:
				result = result + getProperty(locale, "3X") + " " + NumberToWord(locale, resto)
			case 4:
				result = result + getProperty(locale, "4X") + " " + NumberToWord(locale, resto)
			case 5:
				result = result + getProperty(locale, "5X") + " " + NumberToWord(locale, resto)
			case 6:
				result = result + getProperty(locale, "6X") + " " + NumberToWord(locale, resto)
			case 7:
				result = result + getProperty(locale, "7X") + " " + NumberToWord(locale, resto)
			case 8:
				result = result + getProperty(locale, "8X") + " " + NumberToWord(locale, resto)
			case 9:
				result = result + getProperty(locale, "9X") + " " + NumberToWord(locale, resto)
			}
		}
	} else if len(strconv.Itoa(number)) == 1 { // Unit Case
		switch number {
		case 0:
			result = getProperty(locale, "0")
		case 1:
			result = result + getProperty(locale, "1")
		case 2:
			result = result + getProperty(locale, "2")
		case 3:
			result = result + getProperty(locale, "3")
		case 4:
			result = result + getProperty(locale, "4")
		case 5:
			result = result + getProperty(locale, "5")
		case 6:
			result = result + getProperty(locale, "6")
		case 7:
			result = result + getProperty(locale, "7")
		case 8:
			result = result + getProperty(locale, "8")
		case 9:
			result = result + getProperty(locale, "9")
		}
	}
	return result
}

func getProperty(locale string, name string) string {
	var p *properties.Properties
	p = properties.MustLoadFile("./locales/locale.properties", properties.UTF8)
	property := p.MustGet(locale + name)
	return property
}
