package iso639

import "strings"

type Language struct {
	Family     string `json:"family,omitempty"`
	Name       string `json:"name"`
	NativeName string `json:"native_name,omitempty"`
	Code1      string `json:"code1"`            // ISO639-1 code
	Code2      string `json:"code2"`            // ISO639-2 code
	Code2B     string `json:"code2b,omitempty"` // ISO639-2/B code
	Code3      string `json:"code3,omitempty"`  // ISO639-3 code
	Code6      string `json:"code6,omitempty"`  // ISO639-6 code
}

func ParseISO639_1(code string) (*Language, error) {
	code = strings.Trim(code, " ")
	if r, ok := idxISO639_1[code]; ok {
		return r, nil
	}
	return nil, ErrNotFound
}

func ParseISO639_2(code string) (*Language, error) {
	code = strings.Trim(code, " ")
	if r, ok := idxISO639_2[code]; ok {
		return r, nil
	}
	return nil, ErrNotFound
}

func ParseLanguage(lang string) (*Language, error) {
	name := strings.Trim(lang, " ")
	if len(name) == 2 {
		r, err := ParseISO639_1(name)
		if err == nil {
			return r, nil
		}
	}
	if len(name) == 3 {
		r, err := ParseISO639_2(name)
		if err == nil {
			return r, nil
		}
	}

	name = strings.ToLower(name)
	for len(name) > 3 {
		if r, ok := idxByName[name]; ok {
			return r, nil
		}
		name = name[:len(name)-1]
	}
	return nil, ErrNotFound
}

func MustParseLanguage(lang string) *Language {
	r, err := ParseLanguage(lang)
	if err != nil {
		panic(err)
	}
	return r
}
