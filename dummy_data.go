package main

import "learn_words/common"

var data = common.Words{
	//{
	//	Original:    "Cesta",
	//	Translation: "Путь",
	//}, {
	//	Original:    "Zeď",
	//	Translation: "Стена",
	//}, {
	//	Original:    "Zítra",
	//	Translation: "Завтра",
	//}, {
	//	Original:    "Střecha",
	//	Translation: "Крыша",
	//}, {
	//	Original:    "Cestovat",
	//	Translation: "Путешествовать",
	//}, {
	//	Original:    "Co znamená..?",
	//	Translation: "Что значит ...",
	//}, {
	//	Original:    "Jak se řekne",
	//	Translation: "Как это произносится",
	//}, {
	//	Original:    "Náhoda",
	//	Translation: "Совпадение",
	//}, {
	//	Original:    "Určitě",
	//	Translation: "Конечно, однозначно, определенно",
	//}, {
	//	Original:    "Nizozemsko",
	//	Translation: "Нидерланды",
	//}, {
	//	Original:    "Žák",
	//	Translation: "Школьник/ученик",
	//}, {
	//	Original:    "Počasí",
	//	Translation: "Погода",
	//}, {
	//	Original:    "Zajímavý",
	//	Translation: "Интересный",
	//}, {
	//	Original:    "Budova",
	//	Translation: "Здание",
	//}, {
	//	Original:    "Ošklivý",
	//	Translation: "Безобразный",
	//}, {
	//	Original:    "Obyčejný",
	//	Translation: "Обычный",
	//}, {
	//	Original:    "Nudný",
	//	Translation: "Скучный",
	//}, {
	//	Original:    "Pohodlný",
	//	Translation: "Удобный",
	//}, {
	//	Original:    "Obývák",
	//	Translation: "Зал/гостинная",
	//}, {
	//	Original:    "Předsíň",
	//	Translation: "Прихожая",
	//}, {
	//	Original:    "Levní",
	//	Translation: "Дешевый",
	//}, {
	//	Original:    "Mrakodrap",
	//	Translation: "Небоскреб",
	//}, {
	//	Original:    "Sklep",
	//	Translation: "Подвал",
	//}, {
	//	Original:    "Zahrada",
	//	Translation: "Придомовая территория",
	//}, {
	//	Original:    "Záchod",
	//	Translation: "Туалет",
	//}, {
	//	Original:    "Gauč",
	//	Translation: "Диван",
	//}, {
	//	Original:    "Koberec",
	//	Translation: "Ковер",
	//}, {
	//	Original:    "Skříň",
	//	Translation: "Шкаф",
	//}, {
	//	Original:    "Zrcadlo",
	//	Translation: "Зеркало",
	//}, {
	//	Original:    "Šedý",
	//	Translation: "Серый/седой",
	//}, {
	//	Original:    "Nic",
	//	Translation: "Ничего",
	//}, {
	//	Original:    "Dovolená",
	//	Translation: "Отпуск",
	//}, {
	//	Original:    "Uprostřed",
	//	Translation: "Посередине",
	//}, {
	//	Original:    "Kde se ti libilo?",
	//	Translation: "Где тебе понравилось",
	//}, {
	//	Original:    "Jet na výlet",
	//	Translation: "Отправиться в путешествие",
	//}, {
	//	Original:    "Nějaky",
	//	Translation: "Какой-нибудь",
	//}, {
	//	Original:    "Zpǎtky",
	//	Translation: "Назад",
	//}, {
	//	Original:    "Musíte",
	//	Translation: "Must",
	//}, {
	//	Original:    "Křižovatka",
	//	Translation: "Перекресток",
	//}, {
	//	Original:    "Byla návštěva",
	//	Translation: "Были гостях",
	//}, {
	//	Original:    "Navštívit",
	//	Translation: "Посещать, ходить в гости",
	//}, {
	//	Original:    "Dávat dárky",
	//	Translation: "Дарить подарки",
	//}, {
	//	Original:    "Dostat",
	//	Translation: "Получить",
	//}, {
	//	Original:    "Svetr",
	//	Translation: "Свитер",
	//}, {
	//	Original:    "Směr",
	//	Translation: "Наприавление",
	//}, {
	//	Original:    "Přítel",
	//	Translation: "Друг",
	//}, {
	//	Original:    "Herec",
	//	Translation: "Актер",
	//}, {
	//	Original:    "Úředník",
	//	Translation: "Чиновник/клерк",
	//}, {
	//	Original:    "Strom",
	//	Translation: "Дерево",
	//}, {
	//	Original:    "Snídat",
	//	Translation: "Завртакать",
	//}, {
	//	Original:    "Snídaně",
	//	Translation: "Завртак",
	//}, {
	//	Original:    "Večeřet",
	//	Translation: "Ужинать",
	//}, {
	//	Original:    "Večeře",
	//	Translation: "Ужин",
	//}, {
	//	Original:    "Hledat",
	//	Translation: "Искать",
	//}, {
	//	Original:    "Spěchat",
	//	Translation: "Спешить",
	//}, {
	//	Original:    "Zřídka",
	//	Translation: "Редко",
	//}, {
	//	Original:    "Občas",
	//	Translation: "Иногда",
	//}, {
	//	Original:    "Hodně",
	//	Translation: "Много",
	//}, {
	//	Original:    "Zavináč",
	//	Translation: "@",
	//}, {
	//	Original:    "Žízeň",
	//	Translation: "Жажда",
	//}, {
	//	Original:    "Hlad",
	//	Translation: "Голод",
	//}, {
	//	Original:    "Chlad",
	//	Translation: "Холод",
	//}, {
	//	Original:    "Zkouška",
	//	Translation: "Экзамен",
	//}, {
	//	Original:    "Opakovat",
	//	Translation: "Повторять",
	//}, {
	//	Original:    "Zvyknout",
	//	Translation: "Привыкать",
	//}, {
	//	Original:    "Sejdeme se",
	//	Translation: "Встретимся/пересечемся",
	//}, {
	//	Original:    "Chutnat",
	//	Translation: "Нравиться по вкусу",
	//},
	{
		Original:    "Vyhovovat",
		Translation: "Подходить, соответсвовать",
	}, {
		Original:    "Zůstat",
		Translation: "Остаться",
	}, {
		Original:    "cvičit",
		Translation: "Делать упражнения, тренироваться",
	}, {
		Original:    "Sportovec",
		Translation: "Спортсмен",
	}, {
		Original:    "kluk",
		Translation: "Мальчик",
	}, {
		Original:    "lidi",
		Translation: "Люди",
	}, {
		Original:    "kuchař",
		Translation: "Повар",
	}, {
		Original:    "dlouho",
		Translation: "Долго",
	}, {
		Original:    "procházet se",
		Translation: "Ходить, прогуливаться",
	}, {
		Original:    "vůbec",
		Translation: "Вообще",
	}, {
		Original:    "teď",
		Translation: "Сейчас",
	},
}
