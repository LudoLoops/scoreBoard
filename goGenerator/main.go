// generate.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	UserID         string   `json:"userId"`
	Nom            string   `json:"nom"`
	Note           int      `json:"note"`
	Bonus          int      `json:"bonus,omitempty"`
	XP             int      `json:"xp"`
	DateSoumission string   `json:"dateSoumission"`
	Tags           []string `json:"tags,omitempty"`
}

type Projet struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Users []User `json:"users"`
}

var possibleTags = []string{
	"créativité", "leadership", "travail d'équipe",
	"puntualité", "autonomie", "rigueur",
	"communication", "innovation", "proactivité",
}

var projectNames []string

func init() {
	// gofakeit.Seed(0)
}

func randomTags() []string {
	count := 1 + rand.Intn(3)
	tags := make(map[string]bool)
	var result []string
	for len(result) < count {
		tag := possibleTags[rand.Intn(len(possibleTags))]
		if !tags[tag] {
			tags[tag] = true
			result = append(result, tag)
		}
	}
	return result
}

func randomDate() string {
	daysAgo := rand.Intn(20)
	return time.Now().AddDate(0, 0, -daysAgo).Format("2006-01-02")
}

func generateSlug(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func main() {
	var maxUsers int
	var maxProjects int
	flag.IntVar(&maxUsers, "u", 50, "nombre d'utilisateurs à générer")

	flag.IntVar(&maxProjects, "p", 10, "nombre de projets à générer")
	flag.Parse()

	if maxProjects == 0 {
		log.Fatal("❌ Le nombre projets ne doit pas être 0")
	}

	if maxUsers < 40 {
		log.Fatal("❌ Le nombre d'utilisateurs doit être > 40")
	}

	fmt.Printf("📊 Génération pour %d utilisateurs\n", maxUsers)
	// Générer utilisateurs uniques
	users := make(map[string]string)
	var userIds []string

	for i := range maxUsers {
		userID := strconv.Itoa(i)
		name := gofakeit.Gamertag()
		users[userID] = name
		userIds = append(userIds, userID)
	}

	// Générer projets
	var data []Projet
	for p := 1; p <= maxProjects; p++ {

		projetKey := gofakeit.ProductName()
		slug := generateSlug(projetKey)
		minUsers := int(float64(maxUsers) * 0.6)
		playerCount := minUsers + rand.Intn(maxUsers-minUsers+1)

		// Mélanger les utilisateurs
		rand.Shuffle(len(userIds), func(i, j int) {
			userIds[i], userIds[j] = userIds[j], userIds[i]
		})

		var userlist []User
		for i := range playerCount {

			userID := userIds[i]
			nom := users[userID]
			note := 10 + rand.Intn(41)
			bonus := rand.Intn(21)
			xp := note + bonus

			userlist = append(userlist, User{
				UserID:         userID,
				Nom:            nom,
				Note:           note,
				Bonus:          bonus,
				XP:             xp,
				DateSoumission: randomDate(),
				Tags:           randomTags(),
			})
		}
		projet := Projet{
			Name:  projetKey,
			Slug:  slug,
			Users: userlist,
		}
		data = append(data, projet)
	}

	// ✅ Écrire equipage.json DANS LE DOSSIER COURANT
	file, err := os.Create("equipage.json")
	if err != nil {
		log.Fatal("Erreur lors de la création du fichier equipage.json:", err)
	}
	defer file.Close()

	// Formater le JSON avec une indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)
	if err != nil {
		log.Fatal("Erreur lors de l'écriture du JSON:", err)
	}

	// ✅ Confirmation claire
	fmt.Printf("✅  %d projets, %d utilisateurs générés!\n", maxProjects, maxUsers)
	fmt.Println("➡️  Fichier : ./equipage.json")
}
