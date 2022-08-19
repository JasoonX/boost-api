package enums

// full list for future reference:
//	simple
//	arabic
//	armenian
//	basque
//	catalan
//	danish
//	dutch
//	english
//	finnish
//	french
//	german
//	greek
//	hindi
//	hungarian
//	indonesian
//	irish
//	italian
//	lithuanian
//	nepali
//	norwegian
//	portuguese
//	romanian
//	russian
//	serbian
//	spanish
//	swedish
//	tamil
//	turkish
//	yiddish

var psqlLocaleMappingISO2 = map[string]string{
	defaultLocale: "simple",
	EN:            "english",
	DE:            "german",
	// for the MVP will treat all slovenian languages (with cyrillic) the same way
	RU: "russian",
	UA: "russian",
}

func LocaleFromISO2(iso2 string) string {
	if locale, ok := psqlLocaleMappingISO2[iso2]; ok {
		return locale
	}
	return psqlLocaleMappingISO2[defaultLocale]
}
