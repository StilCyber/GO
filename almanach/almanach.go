// МАПЫ

// создание мапы, с использованием структуры для значений
type characteristicsDogs struct {
	fear      uint8
	agression uint8
}

var stock = map[string]characteristicsDogs{
	"Шпиц":      {90, 152},
	"Чихуа-хуа": {255, 255},
	"Мопс":      {115, 82},
	"Пудель":    {199, 41},
}