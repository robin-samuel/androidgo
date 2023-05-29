package locale

type Locale struct {
	Language string
	Country  string
}

func RandomLocale() Locale {
	return Locale{
		Language: "en",
		Country:  "US",
	}
}

func (l Locale) GetLanguage() string {
	return l.Language
}

func (l Locale) GetCountry() string {
	return l.Country
}

var DE = Locale{"de", "DE"}
var EN = Locale{"en", "US"}
var ES = Locale{"es", "ES"}
var FR = Locale{"fr", "FR"}
var IT = Locale{"it", "IT"}
var JA = Locale{"ja", "JP"}
var KO = Locale{"ko", "KR"}
var PT = Locale{"pt", "BR"}
var RU = Locale{"ru", "RU"}
var ZH = Locale{"zh", "CN"}
