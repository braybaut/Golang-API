package main

var books = Books{
	book{ID: 0,
		Name:        "Infrastructure As Code",
		Description: "this is a great “definitive” book about the concept of Infrastructure As Code (IaC). IaC is a set of principles and patterns for managing infrastructure in the cloud. This book is focused more on general principles and ideas than tools.",
		Autor:       "Kief Morris"},

	book{ID: 1,
		Name:        "Site Reliability Engineering",
		Description: "Members of the SRE team explain how their engagement with the entire software lifecycle has enabled Google to build, deploy, monitor, and maintain some of the largest software systems in the world.",
		Autor:       " etsy Beyer, Chris Jones, Jennifer Petoff and Niall Richard Murphy"},
	book{ID: 2,
		Name:        "The Go Programming Language",
		Description: "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go. It shows how to write clear and idiomatic Go to solve real-world problems. The book does not assume prior knowledge of Go nor experience with any specific language, so you’ll find it accessible whether you’re most comfortable with JavaScript, Ruby, Python, Java, or C++.",
		Autor:       "Alan A. A. Donovan, Brian W. Kernighan"},
}
