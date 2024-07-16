package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomExercises(slice []string, n int) []string {
    rand.Seed(time.Now().UnixNano())
    randomExercises := make([]string, n)
    for i := 0; i < n; i++ {
        index := rand.Intn(len(slice))
        randomExercises[i] = slice[index]
    }
    return randomExercises
}

func main() {
	trePeito := []string{
        "Supino inclinado barra",
        "Supino declinado",
        "Supino reto",
        "Crucifixo peck deck",
        "Crucifixo banco reto com halteres",
        "Cross over",
        "Flexão básica",
        "Peito na paralela",
        "Voador no cabo com banco inclinado",
    }

	treInferiores := []string{
        "Agachamento Livre",
        "Agachamento com pulo",
        "Afundo",
        "Agachamento no hack",
        "Agachamento hack inverso",
        "Cadeira extensora",
        "Cadeira flexora",
        "Leg press",
        "Mesa flexora",
        "Cadeira adutora",
        "Agachamento Smith",
        "Cadeira flexora unilateral",
        "Cadeira extensora unilateral",
        "Cadeira abdutora",
    }
	
	treOmbros := []string{
        "Elevação lateral c/ halteres",
        "Elevação frontal c/ halteres",
        "Elevação lateral c/ haltere tronco inclinado",
        "Elevação frontal corda polia baixa",
        "Desenvolvimento máquina",
        "Desenvolvimento c/ halteres",
        "Elevação lateral tronco inclinado",
    }

	treTriceps := []string{
        "Tríceps mergulho no banco",
        "Tríceps mergulho pernas estendidas",
        "Tríceps coice",
        "Tríceps francês unilateral",
        "Tríceps testa",
        "Flexão de braço",
        "Flexão de braço fechado",
        "Flexão de braço em isometria",
    }

	treCostas := []string{
        "Remada polia baixa",
        "Remada cavalinho",
        "Puxada aberta",
        "Puxada supinada",
        "Puxada triângulo",
        "Puxada aberta neutra",
        "Serrote",
        "Face pull",
        "Remada hammer neutra",
        "Remada hammer pronada",
        "Remada curvada barra curva",
    }
	
    treBiceps := []string{
        "Rosca direta",
        "Rosca alternada com halteres",
        "Rosca concentrada",
        "Rosca martelo",
        "Rosca Scott",
        "Rosca no cabo",
        "Rosca 21",
        "Rosca inclinada",
        "Rosca de concentração com halteres",
        "Rosca de punho",
    }

	dia := "C"

	if dia == "A" {
	
        selectedPeito := getRandomExercises(trePeito, 3)
        selectedOmbros := getRandomExercises(treOmbros, 2)
        selectedTriceps := getRandomExercises(treTriceps, 2)
        
        selectedExercises := append(selectedPeito, selectedOmbros...)
        selectedExercises = append(selectedExercises, selectedTriceps...)
        
        fmt.Println("Exercícios selecionados:")
        for _, exercise := range selectedExercises{
            fmt.Println(exercise)
        }
    
    }else if dia == "B" {
        selectedCostas := getRandomExercises(treCostas, 4)
        selectedBiceps := getRandomExercises(treBiceps, 3)
        
        selectedExercises := append(selectedCostas, selectedBiceps...)
        
        fmt.Println("Exercícios selecionados:")
        for _, exercise := range selectedExercises{
            fmt.Println(exercise)
        }
    
    }else if dia == "C" {
        selectedInferiores := getRandomExercises(treInferiores, 4)

        selectedExercises := append(selectedInferiores)

        fmt.Println("Exercícios selecionados:")
        for _, exercise := range selectedExercises{
            fmt.Println(exercise)
        }
    }else {
        fmt.Println("Opção Inválida.")
    }

	
	
}	