package main

import "fmt"

func fuelGauge(fuel int) {
  fmt.Println(fuel)
}

func calculateFuel(planet string) int {
var fuel int

if planet== "Venus" {
  fuel=300000
} else if planet == "Mercury" {
  fuel=500000
} else if planet == "Mars" {
  fuel = 700000
} 

return fuel
} 

func greetPlanet(planet string) {
  fmt.Println("Welcome to "+ planet)
}

func cantfly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int{
  var fuelRemaining int
  var fuelCost int
  fuelRemaining = fuel
  fuelCost=calculateFuel(planet)

if(fuelRemaining >= fuelCost) {
  greetPlanet(planet)
  fuelCost-=fuelRemaining
} else {
  cantfly()
}

return fuelRemaining
}

func main() {
  // Test your functions!
  fuel := 1000000
  planetChoice := "Mercury"
  // Create `planetChoize` and `fuel`
  fuel =flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  // And then liftoff!
  
}
