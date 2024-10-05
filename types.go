package speedhive

import (
	"encoding/json"
	"strings"
	"time"
)

type SportCategoryValue string

const (
	SportCategoryValue_Motorized SportCategoryValue = "Motorized"
	SportCategoryValue_Active    SportCategoryValue = "Active"
)

type SportValue string

const (
	SportValue_Unknown         SportValue = "Unknown"
	SportValue_Karting         SportValue = "Karting"
	SportValue_MX              SportValue = "MX"
	SportValue_Car             SportValue = "Car"
	SportValue_Bike            SportValue = "Bike"
	SportValue_RC              SportValue = "RC"
	SportValue_Cycling         SportValue = "Cycling"
	SportValue_IceSkating      SportValue = "IceSkating"
	SportValue_Running         SportValue = "Running"
	SportValue_ModelBoatRacing SportValue = "ModelBoatRacing"
	SportValue_InlineSkating   SportValue = "InlineSkating"
	SportValue_StockCar        SportValue = "StockCar"
	SportValue_Swimming        SportValue = "Swimming"
	SportValue_Other           SportValue = "Other"
)

type CountryValue string

const (
	CountryValue_PuertoRico                             CountryValue = "PR"
	CountryValue_Palestine                              CountryValue = "PS"
	CountryValue_Palau                                  CountryValue = "PW"
	CountryValue_Portugal                               CountryValue = "PT"
	CountryValue_Pakistan                               CountryValue = "PK"
	CountryValue_Philippines                            CountryValue = "PH"
	CountryValue_Pitcairn                               CountryValue = "PN"
	CountryValue_Poland                                 CountryValue = "PL"
	CountryValue_SaintPierreAndMiquelon                 CountryValue = "PM"
	CountryValue_Panama                                 CountryValue = "PA"
	CountryValue_FrenchPolynesia                        CountryValue = "PF"
	CountryValue_PapuaNewGuinea                         CountryValue = "PG"
	CountryValue_Peru                                   CountryValue = "PE"
	CountryValue_Andorra                                CountryValue = "AD"
	CountryValue_AntiguaAndBarbuda                      CountryValue = "AG"
	CountryValue_UnitedArabEmirates                     CountryValue = "AE"
	CountryValue_Afghanistan                            CountryValue = "AF"
	CountryValue_Qatar                                  CountryValue = "QA"
	CountryValue_Paraguay                               CountryValue = "PY"
	CountryValue_Nepal                                  CountryValue = "NP"
	CountryValue_Norway                                 CountryValue = "NO"
	CountryValue_Niue                                   CountryValue = "NU"
	CountryValue_Nauru                                  CountryValue = "NR"
	CountryValue_Nicaragua                              CountryValue = "NI"
	CountryValue_NorfolkIsland                          CountryValue = "NF"
	CountryValue_Nigeria                                CountryValue = "NG"
	CountryValue_Netherlands                            CountryValue = "NL"
	CountryValue_Namibia                                CountryValue = "NA"
	CountryValue_Niger                                  CountryValue = "NE"
	CountryValue_NewCaledonia                           CountryValue = "NC"
	CountryValue_Malawi                                 CountryValue = "MW"
	CountryValue_Mexico                                 CountryValue = "MX"
	CountryValue_Mauritius                              CountryValue = "MU"
	CountryValue_Maldives                               CountryValue = "MV"
	CountryValue_Malaysia                               CountryValue = "MY"
	CountryValue_Mozambique                             CountryValue = "MZ"
	CountryValue_Oman                                   CountryValue = "OM"
	CountryValue_NewZealand                             CountryValue = "NZ"
	CountryValue_Tuvalu                                 CountryValue = "TV"
	CountryValue_Taiwan                                 CountryValue = "TW"
	CountryValue_TrinidadAndTobago                      CountryValue = "TT"
	CountryValue_Germany                                CountryValue = "DE"
	CountryValue_Djibouti                               CountryValue = "DJ"
	CountryValue_Tanzania                               CountryValue = "TZ"
	CountryValue_Denmark                                CountryValue = "DK"
	CountryValue_Tunisia                                CountryValue = "TN"
	CountryValue_Tonga                                  CountryValue = "TO"
	CountryValue_TimorLeste                             CountryValue = "TL"
	CountryValue_Turkmenistan                           CountryValue = "TM"
	CountryValue_Turkey                                 CountryValue = "TR"
	CountryValue_Cuba                                   CountryValue = "CU"
	CountryValue_FrenchSouthernTerritories              CountryValue = "TF"
	CountryValue_CapeVerde                              CountryValue = "CV"
	CountryValue_Togo                                   CountryValue = "TG"
	CountryValue_Chad                                   CountryValue = "TD"
	CountryValue_Cyprus                                 CountryValue = "CY"
	CountryValue_Tajikistan                             CountryValue = "TJ"
	CountryValue_CzechRepublic                          CountryValue = "CZ"
	CountryValue_Tokelau                                CountryValue = "TK"
	CountryValue_Curacao                                CountryValue = "CW"
	CountryValue_Thailand                               CountryValue = "TH"
	CountryValue_ChristmasIsland                        CountryValue = "CX"
	CountryValue_Cameroon                               CountryValue = "CM"
	CountryValue_China                                  CountryValue = "CN"
	CountryValue_CookIslands                            CountryValue = "CK"
	CountryValue_Chile                                  CountryValue = "CL"
	CountryValue_CostaRica                              CountryValue = "CR"
	CountryValue_TurksAndCaicosIslands                  CountryValue = "TC"
	CountryValue_Colombia                               CountryValue = "CO"
	CountryValue_Egypt                                  CountryValue = "EG"
	CountryValue_WesternSahara                          CountryValue = "EH"
	CountryValue_Estonia                                CountryValue = "EE"
	CountryValue_Uruguay                                CountryValue = "UY"
	CountryValue_Uzbekistan                             CountryValue = "UZ"
	CountryValue_UnitedStatesMinorOutlyingIslands       CountryValue = "UM"
	CountryValue_Ecuador                                CountryValue = "EC"
	CountryValue_UnitedStates                           CountryValue = "US"
	CountryValue_Uganda                                 CountryValue = "UG"
	CountryValue_Algeria                                CountryValue = "DZ"
	CountryValue_DominicanRepublic                      CountryValue = "DO"
	CountryValue_Dominica                               CountryValue = "DM"
	CountryValue_Ukraine                                CountryValue = "UA"
	CountryValue_Bangladesh                             CountryValue = "BD"
	CountryValue_Belgium                                CountryValue = "BE"
	CountryValue_RussianFederation                      CountryValue = "RU"
	CountryValue_Barbados                               CountryValue = "BB"
	CountryValue_Serbia                                 CountryValue = "RS"
	CountryValue_Bahrain                                CountryValue = "BH"
	CountryValue_Burundi                                CountryValue = "BI"
	CountryValue_BurkinaFaso                            CountryValue = "BF"
	CountryValue_Bulgaria                               CountryValue = "BG"
	CountryValue_Rwanda                                 CountryValue = "RW"
	CountryValue_Azerbaijan                             CountryValue = "AZ"
	CountryValue_BosniaAndHerzegovina                   CountryValue = "BA"
	CountryValue_Romania                                CountryValue = "RO"
	CountryValue_AmericanSamoa                          CountryValue = "AS"
	CountryValue_Austria                                CountryValue = "AT"
	CountryValue_Reunion                                CountryValue = "RE"
	CountryValue_Antarctica                             CountryValue = "AQ"
	CountryValue_Argentina                              CountryValue = "AR"
	CountryValue_Aruba                                  CountryValue = "AW"
	CountryValue_AlandIslands                           CountryValue = "AX"
	CountryValue_Australia                              CountryValue = "AU"
	CountryValue_Albania                                CountryValue = "AL"
	CountryValue_Anguilla                               CountryValue = "AI"
	CountryValue_Angola                                 CountryValue = "AO"
	CountryValue_Armenia                                CountryValue = "AM"
	CountryValue_NetherlandsAntilles                    CountryValue = "AN"
	CountryValue_CentralAfricanRepublic                 CountryValue = "CF"
	CountryValue_ElSalvador                             CountryValue = "SV"
	CountryValue_CocosIslands                           CountryValue = "CC"
	CountryValue_SouthSudan                             CountryValue = "SS"
	CountryValue_CongoDR                                CountryValue = "CD"
	CountryValue_SaoTomeAndPrincipe                     CountryValue = "ST"
	CountryValue_CoteDIvoire                            CountryValue = "CI"
	CountryValue_SyrianArabRepublic                     CountryValue = "SY"
	CountryValue_Swaziland                              CountryValue = "SZ"
	CountryValue_Congo                                  CountryValue = "CG"
	CountryValue_Switzerland                            CountryValue = "CH"
	CountryValue_SintMaarten                            CountryValue = "SX"
	CountryValue_SanMarino                              CountryValue = "SM"
	CountryValue_Senegal                                CountryValue = "SN"
	CountryValue_Belize                                 CountryValue = "BZ"
	CountryValue_Slovakia                               CountryValue = "SK"
	CountryValue_SierraLeone                            CountryValue = "SL"
	CountryValue_Canada                                 CountryValue = "CA"
	CountryValue_Suriname                               CountryValue = "SR"
	CountryValue_Somalia                                CountryValue = "SO"
	CountryValue_Bhutan                                 CountryValue = "BT"
	CountryValue_Sweden                                 CountryValue = "SE"
	CountryValue_Brazil                                 CountryValue = "BR"
	CountryValue_Seychelles                             CountryValue = "SC"
	CountryValue_Bahamas                                CountryValue = "BS"
	CountryValue_Sudan                                  CountryValue = "SD"
	CountryValue_Slovenia                               CountryValue = "SI"
	CountryValue_Belarus                                CountryValue = "BY"
	CountryValue_SvalbardAndJanMayen                    CountryValue = "SJ"
	CountryValue_BouvetIsland                           CountryValue = "BV"
	CountryValue_Singapore                              CountryValue = "SG"
	CountryValue_Botswana                               CountryValue = "BW"
	CountryValue_SaintHelena                            CountryValue = "SH"
	CountryValue_SaintBarthelemy                        CountryValue = "BL"
	CountryValue_Bermuda                                CountryValue = "BM"
	CountryValue_Benin                                  CountryValue = "BJ"
	CountryValue_SaudiArabia                            CountryValue = "SA"
	CountryValue_Bonaire                                CountryValue = "BQ"
	CountryValue_SolomonIslands                         CountryValue = "SB"
	CountryValue_BruneiDarussalam                       CountryValue = "BN"
	CountryValue_Bolivia                                CountryValue = "BO"
	CountryValue_HongKong                               CountryValue = "HK"
	CountryValue_Honduras                               CountryValue = "HN"
	CountryValue_HeardIslandAndMcdonaldIslands          CountryValue = "HM"
	CountryValue_Guyana                                 CountryValue = "GY"
	CountryValue_GuineaBissau                           CountryValue = "GW"
	CountryValue_EquatorialGuinea                       CountryValue = "GQ"
	CountryValue_Greece                                 CountryValue = "GR"
	CountryValue_Guadeloupe                             CountryValue = "GP"
	CountryValue_Guam                                   CountryValue = "GU"
	CountryValue_SouthGeorgiaAndTheSouthSandwichIslands CountryValue = "GS"
	CountryValue_Guatemala                              CountryValue = "GT"
	CountryValue_Israel                                 CountryValue = "IL"
	CountryValue_BritishIndianOceanTerritory            CountryValue = "IO"
	CountryValue_SouthAfrica                            CountryValue = "ZA"
	CountryValue_IsleOfMan                              CountryValue = "IM"
	CountryValue_India                                  CountryValue = "IN"
	CountryValue_CanaryIslands                          CountryValue = "IC"
	CountryValue_Indonesia                              CountryValue = "ID"
	CountryValue_Mayotte                                CountryValue = "YT"
	CountryValue_Ireland                                CountryValue = "IE"
	CountryValue_Croatia                                CountryValue = "HR"
	CountryValue_Haiti                                  CountryValue = "HT"
	CountryValue_Yemen                                  CountryValue = "YE"
	CountryValue_Hungary                                CountryValue = "HU"
	CountryValue_Finland                                CountryValue = "FI"
	CountryValue_Micronesia                             CountryValue = "FM"
	CountryValue_Fiji                                   CountryValue = "FJ"
	CountryValue_FalklandIslands                        CountryValue = "FK"
	CountryValue_Vietnam                                CountryValue = "VN"
	CountryValue_Vanuatu                                CountryValue = "VU"
	CountryValue_VirginIslandsUS                        CountryValue = "VI"
	CountryValue_VirginIslandsBritish                   CountryValue = "VG"
	CountryValue_HolySee                                CountryValue = "VA"
	CountryValue_Spain                                  CountryValue = "ES"
	CountryValue_Ethiopia                               CountryValue = "ET"
	CountryValue_Venezuela                              CountryValue = "VE"
	CountryValue_Eritrea                                CountryValue = "ER"
	CountryValue_SaintVincentAndTheGrenadines           CountryValue = "VC"
	CountryValue_Gibraltar                              CountryValue = "GI"
	CountryValue_Guernsey                               CountryValue = "GG"
	CountryValue_Ghana                                  CountryValue = "GH"
	CountryValue_Gambia                                 CountryValue = "GM"
	CountryValue_Guinea                                 CountryValue = "GN"
	CountryValue_Greenland                              CountryValue = "GL"
	CountryValue_Gabon                                  CountryValue = "GA"
	CountryValue_UnitedKingdom                          CountryValue = "GB"
	CountryValue_Georgia                                CountryValue = "GE"
	CountryValue_FrenchGuiana                           CountryValue = "GF"
	CountryValue_Samoa                                  CountryValue = "WS"
	CountryValue_Grenada                                CountryValue = "GD"
	CountryValue_FaroeIslands                           CountryValue = "FO"
	CountryValue_WallisAndFutuna                        CountryValue = "WF"
	CountryValue_France                                 CountryValue = "FR"
	CountryValue_Liberia                                CountryValue = "LR"
	CountryValue_Lesotho                                CountryValue = "LS"
	CountryValue_SriLanka                               CountryValue = "LK"
	CountryValue_Liechtenstein                          CountryValue = "LI"
	CountryValue_Lebanon                                CountryValue = "LB"
	CountryValue_SaintLucia                             CountryValue = "LC"
	CountryValue_LaoPDR                                 CountryValue = "LA"
	CountryValue_CaymanIslands                          CountryValue = "KY"
	CountryValue_Kazakhstan                             CountryValue = "KZ"
	CountryValue_Kuwait                                 CountryValue = "KW"
	CountryValue_Macao                                  CountryValue = "MO"
	CountryValue_NorthernMarianaIslands                 CountryValue = "MP"
	CountryValue_Myanmar                                CountryValue = "MM"
	CountryValue_Mongolia                               CountryValue = "MN"
	CountryValue_Montserrat                             CountryValue = "MS"
	CountryValue_Malta                                  CountryValue = "MT"
	CountryValue_Martinique                             CountryValue = "MQ"
	CountryValue_Mauritania                             CountryValue = "MR"
	CountryValue_Madagascar                             CountryValue = "MG"
	CountryValue_MarshallIslands                        CountryValue = "MH"
	CountryValue_Montenegro                             CountryValue = "ME"
	CountryValue_SaintMartin                            CountryValue = "MF"
	CountryValue_Macedonia                              CountryValue = "MK"
	CountryValue_Mali                                   CountryValue = "ML"
	CountryValue_Monaco                                 CountryValue = "MC"
	CountryValue_Moldova                                CountryValue = "MD"
	CountryValue_Morocco                                CountryValue = "MA"
	CountryValue_Latvia                                 CountryValue = "LV"
	CountryValue_Lithuania                              CountryValue = "LT"
	CountryValue_Luxembourg                             CountryValue = "LU"
	CountryValue_Libya                                  CountryValue = "LY"
	CountryValue_Jamaica                                CountryValue = "JM"
	CountryValue_Japan                                  CountryValue = "JP"
	CountryValue_Jordan                                 CountryValue = "JO"
	CountryValue_Jersey                                 CountryValue = "JE"
	CountryValue_Zimbabwe                               CountryValue = "ZW"
	CountryValue_Zambia                                 CountryValue = "ZM"
	CountryValue_Iceland                                CountryValue = "IS"
	CountryValue_Italy                                  CountryValue = "IT"
	CountryValue_Iraq                                   CountryValue = "IQ"
	CountryValue_Iran                                   CountryValue = "IR"
	CountryValue_Comoros                                CountryValue = "KM"
	CountryValue_SaintKittsAndNevis                     CountryValue = "KN"
	CountryValue_KoreaRepublicOf                        CountryValue = "KR"
	CountryValue_KoreaDPR                               CountryValue = "KP"
	CountryValue_Kenya                                  CountryValue = "KE"
	CountryValue_Kiribati                               CountryValue = "KI"
	CountryValue_Kyrgyzstan                             CountryValue = "KG"
	CountryValue_Cambodia                               CountryValue = "KH"
)

