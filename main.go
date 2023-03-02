package main

import (
	"fmt"
)

type Profile struct {
	Name 	string
	Health 	int
	Power 	int
	Exp 	int
}

func main() {
	// Profil Sasuke
	fmt.Println("//This Profile Belongs To Sasuke//")
	profileSasuke := MakeProfileSasuke("Sasuke")
	fmt.Println("Name : ", profileSasuke.Name)
	fmt.Println("Health : ", profileSasuke.Health)
	fmt.Println("Power : ", profileSasuke.Power)
	fmt.Println("Exp : ", profileSasuke.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUpSasuke := PowerUpSasuke(profileSasuke, 3)
	fmt.Println("Name : ", powerUpSasuke.Name)
	fmt.Println("Health : ", powerUpSasuke.Health)
	fmt.Println("Power : ", powerUpSasuke.Power)
	fmt.Println("Exp : ", powerUpSasuke.Exp)

	fmt.Println("================================")

	// Profil Goku
	fmt.Println("//This Profile Belongs To Goku//")
	profileGoku := MakeProfileGoku("Goku")
	fmt.Println("Name : ", profileGoku.Name)
	fmt.Println("Health : ", profileGoku.Health)
	fmt.Println("Power : ", profileGoku.Power)
	fmt.Println("Exp : ", profileGoku.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUpGoku := PowerUpGoku(profileGoku, 3)
	fmt.Println("Name : ", powerUpGoku.Name)
	fmt.Println("Health : ", powerUpGoku.Health)
	fmt.Println("Power : ", powerUpGoku.Power)
	fmt.Println("Exp : ", powerUpGoku.Exp)

	fmt.Println("================================")

	// Profil Naruto
	fmt.Println("//This Profile Belongs To Naruto//")
	profileNaruto := MakeProfileNaruto("Naruto")
	fmt.Println("Name : ", profileNaruto.Name)
	fmt.Println("Health : ", profileNaruto.Health)
	fmt.Println("Power : ", profileNaruto.Power)
	fmt.Println("Exp : ", profileNaruto.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUpNaruto := PowerUpNaruto(profileNaruto, 3)
	fmt.Println("Name : ", powerUpNaruto.Name)
	fmt.Println("Health : ", powerUpNaruto.Health)
	fmt.Println("Power : ", powerUpNaruto.Power)
	fmt.Println("Exp : ", powerUpNaruto.Exp)
}

// Function Sasuke
func MakeProfileSasuke(character string) Profile {
	var profileSasuke = Profile{}
	if character == "Sasuke" {
		profileSasuke.Name = "Sasuke"
		profileSasuke.Health = 200
		profileSasuke.Power = 100
		profileSasuke.Exp = 0
		return profileSasuke
	} 
	
	return profileSasuke
}

func PowerUpSasuke(prosedur Profile, multiplier int) Profile {
	prosedur.Health = prosedur.Health + (prosedur.Health * multiplier)
	prosedur.Power = prosedur.Power + (prosedur.Power * multiplier)
	prosedur.Exp = prosedur.Exp + (prosedur.Exp * multiplier)
	return prosedur
}

// Function Goku
func MakeProfileGoku(character string) Profile {
	var profileGoku = Profile{}
	if character == "Goku" {
		profileGoku.Name = "Goku"
		profileGoku.Health = 400
		profileGoku.Power = 300
		profileGoku.Exp = 100
		return profileGoku
	} 
	
	return profileGoku
}

func PowerUpGoku(prosedur Profile, multiplier int) Profile {
	prosedur.Health = prosedur.Health + (prosedur.Health * multiplier)
	prosedur.Power = prosedur.Power + (prosedur.Power * multiplier)
	prosedur.Exp = prosedur.Exp + (prosedur.Exp * multiplier)
	return prosedur
}

// Function Naruto
func MakeProfileNaruto(character string) Profile {
	var profileNaruto = Profile{}
	if character == "Naruto" {
		profileNaruto.Name = "Naruto"
		profileNaruto.Health = 150
		profileNaruto.Power = 200
		profileNaruto.Exp = 50
		return profileNaruto
	} 
	
	return profileNaruto
}

func PowerUpNaruto(prosedur Profile, multiplier int) Profile {
	prosedur.Health = prosedur.Health + (prosedur.Health * multiplier)
	prosedur.Power = prosedur.Power + (prosedur.Power * multiplier)
	prosedur.Exp = prosedur.Exp + (prosedur.Exp * multiplier)
	return prosedur
}