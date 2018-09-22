package html

import "github.com/alisdairrankine/nevis/dom"

var (

	//H1 represents the H1 heading tag
	H1 = dom.ElementCreator("h1")

	//H2 represents the H2 heading tag
	H2 = dom.ElementCreator("h2")

	//H3 represents the H3 heading tag
	H3 = dom.ElementCreator("h3")

	//H4 represents the H4 heading tag
	H4 = dom.ElementCreator("h4")

	//H5 represents the H5 heading tag
	H5 = dom.ElementCreator("h5")

	//H6 represents the H6 heading tag
	H6 = dom.ElementCreator("h6")

	A = dom.ElementCreator("a")
	P = dom.ElementCreator("p")

	Div  = dom.ElementCreator("div")
	Span = dom.ElementCreator("span")

	Img = dom.ElementCreator("img")

	Ol = dom.ElementCreator("ol")
	Ul = dom.ElementCreator("ul")
	Li = dom.ElementCreator("li")

	Table = dom.ElementCreator("table")
	Tr    = dom.ElementCreator("tr")
	Td    = dom.ElementCreator("td")
	Th    = dom.ElementCreator("th")

	Button = dom.ElementCreator("button")
)
