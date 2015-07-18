package iso3166

import "strings"

type Region struct {
	Name   string            `json:"name"`
	Alpha2 string            `json:"code1"` // ISO3166-1 alpha-2
	Alpha3 string            `json:"code2"` // ISO3166-1 alpha-3
	Code3  string            `json:"code3"` // ISO3166-1 numeric
	Alias  string            `json:"alias,omitempty"`
	Names  map[string]string `json:"names,omitempty"`
}

func ParseISO3166_1(s string) (*Region, error) {
	s = strings.Trim(s, " ")
	s = strings.ToUpper(s)
	if r, ok := idxISO3166_1[s]; ok {
		return r, nil
	}
	return nil, ErrNotFound
}

func ParseISO3166_2(s string) (*Region, error) {
	s = strings.Trim(s, " ")
	s = strings.ToUpper(s)
	if r, ok := idxISO3166_2[s]; ok {
		return r, nil
	}
	return nil, ErrNotFound
}

func ParseRegion(name string) (*Region, error) {
	name = strings.Trim(name, " ")
	if len(name) == 2 {
		r, err := ParseISO3166_1(name)
		if err == nil {
			return r, nil
		}
	}
	if len(name) == 3 {
		r, err := ParseISO3166_2(name)
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

func MustParseRegion(name string) *Region {
	r, err := ParseRegion(name)
	if err != nil {
		panic(err)
	}
	return r
}
