package importador

import (
	"time"

	"github.com/tealeg/xlsx/v3"
	"gopkg.in/guregu/null.v3"
)

func leerComoBool(data *xlsx.Cell) bool {

	dataString := leerComoString(data)

	return dataString == verdadero
}

func leerComoNullBool(data *xlsx.Cell) null.Bool {
	dataString := leerComoString(data)

	if len(dataString) == 0 {
		return null.NewBool(false, false)
	}

	return null.NewBool(leerComoBool(data), true)
}

func leerComoInt(data *xlsx.Cell) (int, error) {
	dataInt, err := data.Int()
	if err != nil {
		return dataInt, err
	}

	return dataInt, nil
}

func leerComoString(data *xlsx.Cell) string {
	dataString := data.String()
	return dataString
}

func leerComoNullString(data *xlsx.Cell) null.String {
	dataString := leerComoString(data)

	if len(dataString) == 0 {
		return null.NewString(dataString, false)
	}

	return null.NewString(dataString, true)
}

func leerComoDatetime(data *xlsx.Cell) (time.Time, error) {
	time, err := data.GetTime(false)
	if err != nil {
		return time, err
	}

	return time, nil
}


