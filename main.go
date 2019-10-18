package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "lcego"
	app.Author = "yasukotelin"
	app.Version = "1.0.0"
	app.Usage = "generates OSS Licence selected you one from list."
	app.Description = "lcego is OSS Licence generator CLI tool."
	app.Action = mainAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainAction(c *cli.Context) error {
	licenses, err := getLicenses()
	if err != nil {
		return err
	}
	license := askToSelectLicense(licenses)

	detail, err := getLicenseDetail(license)
	if err != nil {
		return err
	}

	writeLicenseFile(detail.Body)

	return nil
}

func askToSelectLicense(licenses []License) *License {
	for i, license := range licenses {
		fmt.Printf("[%d]\t%s\n", i, license.Name)
	}
	fmt.Println()
	fmt.Print("Select a License: ")

	licensesLen := len(licenses)-1
	for {
		var numStr string
		fmt.Scan(&numStr)

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if num < 0 || num > licensesLen {
			fmt.Println("your selected number is out of range.")
			continue
		}

		return &licenses[num]
	}
}

func writeLicenseFile(text string) error {
	file, err := os.OpenFile("LICENSE", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, text)

	return err
}
