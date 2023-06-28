package main

import "fmt"

func main() {
	//Set Variables
	var totalSpaces int
	var parks int
	var communityCentres int
	var publicBeaches int
	var publicPlaygrounds int
	
	//Prompt User to enter total number of public spaces
	fmt.Println("Please enter the total number of public spaces in your area:")
	fmt.Scan(&totalSpaces)
	
	//Prompt user to enter number of parks
	fmt.Println("Please enter the number of parks in your area:")
	fmt.Scan(&parks)
	
	//Prompt user to enter number of community centres
	fmt.Println("Please enter the number of community centres in your area:")
	fmt.Scan(&communityCentres)
	
	//Prompt user to enter number of public beaches
	fmt.Println("Please enter the number of public beaches in your area:")
	fmt.Scan(&publicBeaches)
	
	//Prompt user to enter number of public playgrounds
	fmt.Println("Please enter the number of public playgrounds in your area:")
	fmt.Scan(&publicPlaygrounds)
	
	//Calculate remaining
	remainingSpaces := totalSpaces - parks - communityCentres - publicBeaches - publicPlaygrounds
	
	//Print Stats
	fmt.Println("Total Number of Public Spaces:", totalSpaces)
	fmt.Println("Number of Parks:", parks)
	fmt.Println("Number of Community Centres:", communityCentres)
	fmt.Println("Number of Public Beaches:", publicBeaches)
	fmt.Println("Number of Public Playgrounds:", publicPlaygrounds)
	fmt.Println("Remaining Public Spaces:", remainingSpaces)
	
	//Calculate Percentage of Each Space
	parkPercent := float64(parks) / float64(totalSpaces) * 100
	communityPercent := float64(communityCentres) / float64(totalSpaces) * 100
	beachPercent := float64(publicBeaches) / float64(totalSpaces) * 100
	playgroundPercent := float64(publicPlaygrounds) / float64(totalSpaces) * 100
	remainingPercent := float64(remainingSpaces) / float64(totalSpaces) * 100
	
	//Print Percentage of Each Space
	fmt.Printf("\nPercentage of Parks: %.2f%%\n", parkPercent)
	fmt.Printf("Percentage of Community Centres: %.2f%%\n", communityPercent)
	fmt.Printf("Percentage of Public Beaches: %.2f%%\n", beachPercent)
	fmt.Printf("Percentage of Public Playgrounds: %.2f%%\n", playgroundPercent)
	fmt.Printf("Percentage of Remaining Public Spaces: %.2f%%\n", remainingPercent)

}