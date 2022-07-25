// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package unicode

import (
	q "unicode"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "unicode",
		Path:       "unicode",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CaseRange":   reflect.TypeOf((*q.CaseRange)(nil)).Elem(),
			"Range16":     reflect.TypeOf((*q.Range16)(nil)).Elem(),
			"Range32":     reflect.TypeOf((*q.Range32)(nil)).Elem(),
			"RangeTable":  reflect.TypeOf((*q.RangeTable)(nil)).Elem(),
			"SpecialCase": reflect.TypeOf((*q.SpecialCase)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ASCII_Hex_Digit":                    reflect.ValueOf(&q.ASCII_Hex_Digit),
			"Adlam":                              reflect.ValueOf(&q.Adlam),
			"Ahom":                               reflect.ValueOf(&q.Ahom),
			"Anatolian_Hieroglyphs":              reflect.ValueOf(&q.Anatolian_Hieroglyphs),
			"Arabic":                             reflect.ValueOf(&q.Arabic),
			"Armenian":                           reflect.ValueOf(&q.Armenian),
			"Avestan":                            reflect.ValueOf(&q.Avestan),
			"AzeriCase":                          reflect.ValueOf(&q.AzeriCase),
			"Balinese":                           reflect.ValueOf(&q.Balinese),
			"Bamum":                              reflect.ValueOf(&q.Bamum),
			"Bassa_Vah":                          reflect.ValueOf(&q.Bassa_Vah),
			"Batak":                              reflect.ValueOf(&q.Batak),
			"Bengali":                            reflect.ValueOf(&q.Bengali),
			"Bhaiksuki":                          reflect.ValueOf(&q.Bhaiksuki),
			"Bidi_Control":                       reflect.ValueOf(&q.Bidi_Control),
			"Bopomofo":                           reflect.ValueOf(&q.Bopomofo),
			"Brahmi":                             reflect.ValueOf(&q.Brahmi),
			"Braille":                            reflect.ValueOf(&q.Braille),
			"Buginese":                           reflect.ValueOf(&q.Buginese),
			"Buhid":                              reflect.ValueOf(&q.Buhid),
			"C":                                  reflect.ValueOf(&q.C),
			"Canadian_Aboriginal":                reflect.ValueOf(&q.Canadian_Aboriginal),
			"Carian":                             reflect.ValueOf(&q.Carian),
			"CaseRanges":                         reflect.ValueOf(&q.CaseRanges),
			"Categories":                         reflect.ValueOf(&q.Categories),
			"Caucasian_Albanian":                 reflect.ValueOf(&q.Caucasian_Albanian),
			"Cc":                                 reflect.ValueOf(&q.Cc),
			"Cf":                                 reflect.ValueOf(&q.Cf),
			"Chakma":                             reflect.ValueOf(&q.Chakma),
			"Cham":                               reflect.ValueOf(&q.Cham),
			"Cherokee":                           reflect.ValueOf(&q.Cherokee),
			"Chorasmian":                         reflect.ValueOf(&q.Chorasmian),
			"Co":                                 reflect.ValueOf(&q.Co),
			"Common":                             reflect.ValueOf(&q.Common),
			"Coptic":                             reflect.ValueOf(&q.Coptic),
			"Cs":                                 reflect.ValueOf(&q.Cs),
			"Cuneiform":                          reflect.ValueOf(&q.Cuneiform),
			"Cypriot":                            reflect.ValueOf(&q.Cypriot),
			"Cyrillic":                           reflect.ValueOf(&q.Cyrillic),
			"Dash":                               reflect.ValueOf(&q.Dash),
			"Deprecated":                         reflect.ValueOf(&q.Deprecated),
			"Deseret":                            reflect.ValueOf(&q.Deseret),
			"Devanagari":                         reflect.ValueOf(&q.Devanagari),
			"Diacritic":                          reflect.ValueOf(&q.Diacritic),
			"Digit":                              reflect.ValueOf(&q.Digit),
			"Dives_Akuru":                        reflect.ValueOf(&q.Dives_Akuru),
			"Dogra":                              reflect.ValueOf(&q.Dogra),
			"Duployan":                           reflect.ValueOf(&q.Duployan),
			"Egyptian_Hieroglyphs":               reflect.ValueOf(&q.Egyptian_Hieroglyphs),
			"Elbasan":                            reflect.ValueOf(&q.Elbasan),
			"Elymaic":                            reflect.ValueOf(&q.Elymaic),
			"Ethiopic":                           reflect.ValueOf(&q.Ethiopic),
			"Extender":                           reflect.ValueOf(&q.Extender),
			"FoldCategory":                       reflect.ValueOf(&q.FoldCategory),
			"FoldScript":                         reflect.ValueOf(&q.FoldScript),
			"Georgian":                           reflect.ValueOf(&q.Georgian),
			"Glagolitic":                         reflect.ValueOf(&q.Glagolitic),
			"Gothic":                             reflect.ValueOf(&q.Gothic),
			"Grantha":                            reflect.ValueOf(&q.Grantha),
			"GraphicRanges":                      reflect.ValueOf(&q.GraphicRanges),
			"Greek":                              reflect.ValueOf(&q.Greek),
			"Gujarati":                           reflect.ValueOf(&q.Gujarati),
			"Gunjala_Gondi":                      reflect.ValueOf(&q.Gunjala_Gondi),
			"Gurmukhi":                           reflect.ValueOf(&q.Gurmukhi),
			"Han":                                reflect.ValueOf(&q.Han),
			"Hangul":                             reflect.ValueOf(&q.Hangul),
			"Hanifi_Rohingya":                    reflect.ValueOf(&q.Hanifi_Rohingya),
			"Hanunoo":                            reflect.ValueOf(&q.Hanunoo),
			"Hatran":                             reflect.ValueOf(&q.Hatran),
			"Hebrew":                             reflect.ValueOf(&q.Hebrew),
			"Hex_Digit":                          reflect.ValueOf(&q.Hex_Digit),
			"Hiragana":                           reflect.ValueOf(&q.Hiragana),
			"Hyphen":                             reflect.ValueOf(&q.Hyphen),
			"IDS_Binary_Operator":                reflect.ValueOf(&q.IDS_Binary_Operator),
			"IDS_Trinary_Operator":               reflect.ValueOf(&q.IDS_Trinary_Operator),
			"Ideographic":                        reflect.ValueOf(&q.Ideographic),
			"Imperial_Aramaic":                   reflect.ValueOf(&q.Imperial_Aramaic),
			"Inherited":                          reflect.ValueOf(&q.Inherited),
			"Inscriptional_Pahlavi":              reflect.ValueOf(&q.Inscriptional_Pahlavi),
			"Inscriptional_Parthian":             reflect.ValueOf(&q.Inscriptional_Parthian),
			"Javanese":                           reflect.ValueOf(&q.Javanese),
			"Join_Control":                       reflect.ValueOf(&q.Join_Control),
			"Kaithi":                             reflect.ValueOf(&q.Kaithi),
			"Kannada":                            reflect.ValueOf(&q.Kannada),
			"Katakana":                           reflect.ValueOf(&q.Katakana),
			"Kayah_Li":                           reflect.ValueOf(&q.Kayah_Li),
			"Kharoshthi":                         reflect.ValueOf(&q.Kharoshthi),
			"Khitan_Small_Script":                reflect.ValueOf(&q.Khitan_Small_Script),
			"Khmer":                              reflect.ValueOf(&q.Khmer),
			"Khojki":                             reflect.ValueOf(&q.Khojki),
			"Khudawadi":                          reflect.ValueOf(&q.Khudawadi),
			"L":                                  reflect.ValueOf(&q.L),
			"Lao":                                reflect.ValueOf(&q.Lao),
			"Latin":                              reflect.ValueOf(&q.Latin),
			"Lepcha":                             reflect.ValueOf(&q.Lepcha),
			"Letter":                             reflect.ValueOf(&q.Letter),
			"Limbu":                              reflect.ValueOf(&q.Limbu),
			"Linear_A":                           reflect.ValueOf(&q.Linear_A),
			"Linear_B":                           reflect.ValueOf(&q.Linear_B),
			"Lisu":                               reflect.ValueOf(&q.Lisu),
			"Ll":                                 reflect.ValueOf(&q.Ll),
			"Lm":                                 reflect.ValueOf(&q.Lm),
			"Lo":                                 reflect.ValueOf(&q.Lo),
			"Logical_Order_Exception":            reflect.ValueOf(&q.Logical_Order_Exception),
			"Lower":                              reflect.ValueOf(&q.Lower),
			"Lt":                                 reflect.ValueOf(&q.Lt),
			"Lu":                                 reflect.ValueOf(&q.Lu),
			"Lycian":                             reflect.ValueOf(&q.Lycian),
			"Lydian":                             reflect.ValueOf(&q.Lydian),
			"M":                                  reflect.ValueOf(&q.M),
			"Mahajani":                           reflect.ValueOf(&q.Mahajani),
			"Makasar":                            reflect.ValueOf(&q.Makasar),
			"Malayalam":                          reflect.ValueOf(&q.Malayalam),
			"Mandaic":                            reflect.ValueOf(&q.Mandaic),
			"Manichaean":                         reflect.ValueOf(&q.Manichaean),
			"Marchen":                            reflect.ValueOf(&q.Marchen),
			"Mark":                               reflect.ValueOf(&q.Mark),
			"Masaram_Gondi":                      reflect.ValueOf(&q.Masaram_Gondi),
			"Mc":                                 reflect.ValueOf(&q.Mc),
			"Me":                                 reflect.ValueOf(&q.Me),
			"Medefaidrin":                        reflect.ValueOf(&q.Medefaidrin),
			"Meetei_Mayek":                       reflect.ValueOf(&q.Meetei_Mayek),
			"Mende_Kikakui":                      reflect.ValueOf(&q.Mende_Kikakui),
			"Meroitic_Cursive":                   reflect.ValueOf(&q.Meroitic_Cursive),
			"Meroitic_Hieroglyphs":               reflect.ValueOf(&q.Meroitic_Hieroglyphs),
			"Miao":                               reflect.ValueOf(&q.Miao),
			"Mn":                                 reflect.ValueOf(&q.Mn),
			"Modi":                               reflect.ValueOf(&q.Modi),
			"Mongolian":                          reflect.ValueOf(&q.Mongolian),
			"Mro":                                reflect.ValueOf(&q.Mro),
			"Multani":                            reflect.ValueOf(&q.Multani),
			"Myanmar":                            reflect.ValueOf(&q.Myanmar),
			"N":                                  reflect.ValueOf(&q.N),
			"Nabataean":                          reflect.ValueOf(&q.Nabataean),
			"Nandinagari":                        reflect.ValueOf(&q.Nandinagari),
			"Nd":                                 reflect.ValueOf(&q.Nd),
			"New_Tai_Lue":                        reflect.ValueOf(&q.New_Tai_Lue),
			"Newa":                               reflect.ValueOf(&q.Newa),
			"Nko":                                reflect.ValueOf(&q.Nko),
			"Nl":                                 reflect.ValueOf(&q.Nl),
			"No":                                 reflect.ValueOf(&q.No),
			"Noncharacter_Code_Point":            reflect.ValueOf(&q.Noncharacter_Code_Point),
			"Number":                             reflect.ValueOf(&q.Number),
			"Nushu":                              reflect.ValueOf(&q.Nushu),
			"Nyiakeng_Puachue_Hmong":             reflect.ValueOf(&q.Nyiakeng_Puachue_Hmong),
			"Ogham":                              reflect.ValueOf(&q.Ogham),
			"Ol_Chiki":                           reflect.ValueOf(&q.Ol_Chiki),
			"Old_Hungarian":                      reflect.ValueOf(&q.Old_Hungarian),
			"Old_Italic":                         reflect.ValueOf(&q.Old_Italic),
			"Old_North_Arabian":                  reflect.ValueOf(&q.Old_North_Arabian),
			"Old_Permic":                         reflect.ValueOf(&q.Old_Permic),
			"Old_Persian":                        reflect.ValueOf(&q.Old_Persian),
			"Old_Sogdian":                        reflect.ValueOf(&q.Old_Sogdian),
			"Old_South_Arabian":                  reflect.ValueOf(&q.Old_South_Arabian),
			"Old_Turkic":                         reflect.ValueOf(&q.Old_Turkic),
			"Oriya":                              reflect.ValueOf(&q.Oriya),
			"Osage":                              reflect.ValueOf(&q.Osage),
			"Osmanya":                            reflect.ValueOf(&q.Osmanya),
			"Other":                              reflect.ValueOf(&q.Other),
			"Other_Alphabetic":                   reflect.ValueOf(&q.Other_Alphabetic),
			"Other_Default_Ignorable_Code_Point": reflect.ValueOf(&q.Other_Default_Ignorable_Code_Point),
			"Other_Grapheme_Extend":              reflect.ValueOf(&q.Other_Grapheme_Extend),
			"Other_ID_Continue":                  reflect.ValueOf(&q.Other_ID_Continue),
			"Other_ID_Start":                     reflect.ValueOf(&q.Other_ID_Start),
			"Other_Lowercase":                    reflect.ValueOf(&q.Other_Lowercase),
			"Other_Math":                         reflect.ValueOf(&q.Other_Math),
			"Other_Uppercase":                    reflect.ValueOf(&q.Other_Uppercase),
			"P":                                  reflect.ValueOf(&q.P),
			"Pahawh_Hmong":                       reflect.ValueOf(&q.Pahawh_Hmong),
			"Palmyrene":                          reflect.ValueOf(&q.Palmyrene),
			"Pattern_Syntax":                     reflect.ValueOf(&q.Pattern_Syntax),
			"Pattern_White_Space":                reflect.ValueOf(&q.Pattern_White_Space),
			"Pau_Cin_Hau":                        reflect.ValueOf(&q.Pau_Cin_Hau),
			"Pc":                                 reflect.ValueOf(&q.Pc),
			"Pd":                                 reflect.ValueOf(&q.Pd),
			"Pe":                                 reflect.ValueOf(&q.Pe),
			"Pf":                                 reflect.ValueOf(&q.Pf),
			"Phags_Pa":                           reflect.ValueOf(&q.Phags_Pa),
			"Phoenician":                         reflect.ValueOf(&q.Phoenician),
			"Pi":                                 reflect.ValueOf(&q.Pi),
			"Po":                                 reflect.ValueOf(&q.Po),
			"Prepended_Concatenation_Mark":       reflect.ValueOf(&q.Prepended_Concatenation_Mark),
			"PrintRanges":                        reflect.ValueOf(&q.PrintRanges),
			"Properties":                         reflect.ValueOf(&q.Properties),
			"Ps":                                 reflect.ValueOf(&q.Ps),
			"Psalter_Pahlavi":                    reflect.ValueOf(&q.Psalter_Pahlavi),
			"Punct":                              reflect.ValueOf(&q.Punct),
			"Quotation_Mark":                     reflect.ValueOf(&q.Quotation_Mark),
			"Radical":                            reflect.ValueOf(&q.Radical),
			"Regional_Indicator":                 reflect.ValueOf(&q.Regional_Indicator),
			"Rejang":                             reflect.ValueOf(&q.Rejang),
			"Runic":                              reflect.ValueOf(&q.Runic),
			"S":                                  reflect.ValueOf(&q.S),
			"STerm":                              reflect.ValueOf(&q.STerm),
			"Samaritan":                          reflect.ValueOf(&q.Samaritan),
			"Saurashtra":                         reflect.ValueOf(&q.Saurashtra),
			"Sc":                                 reflect.ValueOf(&q.Sc),
			"Scripts":                            reflect.ValueOf(&q.Scripts),
			"Sentence_Terminal":                  reflect.ValueOf(&q.Sentence_Terminal),
			"Sharada":                            reflect.ValueOf(&q.Sharada),
			"Shavian":                            reflect.ValueOf(&q.Shavian),
			"Siddham":                            reflect.ValueOf(&q.Siddham),
			"SignWriting":                        reflect.ValueOf(&q.SignWriting),
			"Sinhala":                            reflect.ValueOf(&q.Sinhala),
			"Sk":                                 reflect.ValueOf(&q.Sk),
			"Sm":                                 reflect.ValueOf(&q.Sm),
			"So":                                 reflect.ValueOf(&q.So),
			"Soft_Dotted":                        reflect.ValueOf(&q.Soft_Dotted),
			"Sogdian":                            reflect.ValueOf(&q.Sogdian),
			"Sora_Sompeng":                       reflect.ValueOf(&q.Sora_Sompeng),
			"Soyombo":                            reflect.ValueOf(&q.Soyombo),
			"Space":                              reflect.ValueOf(&q.Space),
			"Sundanese":                          reflect.ValueOf(&q.Sundanese),
			"Syloti_Nagri":                       reflect.ValueOf(&q.Syloti_Nagri),
			"Symbol":                             reflect.ValueOf(&q.Symbol),
			"Syriac":                             reflect.ValueOf(&q.Syriac),
			"Tagalog":                            reflect.ValueOf(&q.Tagalog),
			"Tagbanwa":                           reflect.ValueOf(&q.Tagbanwa),
			"Tai_Le":                             reflect.ValueOf(&q.Tai_Le),
			"Tai_Tham":                           reflect.ValueOf(&q.Tai_Tham),
			"Tai_Viet":                           reflect.ValueOf(&q.Tai_Viet),
			"Takri":                              reflect.ValueOf(&q.Takri),
			"Tamil":                              reflect.ValueOf(&q.Tamil),
			"Tangut":                             reflect.ValueOf(&q.Tangut),
			"Telugu":                             reflect.ValueOf(&q.Telugu),
			"Terminal_Punctuation":               reflect.ValueOf(&q.Terminal_Punctuation),
			"Thaana":                             reflect.ValueOf(&q.Thaana),
			"Thai":                               reflect.ValueOf(&q.Thai),
			"Tibetan":                            reflect.ValueOf(&q.Tibetan),
			"Tifinagh":                           reflect.ValueOf(&q.Tifinagh),
			"Tirhuta":                            reflect.ValueOf(&q.Tirhuta),
			"Title":                              reflect.ValueOf(&q.Title),
			"TurkishCase":                        reflect.ValueOf(&q.TurkishCase),
			"Ugaritic":                           reflect.ValueOf(&q.Ugaritic),
			"Unified_Ideograph":                  reflect.ValueOf(&q.Unified_Ideograph),
			"Upper":                              reflect.ValueOf(&q.Upper),
			"Vai":                                reflect.ValueOf(&q.Vai),
			"Variation_Selector":                 reflect.ValueOf(&q.Variation_Selector),
			"Wancho":                             reflect.ValueOf(&q.Wancho),
			"Warang_Citi":                        reflect.ValueOf(&q.Warang_Citi),
			"White_Space":                        reflect.ValueOf(&q.White_Space),
			"Yezidi":                             reflect.ValueOf(&q.Yezidi),
			"Yi":                                 reflect.ValueOf(&q.Yi),
			"Z":                                  reflect.ValueOf(&q.Z),
			"Zanabazar_Square":                   reflect.ValueOf(&q.Zanabazar_Square),
			"Zl":                                 reflect.ValueOf(&q.Zl),
			"Zp":                                 reflect.ValueOf(&q.Zp),
			"Zs":                                 reflect.ValueOf(&q.Zs),
		},
		Funcs: map[string]reflect.Value{
			"In":         reflect.ValueOf(q.In),
			"Is":         reflect.ValueOf(q.Is),
			"IsControl":  reflect.ValueOf(q.IsControl),
			"IsDigit":    reflect.ValueOf(q.IsDigit),
			"IsGraphic":  reflect.ValueOf(q.IsGraphic),
			"IsLetter":   reflect.ValueOf(q.IsLetter),
			"IsLower":    reflect.ValueOf(q.IsLower),
			"IsMark":     reflect.ValueOf(q.IsMark),
			"IsNumber":   reflect.ValueOf(q.IsNumber),
			"IsOneOf":    reflect.ValueOf(q.IsOneOf),
			"IsPrint":    reflect.ValueOf(q.IsPrint),
			"IsPunct":    reflect.ValueOf(q.IsPunct),
			"IsSpace":    reflect.ValueOf(q.IsSpace),
			"IsSymbol":   reflect.ValueOf(q.IsSymbol),
			"IsTitle":    reflect.ValueOf(q.IsTitle),
			"IsUpper":    reflect.ValueOf(q.IsUpper),
			"SimpleFold": reflect.ValueOf(q.SimpleFold),
			"To":         reflect.ValueOf(q.To),
			"ToLower":    reflect.ValueOf(q.ToLower),
			"ToTitle":    reflect.ValueOf(q.ToTitle),
			"ToUpper":    reflect.ValueOf(q.ToUpper),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"LowerCase":       {"untyped int", constant.MakeInt64(int64(q.LowerCase))},
			"MaxASCII":        {"untyped rune", constant.MakeInt64(int64(q.MaxASCII))},
			"MaxCase":         {"untyped int", constant.MakeInt64(int64(q.MaxCase))},
			"MaxLatin1":       {"untyped rune", constant.MakeInt64(int64(q.MaxLatin1))},
			"MaxRune":         {"untyped rune", constant.MakeInt64(int64(q.MaxRune))},
			"ReplacementChar": {"untyped rune", constant.MakeInt64(int64(q.ReplacementChar))},
			"TitleCase":       {"untyped int", constant.MakeInt64(int64(q.TitleCase))},
			"UpperCase":       {"untyped int", constant.MakeInt64(int64(q.UpperCase))},
			"UpperLower":      {"untyped rune", constant.MakeInt64(int64(q.UpperLower))},
			"Version":         {"untyped string", constant.MakeString(string(q.Version))},
		},
	})
}
