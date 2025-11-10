package main

import (
	"fmt"
	// "log"
	// "net/url"
	// "os"
)



func main() {

	CPU_CORES := 10
	var choice int;

	fmt.Println("Welcome to the PageRank Algorithm Implementation in GoLang!");
	fmt.Println("Please choose one of the following options to proceed:");

	fmt.Println("1 -- Low Level Stimulated Directed Graph of Webpages. N = 20, E = 62");
	fmt.Println("2 -- Medium Level Stimulated Directed Graph of Webpages using the 'WEB-STANFORD.txt' Dataset.");
	fmt.Println("3 -- High Level Stimulated Directed Graph of Webpages using the 'WEB-GOOGLE.txt' Dataset.");

	fmt.Scanln(&choice);

	fmt.Println("The PageRank Algorithm will run in parallel using", CPU_CORES, "CPU cores.");

	AdjacencyList adjacencyList = constructGraphFromDataset(choice);

	


}
