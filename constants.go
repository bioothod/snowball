/*
	Constants used in the English Snowball stemmer.
*/
package snowball

const vowels string = "aeiouy"

var doubleConsonants = [...]string{"bb", "dd", "ff", "gg", "mm", "nn",
	"pp", "rr", "tt"}
var liEnding = "cdeghkmnrt"
var step0Suffixes = [...]string{"'s'", "'s", "'"}
var step1aSuffixes = [...]string{"sses", "ied", "ies", "us", "ss", "s"}
var step1bSuffixes = [...]string{"eedly", "ingly", "edly", "eed", "ing", "ed"}
var step2Suffixes = [...]string{
	"ization", "ational", "fulness", "ousness",
	"iveness", "tional", "biliti", "lessli",
	"entli", "ation", "alism", "aliti", "ousli",
	"iviti", "fulli", "enci", "anci", "abli",
	"izer", "ator", "alli", "bli", "ogi", "li",
}
var step3Suffixes = [...]string{
	"ational", "tional", "alize", "icate", "iciti",
	"ative", "ical", "ness", "ful",
}
var step4Suffixes = [...]string{
	"ement", "ance", "ence", "able", "ible", "ment",
	"ant", "ent", "ism", "ate", "iti", "ous",
	"ive", "ize", "ion", "al", "er", "ic",
}
var step5Suffixes = [...]string{"e", "l"}
var specialWords = map[string]string{
	"skis":       "ski",
	"skies":      "sky",
	"dying":      "die",
	"lying":      "lie",
	"tying":      "tie",
	"idly":       "idl",
	"gently":     "gentl",
	"ugly":       "ugli",
	"early":      "earli",
	"only":       "onli",
	"singly":     "singl",
	"sky":        "sky",
	"news":       "news",
	"howe":       "howe",
	"atlas":      "atlas",
	"cosmos":     "cosmos",
	"bias":       "bias",
	"andes":      "andes",
	"inning":     "inning",
	"innings":    "inning",
	"outing":     "outing",
	"outings":    "outing",
	"canning":    "canning",
	"cannings":   "canning",
	"herring":    "herring",
	"herrings":   "herring",
	"earring":    "earring",
	"earrings":   "earring",
	"proceed":    "proceed",
	"proceeds":   "proceed",
	"proceeded":  "proceed",
	"proceeding": "proceed",
	"exceed":     "exceed",
	"exceeds":    "exceed",
	"exceeded":   "exceed",
	"exceeding":  "exceed",
	"succeed":    "succeed",
	"succeeds":   "succeed",
	"succeeded":  "succeed",
	"succeeding": "succeed",
}

var stopWords = map[string]bool{
	"a":          true,
	"about":      true,
	"above":      true,
	"after":      true,
	"again":      true,
	"against":    true,
	"all":        true,
	"am":         true,
	"an":         true,
	"and":        true,
	"any":        true,
	"are":        true,
	"as":         true,
	"at":         true,
	"be":         true,
	"because":    true,
	"been":       true,
	"before":     true,
	"being":      true,
	"below":      true,
	"between":    true,
	"both":       true,
	"but":        true,
	"by":         true,
	"can":        true,
	"did":        true,
	"do":         true,
	"does":       true,
	"doing":      true,
	"don":        true,
	"down":       true,
	"during":     true,
	"each":       true,
	"few":        true,
	"for":        true,
	"from":       true,
	"further":    true,
	"had":        true,
	"has":        true,
	"have":       true,
	"having":     true,
	"he":         true,
	"her":        true,
	"here":       true,
	"hers":       true,
	"herself":    true,
	"him":        true,
	"himself":    true,
	"his":        true,
	"how":        true,
	"i":          true,
	"if":         true,
	"in":         true,
	"into":       true,
	"is":         true,
	"it":         true,
	"its":        true,
	"itself":     true,
	"just":       true,
	"me":         true,
	"more":       true,
	"most":       true,
	"my":         true,
	"myself":     true,
	"no":         true,
	"nor":        true,
	"not":        true,
	"now":        true,
	"of":         true,
	"off":        true,
	"on":         true,
	"once":       true,
	"only":       true,
	"or":         true,
	"other":      true,
	"our":        true,
	"ours":       true,
	"ourselves":  true,
	"out":        true,
	"over":       true,
	"own":        true,
	"s":          true,
	"same":       true,
	"she":        true,
	"should":     true,
	"so":         true,
	"some":       true,
	"such":       true,
	"t":          true,
	"than":       true,
	"that":       true,
	"the":        true,
	"their":      true,
	"theirs":     true,
	"them":       true,
	"themselves": true,
	"then":       true,
	"there":      true,
	"these":      true,
	"they":       true,
	"this":       true,
	"those":      true,
	"through":    true,
	"to":         true,
	"too":        true,
	"under":      true,
	"until":      true,
	"up":         true,
	"very":       true,
	"was":        true,
	"we":         true,
	"were":       true,
	"what":       true,
	"when":       true,
	"where":      true,
	"which":      true,
	"while":      true,
	"who":        true,
	"whom":       true,
	"why":        true,
	"will":       true,
	"with":       true,
	"you":        true,
	"your":       true,
	"yours":      true,
	"yourself":   true,
	"yourselves": true,
}