package ui

import (
	cui "github.com/UncleJunVIP/nextui-pak-shared-functions/ui"
	"github.com/scalysoot/nextui-pak-store/models"
	"github.com/scalysoot/nextui-pak-store/state"
	"qlova.tech/sum"
	"slices"
)

type PakList struct {
	AppState state.AppState
	Category string
}

func InitPakList(appState state.AppState, category string) PakList {
	return PakList{
		AppState: appState,
		Category: category,
	}
}

func (pl PakList) Name() sum.Int[models.ScreenName] {
	return models.ScreenNames.PakList
}

func (pl PakList) Draw() (selection models.ScreenReturn, exitCode int, e error) {
	title := pl.Category
	options := models.MenuItems{Items: []string{}}
	for _, p := range pl.AppState.BrowsePaks[pl.Category] {
		options.Items = append(options.Items, p.StorefrontName)
	}

	slices.Sort(options.Items)

	s, err := cui.DisplayList(options, title, "")
	if err != nil {
		return nil, -1, err
	}

	selectedPak := pl.AppState.BrowsePaks[pl.Category][s.SelectedValue]

	return selectedPak, s.ExitCode, nil
}
