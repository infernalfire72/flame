package utils

import (
	"unsafe"
)

var countryCodes = [253]int16{
	0x5858, // XX Unknown Country
	0x434f, // OC Oceania Continent
	0x5545, // EU Europe
	0x4441, // AD Andorra
	0x4541, // AE United Arabian Emirates
	0x4641, // AF Afghanistan
	0x4741, // AG Antigua
	0x4941, // AI Anguilla
	0x4c41, // AL Albania
	0x4d41, // AM Armenia
	0x4e41, // AN Antilles
	0x4f41, // AO Angola
	0x5141, // AQ Antarctica
	0x5241, // AR Argentina
	0x5341, // AS American Samoa
	0x5441, // AT Austria
	0x5541, // AU Australia
	0x5741, // AW Aruba
	0x5a41, // AZ Azerbaijan
	0x4142, // BA Bosnia
	0x4242, // BB Barbados
	0x4442, // BD Bangladesh
	0x4542, // BE Belgium
	0x4642, // BF Burkina Faso
	0x4742, // BG Bulgaria
	0x4842, // BH Bahrain
	0x4942, // BI Burundi
	0x4a42, // BJ Benin
	0x4d42, // BM Bermuda
	0x4e42, // BN Brunei Darussalam
	0x4f42, // BO Bolivia
	0x5242, // BR Brazil
	0x5342, // BS Bahamas
	0x5442, // BT Bhutan
	0x5642, // BV Bouvet Island
	0x5742, // BW Botswana
	0x5942, // BY Belarus
	0x5a42, // BZ Belize
	0x4143, // CA Canada
	0x4343, // CC Cocos (Keeling) Islands
	0x4443, // CD Congo, The Democratic Republic of them (still only Congo to osu!)
	0x4643, // CF Central African Republic
	0x4743, // CG Congo
	0x4843, // CH Switzerland
	0x4943, // CI Cote d'Ivoire
	0x4b43, // CK Cook Islands
	0x4c43, // CL Chile
	0x4d43, // CM Cameroon
	0x4e43, // CN China
	0x4f43, // CO Colombia
	0x5243, // CR Costa Rica
	0x5543, // CU Cuba
	0x5643, // CV Cape Verde
	0x5843, // CX Christmas Island
	0x5943, // CY Cyprus
	0x5a43, // CZ Czech Republic
	0x4544, // DE Germany
	0x4a44, // DJ Djibouti
	0x4b44, // DK Denmark
	0x4d44, // DM Dominica
	0x4f44, // DO Dominican Republic
	0x5a44, // DZ Algeria
	0x4345, // EC Ecuador
	0x4545, // EE Estonia
	0x4745, // EG Egypt
	0x4845, // EH Western Sahara
	0x5245, // ER Eritrea
	0x5345, // ES Spain
	0x5445, // ET Ethiopia
	0x4946, // FI Finland
	0x4a46, // FJ Fiji
	0x4b46, // FK Falkland Islands (Malvinas)
	0x4d46, // FM Micronesia, Federated States of
	0x4f46, // FO Faroe Islands
	0x5246, // FR France
	0x5846, // FX Metropolitan France
	0x4147, // GA Gabon
	0x4247, // GB United Kingdom
	0x4447, // GD Grenada
	0x4547, // GE Georgia
	0x4647, // GF French Guiana
	0x4847, // GH Ghana
	0x4947, // GI Gibraltar
	0x4c47, // GL Greenland
	0x4d47, // GM Gambia
	0x4e47, // GN Guinea
	0x5047, // GP Guadeloupe
	0x5147, // GQ Equatorial Guinea
	0x5247, // GR Greece
	0x5347, // GS South Georgia and the South Sandwich Islands
	0x5447, // GT Guatemala
	0x5547, // GU Guam
	0x5747, // GW Guinea-Bissau
	0x5947, // GY Guyana
	0x4b48, // HK Hong Kong
	0x4d48, // HM Heard Island and McDonald Islands
	0x4e48, // HN Honduras
	0x5248, // HR Croatia
	0x5448, // HT Haiti
	0x5548, // HU Hungary
	0x4449, // ID Indonesia
	0x4549, // IE Ireland
	0x4c49, // IL Israel
	0x4e49, // IN India
	0x4f49, // IO British Indian Ocean Territory
	0x5149, // IQ Iraq
	0x5249, // IR Iran, Islamic Republic of
	0x5349, // IS Iceland
	0x5449, // IT Italy
	0x4d4a, // JM Jamaica
	0x4f4a, // JO Jordan
	0x504a, // JP Japan
	0x454b, // KE Kenya
	0x474b, // KG Kyrgyzstan
	0x484b, // KH Cambodia
	0x494b, // KI Kiribati
	0x4d4b, // KM Comoros
	0x4e4b, // KN Saint Kitts and Nevis
	0x504b, // KP Korea, Democratic People's Republic of
	0x524b, // KR Korea, Republic of
	0x574b, // KW Kuwait
	0x594b, // KY Cayman Islands
	0x5a4b, // KZ Kazakhstan
	0x414c, // LA Lao People's Democratic Republic
	0x424c, // LB Lebanon
	0x434c, // LC Saint Lucia
	0x494c, // LI Liechtenstein
	0x4b4c, // LK Sri Lanka
	0x524c, // LR Liberia
	0x534c, // LS Lesotho
	0x544c, // LT Lithuania
	0x554c, // LU Luxembourg
	0x564c, // LV Latvia
	0x594c, // LY Libyan Arab Jamahiriya
	0x414d, // MA Morocco
	0x434d, // MC Monaco
	0x444d, // MD Moldova, Republic of
	0x474d, // MG Madagascar
	0x484d, // MH Marshall Islands
	0x4b4d, // MK Macedonia
	0x4c4d, // ML Mali
	0x4d4d, // MM Myanmar
	0x4e4d, // MN Mongolia
	0x4f4d, // MO Macao
	0x504d, // MP Northern Mariana Islands
	0x514d, // MQ Martinique
	0x524d, // MR Mauritania
	0x534d, // MS Montserrat
	0x544d, // MT Malta
	0x554d, // MU Mauritius
	0x564d, // MV Maldives
	0x574d, // MW Malawi
	0x584d, // MX Mexico
	0x594d, // MY Malaysia
	0x5a4d, // MZ Mozambique
	0x414e, // NA Namibia
	0x434e, // NC New Caledonia
	0x454e, // NE Niger
	0x464e, // NF Norfolk Island
	0x474e, // NG Nigeria
	0x494e, // NI Nicaragua
	0x4c4e, // NL Netherlands
	0x4f4e, // NO Norway
	0x504e, // NP Nepal
	0x524e, // NR Nauru
	0x554e, // NU Niue
	0x5a4e, // NZ New Zealand
	0x4d4f, // OM Oman
	0x4150, // PA Panama
	0x4550, // PE Peru
	0x4650, // PF French Polynesia
	0x4750, // PG Papua New Guinea
	0x4850, // PH Philippines
	0x4b50, // PK Pakistan
	0x4c50, // PL Poland
	0x4d50, // PM Saint Pierre and Miquelon
	0x4e50, // PN Pitcairn
	0x5250, // PR Puerto Rico
	0x5350, // PS Palestinian Territory
	0x5450, // PT Portugal
	0x5750, // PW Palau
	0x5950, // PY Paraguay
	0x4151, // QA Qatar
	0x4552, // RE Reunion
	0x4f52, // RO Romania
	0x5552, // RU Russian Federation
	0x5752, // RW Rwanda
	0x4153, // SA Saudi Arabia
	0x4253, // SB Solomon Islands
	0x4353, // SC Seychelles
	0x4453, // SD Sudan
	0x4553, // SE Sweden
	0x4753, // SG Singapore
	0x4853, // SH Saint Helena
	0x4953, // SI Slovenia
	0x4a53, // SJ Svalbard and Jan Mayen
	0x4b53, // SK Slovakia
	0x4c53, // SL Sierra Leone
	0x4d53, // SM San Marino
	0x4e53, // SN Senegal
	0x4f53, // SO Somalia
	0x5253, // SR Suriname
	0x5453, // ST Sao Tome and Principe
	0x5653, // SV El Salvador
	0x5953, // SY Syrian Arab Republic
	0x5a53, // SZ Swaziland
	0x4354, // TC Turks and Caicos Islands
	0x4454, // TD Chad
	0x4654, // TF French Southern Territories
	0x4754, // TG Togo
	0x4854, // TH Thailand
	0x4a54, // TJ Tajikistan
	0x4b54, // TK Tokelau
	0x4d54, // TM Turkmenistan
	0x4e54, // TN Tunisia
	0x4f54, // TO Tonga
	0x4c54, // TL Timor-Leste
	0x5254, // TR Turkey
	0x5454, // TT Trinidad and Tobago
	0x5654, // TV Tuvalu
	0x5754, // TW Taiwan
	0x5a54, // TZ Tanzania, United Republic of
	0x4155, // UA Ukraine
	0x4755, // UG Uganda
	0x4d55, // UM United States Minor Outlying Islands
	0x5355, // US United States
	0x5955, // UY Uruguay
	0x5a55, // UZ Uzbekistan
	0x4156, // VA Holy See (Vatican City State)
	0x4356, // VC Saint Vincent and the Grenadines
	0x4556, // VE Venezuela
	0x4756, // VG Virgin Islands, British
	0x4956, // VI Virgin Islands, U.S.
	0x4e56, // VN Vietnam
	0x5556, // VU Vanuatu
	0x4657, // WF Wallis and Futuna
	0x5357, // WS Samoa
	0x4559, // YE Yemen
	0x5459, // YT Mayotte
	0x5352, // RS Serbia
	0x415a, // ZA South Africa
	0x4d5a, // ZM Zambia
	0x454d, // ME Montenegro
	0x575a, // ZW Zimbabwe
	0x5858, // XX Unknown
	0x3241, // A2 Satellite Provider
	0x314f, // O1 Other Country
	0x5841, // AX Aland Islands
	0x4747, // GG Guernsey
	0x4d49, // IM Isle of Man
	0x454a, // JE Jersey
	0x4c42, // BL Saint Barthelemy
	0x464d, // MF Saint Martin
}

func CountryCode(i byte) string {
	slice := struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(&countryCodes[i]), 2, 2}
	return *(*string)(unsafe.Pointer(&slice))
}

func CountryByte(code string) byte {
	if len(code) != 2 {
		return 0
	}

	v := *(*int16)(unsafe.Pointer(&[]byte(code)[0]))

	for i := byte(0); i < 253; i++ {
		if countryCodes[i] == v {
			return i
		}
	}
	return 0
}
