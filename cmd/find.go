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

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	raw, err := ioutil.ReadFile("./data/pokemon.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var monsters []Monster
	json.Unmarshal(raw, &monsters)
	for _, monster := range monsters {
		fmt.Println("--------------------------------------------")
		fmt.Println("No." + strconv.Itoa(monster.Number))
		fmt.Println("Name: " + monster.Name)
		fmt.Println("H: " + strconv.Itoa(monster.HitPoint))
		fmt.Println("A: " + strconv.Itoa(monster.Attack))
		fmt.Println("B: " + strconv.Itoa(monster.Defense))
		fmt.Println("C: " + strconv.Itoa(monster.SpecialAttack))
		fmt.Println("D: " + strconv.Itoa(monster.SpecialDefense))
		fmt.Println("S: " + strconv.Itoa(monster.Speed))
		fmt.Printf("types: %v\n", monster.Types)
		fmt.Println("--------------------------------------------")
	}
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