type OrderValue string

const (
	OrderValue_Ascending  OrderValue = "asc"
	OrderValue_Descending OrderValue = "desc"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	t := time.Duration(0)

	// Posibilities are:
	// 1. 1:23:45.678
	// 2. 23:45.678
	// 3. 45.678
	// 4. 0.678
	// And the negative versions of the above

	negative := false
	if strings.HasPrefix(value, "-") {
		negative = true
		value = value[1:]
	}

	parts := strings.Split(value, ":")
	switch len(parts) {
	case 1:
		t, _ = time.ParseDuration(value + "s")
	case 2:
		t, _ = time.ParseDuration(parts[0] + "m" + parts[1] + "s")
	case 3:
		t, _ = time.ParseDuration(parts[0] + "h" + parts[1] + "m" + parts[2] + "s")
	}

	if negative {
		t *= -1
	}

	*d = Duration(t)
	return nil
}

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// Time may be in the formats:
	// - 2006-01-02
	// - 2006-01-02T15:04:05
	// - 2006-01-02T15:04:05+07:00

	var err error
	var tm time.Time

	if strings.Contains(value, "T") {
		tm, err = time.Parse(time.RFC3339, value)
		if err != nil {
			tm, err = time.Parse("2006-01-02T15:04:05", value)
		}
	} else {
		tm, err = time.Parse("2006-01-02", value)
	}

	if err != nil {
		return err
	}

	*t = Time(tm)
	return nil
}
