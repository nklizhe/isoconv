package iso3166

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrNotFound = errors.New("not found")

	allRegions   []Region
	iso3166Data  = `[{"name":"Afghanistan","code1":"AF","code2":"AFG"},{"name":"","code1":"AX","code2":"ALA"},{"name":"Albania","code1":"AL","code2":"ALB"},{"name":"Algeria","code1":"DZ","code2":"DZA"},{"name":"American Samoa","code1":"AS","code2":"ASM"},{"name":"Andorra","code1":"AD","code2":"AND"},{"name":"Angola","code1":"AO","code2":"AGO"},{"name":"Anguilla","code1":"AI","code2":"AIA"},{"name":"Antarctica","code1":"AQ","code2":"ATA"},{"name":"Antigua and Barbuda","code1":"AG","code2":"ATG"},{"name":"Argentina","code1":"AR","code2":"ARG"},{"name":"Armenia","code1":"AM","code2":"ARM"},{"name":"Aruba","code1":"AW","code2":"ABW"},{"name":"Australia","code1":"AU","code2":"AUS"},{"name":"Austria","code1":"AT","code2":"AUT"},{"name":"Azerbaijan","code1":"AZ","code2":"AZE"},{"name":"Bahamas","code1":"BS","code2":"BHS"},{"name":"Bahrain","code1":"BH","code2":"BHR"},{"name":"Bangladesh","code1":"BD","code2":"BGD"},{"name":"Barbados","code1":"BB","code2":"BRB"},{"name":"Belarus","code1":"BY","code2":"BLR"},{"name":"Belgium","code1":"BE","code2":"BEL"},{"name":"Belize","code1":"BZ","code2":"BLZ"},{"name":"Benin","code1":"BJ","code2":"BEN"},{"name":"Bermuda","code1":"BM","code2":"BMU"},{"name":"Bhutan","code1":"BT","code2":"BTN"},{"name":"Bolivia (Plurinational State of)","code1":"BO","code2":"BOL"},{"name":"Bonaire, Sint Eustatius and Saba","code1":"BQ","code2":"BES"},{"name":"Bosnia and Herzegovina","code1":"BA","code2":"BIH"},{"name":"Botswana","code1":"BW","code2":"BWA"},{"name":"Bouvet Island","code1":"BV","code2":"BVT"},{"name":"Brazil","code1":"BR","code2":"BRA"},{"name":"British Indian Ocean Territory","code1":"IO","code2":"IOT"},{"name":"Brunei Darussalam","code1":"BN","code2":"BRN"},{"name":"Bulgaria","code1":"BG","code2":"BGR"},{"name":"Burkina Faso","code1":"BF","code2":"BFA"},{"name":"Burundi","code1":"BI","code2":"BDI"},{"name":"Cambodia","code1":"KH","code2":"KHM"},{"name":"Cameroon","code1":"CM","code2":"CMR"},{"name":"Canada","code1":"CA","code2":"CAN"},{"name":"Cabo Verde","code1":"CV","code2":"CPV"},{"name":"Cayman Islands","code1":"KY","code2":"CYM"},{"name":"Central African Republic","code1":"CF","code2":"CAF"},{"name":"Chad","code1":"TD","code2":"TCD"},{"name":"Chile","code1":"CL","code2":"CHL"},{"name":"China","code1":"CN","code2":"CHN"},{"name":"Christmas Island","code1":"CX","code2":"CXR"},{"name":"Cocos (Keeling) Islands","code1":"CC","code2":"CCK"},{"name":"Colombia","code1":"CO","code2":"COL"},{"name":"Comoros","code1":"KM","code2":"COM"},{"name":"Congo","code1":"CG","code2":"COG"},{"name":"Congo (Democratic Republic of the)","code1":"CD","code2":"COD"},{"name":"Cook Islands","code1":"CK","code2":"COK"},{"name":"Costa Rica","code1":"CR","code2":"CRI"},{"name":"","code1":"CI","code2":"CIV"},{"name":"Croatia","code1":"HR","code2":"HRV"},{"name":"Cuba","code1":"CU","code2":"CUB"},{"name":"","code1":"CW","code2":"CUW"},{"name":"Cyprus","code1":"CY","code2":"CYP"},{"name":"Czech Republic","code1":"CZ","code2":"CZE"},{"name":"Denmark","code1":"DK","code2":"DNK"},{"name":"Djibouti","code1":"DJ","code2":"DJI"},{"name":"Dominica","code1":"DM","code2":"DMA"},{"name":"Dominican Republic","code1":"DO","code2":"DOM"},{"name":"Ecuador","code1":"EC","code2":"ECU"},{"name":"Egypt","code1":"EG","code2":"EGY"},{"name":"El Salvador","code1":"SV","code2":"SLV"},{"name":"Equatorial Guinea","code1":"GQ","code2":"GNQ"},{"name":"Eritrea","code1":"ER","code2":"ERI"},{"name":"Estonia","code1":"EE","code2":"EST"},{"name":"Ethiopia","code1":"ET","code2":"ETH"},{"name":"Falkland Islands (Malvinas)","code1":"FK","code2":"FLK"},{"name":"Faroe Islands","code1":"FO","code2":"FRO"},{"name":"Fiji","code1":"FJ","code2":"FJI"},{"name":"Finland","code1":"FI","code2":"FIN"},{"name":"France","code1":"FR","code2":"FRA"},{"name":"French Guiana","code1":"GF","code2":"GUF"},{"name":"French Polynesia","code1":"PF","code2":"PYF"},{"name":"French Southern Territories","code1":"TF","code2":"ATF"},{"name":"Gabon","code1":"GA","code2":"GAB"},{"name":"Gambia","code1":"GM","code2":"GMB"},{"name":"Georgia","code1":"GE","code2":"GEO"},{"name":"Germany","code1":"DE","code2":"DEU"},{"name":"Ghana","code1":"GH","code2":"GHA"},{"name":"Gibraltar","code1":"GI","code2":"GIB"},{"name":"Greece","code1":"GR","code2":"GRC"},{"name":"Greenland","code1":"GL","code2":"GRL"},{"name":"Grenada","code1":"GD","code2":"GRD"},{"name":"Guadeloupe","code1":"GP","code2":"GLP"},{"name":"Guam","code1":"GU","code2":"GUM"},{"name":"Guatemala","code1":"GT","code2":"GTM"},{"name":"Guernsey","code1":"GG","code2":"GGY"},{"name":"Guinea","code1":"GN","code2":"GIN"},{"name":"Guinea-Bissau","code1":"GW","code2":"GNB"},{"name":"Guyana","code1":"GY","code2":"GUY"},{"name":"Haiti","code1":"HT","code2":"HTI"},{"name":"Heard Island and McDonald Islands","code1":"HM","code2":"HMD"},{"name":"Holy See","code1":"VA","code2":"VAT"},{"name":"Honduras","code1":"HN","code2":"HND"},{"name":"Hong Kong","code1":"HK","code2":"HKG"},{"name":"Hungary","code1":"HU","code2":"HUN"},{"name":"Iceland","code1":"IS","code2":"ISL"},{"name":"India","code1":"IN","code2":"IND"},{"name":"Indonesia","code1":"ID","code2":"IDN"},{"name":"Iran (Islamic Republic of)","code1":"IR","code2":"IRN"},{"name":"Iraq","code1":"IQ","code2":"IRQ"},{"name":"Ireland","code1":"IE","code2":"IRL"},{"name":"Isle of Man","code1":"IM","code2":"IMN"},{"name":"Israel","code1":"IL","code2":"ISR"},{"name":"Italy","code1":"IT","code2":"ITA"},{"name":"Jamaica","code1":"JM","code2":"JAM"},{"name":"Japan","code1":"JP","code2":"JPN"},{"name":"Jersey","code1":"JE","code2":"JEY"},{"name":"Jordan","code1":"JO","code2":"JOR"},{"name":"Kazakhstan","code1":"KZ","code2":"KAZ"},{"name":"Kenya","code1":"KE","code2":"KEN"},{"name":"Kiribati","code1":"KI","code2":"KIR"},{"name":"Korea (Democratic People's Republic of)","code1":"KP","code2":"PRK"},{"name":"Korea (Republic of)","code1":"KR","code2":"KOR"},{"name":"Kuwait","code1":"KW","code2":"KWT"},{"name":"Kyrgyzstan","code1":"KG","code2":"KGZ"},{"name":"Lao People's Democratic Republic","code1":"LA","code2":"LAO"},{"name":"Latvia","code1":"LV","code2":"LVA"},{"name":"Lebanon","code1":"LB","code2":"LBN"},{"name":"Lesotho","code1":"LS","code2":"LSO"},{"name":"Liberia","code1":"LR","code2":"LBR"},{"name":"Libya","code1":"LY","code2":"LBY"},{"name":"Liechtenstein","code1":"LI","code2":"LIE"},{"name":"Lithuania","code1":"LT","code2":"LTU"},{"name":"Luxembourg","code1":"LU","code2":"LUX"},{"name":"Macao","code1":"MO","code2":"MAC"},{"name":"Macedonia (the former Yugoslav Republic of)","code1":"MK","code2":"MKD"},{"name":"Madagascar","code1":"MG","code2":"MDG"},{"name":"Malawi","code1":"MW","code2":"MWI"},{"name":"Malaysia","code1":"MY","code2":"MYS"},{"name":"Maldives","code1":"MV","code2":"MDV"},{"name":"Mali","code1":"ML","code2":"MLI"},{"name":"Malta","code1":"MT","code2":"MLT"},{"name":"Marshall Islands","code1":"MH","code2":"MHL"},{"name":"Martinique","code1":"MQ","code2":"MTQ"},{"name":"Mauritania","code1":"MR","code2":"MRT"},{"name":"Mauritius","code1":"MU","code2":"MUS"},{"name":"Mayotte","code1":"YT","code2":"MYT"},{"name":"Mexico","code1":"MX","code2":"MEX"},{"name":"Micronesia (Federated States of)","code1":"FM","code2":"FSM"},{"name":"Moldova (Republic of)","code1":"MD","code2":"MDA"},{"name":"Monaco","code1":"MC","code2":"MCO"},{"name":"Mongolia","code1":"MN","code2":"MNG"},{"name":"Montenegro","code1":"ME","code2":"MNE"},{"name":"Montserrat","code1":"MS","code2":"MSR"},{"name":"Morocco","code1":"MA","code2":"MAR"},{"name":"Mozambique","code1":"MZ","code2":"MOZ"},{"name":"Myanmar","code1":"MM","code2":"MMR"},{"name":"Namibia","code1":"NA","code2":"NAM"},{"name":"Nauru","code1":"NR","code2":"NRU"},{"name":"Nepal","code1":"NP","code2":"NPL"},{"name":"Netherlands","code1":"NL","code2":"NLD"},{"name":"New Caledonia","code1":"NC","code2":"NCL"},{"name":"New Zealand","code1":"NZ","code2":"NZL"},{"name":"Nicaragua","code1":"NI","code2":"NIC"},{"name":"Niger","code1":"NE","code2":"NER"},{"name":"Nigeria","code1":"NG","code2":"NGA"},{"name":"Niue","code1":"NU","code2":"NIU"},{"name":"Norfolk Island","code1":"NF","code2":"NFK"},{"name":"Northern Mariana Islands","code1":"MP","code2":"MNP"},{"name":"Norway","code1":"NO","code2":"NOR"},{"name":"Oman","code1":"OM","code2":"OMN"},{"name":"Pakistan","code1":"PK","code2":"PAK"},{"name":"Palau","code1":"PW","code2":"PLW"},{"name":"Palestine, State of","code1":"PS","code2":"PSE"},{"name":"Panama","code1":"PA","code2":"PAN"},{"name":"Papua New Guinea","code1":"PG","code2":"PNG"},{"name":"Paraguay","code1":"PY","code2":"PRY"},{"name":"Peru","code1":"PE","code2":"PER"},{"name":"Philippines","code1":"PH","code2":"PHL"},{"name":"Pitcairn","code1":"PN","code2":"PCN"},{"name":"Poland","code1":"PL","code2":"POL"},{"name":"Portugal","code1":"PT","code2":"PRT"},{"name":"Puerto Rico","code1":"PR","code2":"PRI"},{"name":"Qatar","code1":"QA","code2":"QAT"},{"name":"","code1":"RE","code2":"REU"},{"name":"Romania","code1":"RO","code2":"ROU"},{"name":"Russian Federation","code1":"RU","code2":"RUS"},{"name":"Rwanda","code1":"RW","code2":"RWA"},{"name":"","code1":"BL","code2":"BLM"},{"name":"Saint Helena, Ascension and Tristan da Cunha","code1":"SH","code2":"SHN"},{"name":"Saint Kitts and Nevis","code1":"KN","code2":"KNA"},{"name":"Saint Lucia","code1":"LC","code2":"LCA"},{"name":"Saint Martin (French part)","code1":"MF","code2":"MAF"},{"name":"Saint Pierre and Miquelon","code1":"PM","code2":"SPM"},{"name":"Saint Vincent and the Grenadines","code1":"VC","code2":"VCT"},{"name":"Samoa","code1":"WS","code2":"WSM"},{"name":"San Marino","code1":"SM","code2":"SMR"},{"name":"Sao Tome and Principe","code1":"ST","code2":"STP"},{"name":"Saudi Arabia","code1":"SA","code2":"SAU"},{"name":"Senegal","code1":"SN","code2":"SEN"},{"name":"Serbia","code1":"RS","code2":"SRB"},{"name":"Seychelles","code1":"SC","code2":"SYC"},{"name":"Sierra Leone","code1":"SL","code2":"SLE"},{"name":"Singapore","code1":"SG","code2":"SGP"},{"name":"Sint Maarten (Dutch part)","code1":"SX","code2":"SXM"},{"name":"Slovakia","code1":"SK","code2":"SVK"},{"name":"Slovenia","code1":"SI","code2":"SVN"},{"name":"Solomon Islands","code1":"SB","code2":"SLB"},{"name":"Somalia","code1":"SO","code2":"SOM"},{"name":"South Africa","code1":"ZA","code2":"ZAF"},{"name":"South Georgia and the South Sandwich Islands","code1":"GS","code2":"SGS"},{"name":"South Sudan","code1":"SS","code2":"SSD"},{"name":"Spain","code1":"ES","code2":"ESP"},{"name":"Sri Lanka","code1":"LK","code2":"LKA"},{"name":"Sudan","code1":"SD","code2":"SDN"},{"name":"Suriname","code1":"SR","code2":"SUR"},{"name":"Svalbard and Jan Mayen","code1":"SJ","code2":"SJM"},{"name":"Swaziland","code1":"SZ","code2":"SWZ"},{"name":"Sweden","code1":"SE","code2":"SWE"},{"name":"Switzerland","code1":"CH","code2":"CHE"},{"name":"Syrian Arab Republic","code1":"SY","code2":"SYR"},{"name":"Taiwan, Province of China","code1":"TW","code2":"TWN"},{"name":"Tajikistan","code1":"TJ","code2":"TJK"},{"name":"Tanzania, United Republic of","code1":"TZ","code2":"TZA"},{"name":"Thailand","code1":"TH","code2":"THA"},{"name":"Timor-Leste","code1":"TL","code2":"TLS"},{"name":"Togo","code1":"TG","code2":"TGO"},{"name":"Tokelau","code1":"TK","code2":"TKL"},{"name":"Tonga","code1":"TO","code2":"TON"},{"name":"Trinidad and Tobago","code1":"TT","code2":"TTO"},{"name":"Tunisia","code1":"TN","code2":"TUN"},{"name":"Turkey","code1":"TR","code2":"TUR"},{"name":"Turkmenistan","code1":"TM","code2":"TKM"},{"name":"Turks and Caicos Islands","code1":"TC","code2":"TCA"},{"name":"Tuvalu","code1":"TV","code2":"TUV"},{"name":"Uganda","code1":"UG","code2":"UGA"},{"name":"Ukraine","code1":"UA","code2":"UKR"},{"name":"United Arab Emirates","code1":"AE","code2":"ARE"},{"name":"United Kingdom of Great Britain and Northern Ireland","code1":"GB","code2":"GBR"},{"name":"United States of America","code1":"US","code2":"USA"},{"name":"United States Minor Outlying Islands","code1":"UM","code2":"UMI"},{"name":"Uruguay","code1":"UY","code2":"URY"},{"name":"Uzbekistan","code1":"UZ","code2":"UZB"},{"name":"Vanuatu","code1":"VU","code2":"VUT"},{"name":"Venezuela (Bolivarian Republic of)","code1":"VE","code2":"VEN"},{"name":"Viet Nam","code1":"VN","code2":"VNM"},{"name":"Virgin Islands (British)","code1":"VG","code2":"VGB"},{"name":"Virgin Islands (U.S.)","code1":"VI","code2":"VIR"},{"name":"Wallis and Futuna","code1":"WF","code2":"WLF"},{"name":"Western Sahara","code1":"EH","code2":"ESH"},{"name":"Yemen","code1":"YE","code2":"YEM"},{"name":"Zambia","code1":"ZM","code2":"ZMB"},{"name":"Zimbabwe","code1":"ZW","code2":"ZWE"}]`
	idxISO3166_1 = make(map[string]*Region)
	idxISO3166_2 = make(map[string]*Region)
	idxByName    = make(map[string]*Region)
)

func init() {
	err := json.Unmarshal([]byte(iso3166Data), &allRegions)
	if err != nil {
		panic(err)
	}

	// build indices
	for i := range allRegions {
		v := &allRegions[i]
		v.Alpha2 = strings.ToUpper(v.Alpha2)
		v.Alpha3 = strings.ToUpper(v.Alpha3)
		idxISO3166_1[v.Alpha2] = v
		idxISO3166_2[v.Alpha3] = v
		idxByName[strings.ToLower(v.Name)] = v
	}
}