package handlers

import (
	"goscraper/src/helpers"
	"goscraper/src/types"
	"log"
)

func GetUser(token string) (*types.User, error) {
	scraper := helpers.NewCoursePage(token)
	page, err := scraper.GetPage()
	if err != nil {
		log.Fatal(err)
	}

	user, err := helpers.GetUser(page)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil

}
