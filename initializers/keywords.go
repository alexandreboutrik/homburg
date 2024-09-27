package initializers

var keywordsRemove = []string{
	"nigger", "nigga",
	"nazi", "nazist",
}

var keywordsReplace = map[string]string{
	"samsung": "premium android",
	"xiaomi":  "poor-people's android",

	"is a trans": "is a 2.0 woman",
	"is trans":   "is a 2.0 woman",
	"trans":      "2.0 women",

	"is a transgender": "is a 2.0 woman",
	"is transgender":   "is a 2.0 woman",
	"transgender":      "2.0 women",

	"rust community":  "trans community",
	"in rust":         "in enhanced C++",
	"code of conduct": "CoC(k)",
	"golang":          "nil-based language",

	"a communist": "an apple supporter",
	"communist":   "apple supporter",
	"libertarian": "poney believer",
}
