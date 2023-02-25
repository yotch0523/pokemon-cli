/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var name string

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "'pokemon-cli' is a CLI to search for Pokemon by various filter",
	Long: `'pokemon-cli' is a CLI to sarch for Pokemon
`,
	Run: func(cmd *cobra.Command, args []string) {
		getPokemonData()
	},
}

type MonsterCollection struct {
	Monsters []Monster
}

type Monster struct {
	Number         int    `json:"number"`
	Name           string `json:"name"`
	HitPoint       int    `json:"H"`
	Attack         int    `json:"A"`
	Defense        int    `json:"B"`
	SpecialAttack  int    `json:"C"`
	SpecialDefense int    `json:"D"`
	Speed          int    `json:"S"`
	Types          []Type `json:"types"`
}

type Type string

func getPokemonData() {
	if name == "" {
		fmt.Println("No condition is specified")
		os.Exit(1)
	}
	raw, err := ioutil.ReadFile("./data/pokemon.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var monsters []Monster
	json.Unmarshal(raw, &monsters)
	var results []*Monster
	for i := 0; i < len(monsters); i++ {
		if strings.Contains(monsters[i].Name, name) {
			results = append(results, &monsters[i])
		}
	}

	for _, result := range results {
		fmt.Println("--------------------------------------------")
		fmt.Println("No." + strconv.Itoa(result.Number))
		fmt.Println("Name: " + result.Name)
		fmt.Println("H: " + strconv.Itoa(result.HitPoint))
		fmt.Println("A: " + strconv.Itoa(result.Attack))
		fmt.Println("B: " + strconv.Itoa(result.Defense))
		fmt.Println("C: " + strconv.Itoa(result.SpecialAttack))
		fmt.Println("D: " + strconv.Itoa(result.SpecialDefense))
		fmt.Println("S: " + strconv.Itoa(result.Speed))
		fmt.Printf("types: %v\n", result.Types)
		fmt.Println("--------------------------------------------")
	}
}

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().StringVar(&name, "name", "", "Pokemon name condition. Executes a partial match search for Pokemon name value specified in this flag.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
