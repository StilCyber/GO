// ПЕРЕМЕННЫЕ

var wife string
wife = "Кисуля"

wife := "Рыба моя"

var name, age = "Некая тян", 22

var (
	husband = true
	lovers = true
	children = 3
)

const law = "Семейный кодекс"

const mainLaw, codex = "Конституция", "Семейный кодекс"

const (
	mainLaw = "Конституция"
	codex = "Семейный кодекс"
)

Компилятор go требует, чтобы мы обязательно использовали объявленную переменную
var lovers = "Генерал Кастер"
_ = lovers

 Благодаря нижнему подчеркиваю, оно называется blank identifier, мы избежим ошибок при компиляции 
 Это особенно полезно в следующих случаях:
 1. Функция возвращает данные, и ошибку, если не хотим обрабатывать ошибку, заменяем ее _
 2. Импортируем пакет, который пока не хотим использовать, но импорт нужен обязательно или потребуется в будущем

// АРИФМЕТИКА

базовые операции - +, -, *, /
возведение в квадрат - ** 
остаток от деления - %

Одна из самых частых операций - увеличение на единицу 
Для этого есть краткий синтаксис

var i = 0
Вместо того, чтобы писать 
i = i + 1
можно записать так:
i++

Еще один вариант:
i += 1

Кстати, такой синтасиси можно использовать и в других случаях:
var i = 0

i += 1 // 1
i *= 10 // 10
i -= 2 // 8

Для вычитания на единицу синтаксис схож:
i--

Прибавление единицы называется ИНКРЕМЕНТ
Вычитание единицы - ДЕКРЕМЕНТ

// УСЛОВИЯ И ЛОГИКА

базовое сравнение - >, <, >=, <=
равенство - == 
неравенство = !=

логические операции - отрицание, конъюнкция и дизъюнкция

Отрицание
fmt.Println(!true) // false

Конъюнкция (логическое умножение)
Если оба операнда истинны - только тогда будет true
fmt.Println(true && true) // true
fmt.Println(true && false) // false

Дизъюнкция (логическое сложение)
Хотя бы один операнд не должен быть false
fmt.Println(true || true) // true
fmt.Println(true || false) // true
fmt.Println(false || false) // false

Ложные значения в go:
false

В отличие от других языков, в go ложным является только false 
   нули, пустые строки в этом празднике жизни не участвуют

// УСЛОВНОЕ ВЕТВЛЕНИЕ и SWITCH

if условие {
	некий сомнительный код
 } else if другое условие {
	еще более подозрительно
 } else {
	если везде отшили
 }

 switch условие:
	case плохой вариант 1:
		код...
	case плохой вариант 2:
		код...
	default:
		обычная невезуха

// ЦИКЛЫ

// базовый вариант for 
children := 0

for i := 0; i < 3; i++ {
	fmt.Printf("Сейчас детей: %v\n", children)
	fmt.Println("Нужно родить еще")
	children++
}

// for как while
var lovers uint

fmt.Println("Введите количество любовников")
fmt.Scan(&lovers)

for lovers > 0 {
	fmt.Printf("Текущее число любовников: %v\n", lovers)
	lovers--
}

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