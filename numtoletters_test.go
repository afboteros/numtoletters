package numtoletters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesForTestsES = []struct {
	in  int
	out string
}{
	{
		in:  1,
		out: "UNO",
	},
	{
		in:  601,
		out: "SEISCIENTOS UNO",
	},
	{
		in: 125,
		out: "CIENTO VEINTICINCO",
	},
	{
		in: 0,
		out: "CERO",
	},
	{
		in: 00,
		out: "CERO",
	},
	{
		in: 10,
		out: "DIEZ",
	},
	{
		in: -1,
		out: "MENOS UNO",
	},
	{
		in: -1256897,
		out: "MENOS UN MILLÓN DOSCIENTOS CINCUENTA Y SEIS MIL OCHOCIENTOS NOVENTA Y SIETE",
	},
}

var casesForTestsEN = []struct {
	in  int
	out string
}{
	{
		in:  1,
		out: "ONE",
	},
	{
		in:  601,
		out: "SIX HUNDRED AND ONE",
	},
	{
		in: 125,
		out: "ONE HUNDRED AND TWENTY FIVE",
	},
	{
		in: 0,
		out: "ZERO",
	},
	{
		in: 00,
		out: "ZERO",
	},
	{
		in: 10,
		out: "TEN",
	},
	{
		in: -1,
		out: "MINUS ONE",
	},
	{
		in: -1256897,
		out: "MINUS ONE MILLION TWO HUNDRED AND FIFTY SIX THOUSAND EIGHT HUNDRED AND NINETY SEVEN",
	},
}

func TestNumberToWordES(t *testing.T) {
	for i, test := range casesForTestsES {
		actual := NumberToWord(LocaleES, test.in)
		assert.Equal(t, test.out, actual, "Prueba #%d ", i)
	}
}

func TestNumberToWordEN(t *testing.T) {
	for i, test := range casesForTestsEN {
		actual := NumberToWord(LocaleEN, test.in)
		assert.Equal(t, test.out, actual, "Test #%d ", i)
	}
}
