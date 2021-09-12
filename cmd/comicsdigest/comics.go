package main

func comics(title string) {
	HTMLInit(title)

	cOrigin := "class=\"fancybox\""
	cToken := "img src"
	ckOrigin := ""
	ckToken := "data-comic-image-url"

	Generic(cOrigin, cToken, "Andy Capp", "https://www.creators.com/read/andy-capp")
	Generic(cOrigin, cToken, "B.C.", "https://www.creators.com/read/bc")
	Generic(ckOrigin, ckToken, "Beetle Bailey", "https://www.comicskingdom.com/beetle-bailey/")
	GoComics("Betty", "betty")
	GoComics("Broom Hilda", "broomhilda")
	GoComics("Buckles", "buckles")
	GoComics("Calvin and Hobbes", "calvinandhobbes")
	Generic(ckOrigin, ckToken, "Crock", "https://www.comicskingdom.com/crock/")
	Dilbert()
	Explosm()
	GoComics("Garfield", "garfield")
	Generic(ckOrigin, ckToken, "Hagar", "https://www.comicskingdom.com/hagar-the-horrible/")
	GoComics("Peanuts", "peanuts")
	GoComics("Pearls before Swine", "pearlsbeforeswine")
	Generic(ckOrigin, ckToken, "Sam and Silo", "https://www.comicskingdom.com/sam-and-silo/")
	Generic(ckOrigin, ckToken, "Zits", "https://www.comicskingdom.com/zits")
}
