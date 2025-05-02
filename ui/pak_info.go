package ui

import (
	"context"
	"fmt"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/ui"
	"github.com/scalysoot/nextui-pak-store/database"
	"github.com/scalysoot/nextui-pak-store/models"
	"github.com/scalysoot/nextui-pak-store/utils"
	"go.uber.org/zap"
	"path/filepath"
	"qlova.tech/sum"
	"strings"
)

type PakInfoScreen struct {
	Pak       models.Pak
	Category  string
	Installed bool
}

func InitPakInfoScreen(pak models.Pak, category string, installed bool) PakInfoScreen {
	return PakInfoScreen{
		Pak:       pak,
		Category:  category,
		Installed: installed,
	}
}

func (pi PakInfoScreen) Name() sum.Int[models.ScreenName] {
	return models.ScreenNames.PakInfo
}

func (pi PakInfoScreen) Draw() (selection models.ScreenReturn, exitCode int, e error) {
	logger := common.GetLoggerInstance()

	showBanner := true

	banner := pi.Pak.RepoURL + models.RefMainStub + pi.Pak.Banners["BRICK"]
	banner = strings.ReplaceAll(banner, models.GitHubRoot, models.RawGHUC)
	bannerFile, err := utils.DownloadTempFile(banner)
	if err != nil {
		showBanner = false
	}

	var message string
	var options []string

	if showBanner {
		options = []string{
			"--background-image", bannerFile,
			"--cancel-button", "Y",
			"--action-button", "B",
			"--action-text", "BACK",
			"--action-show", "true",
			"--message-alignment", "bottom"}
		message = models.BlankPresenterString
	} else {
		options = []string{
			"--cancel-button", "Y",
			"--action-button", "B",
			"--action-text", "BACK",
			"--action-show", "true",
			"--message-alignment", "middle"}
		message = fmt.Sprintf("%s: %s", pi.Pak.StorefrontName, pi.Pak.Description)
	}

	if !pi.Installed {
		options = append(options, "--confirm-text", "INSTALL", "--confirm-show", "true", "--confirm-button", "X")
	} else {
		message = "Installed!"
	}

	code, err := ui.ShowMessageWithOptions(message, "0", options...)
	if err != nil {
		return nil, -1, err
	}

	if pi.Pak.LargePak {
		code, err = ui.ShowMessageWithOptions("Heads up! This is a very large download!", "0",
			"--cancel-button", "B", "--cancel-show", "false", "--cancel-text", "CANCEL",
			"--confirm-show", "--confirm-text", "I UNDERSTAND")
		if err != nil {
			return nil, -1, err
		}
	}

	if code == 0 {
		tmp, err := utils.DownloadPakArchive(pi.Pak, "Installing")
		if err != nil {
			logger.Error("Unable to download pak archive", zap.Error(err))
			return nil, -1, err
		}

		pakDestination := ""

		if pi.Pak.PakType == models.PakTypes.TOOL {
			pakDestination = filepath.Join(models.ToolRoot, pi.Pak.Name+".pak")
		} else if pi.Pak.PakType == models.PakTypes.EMU {
			pakDestination = filepath.Join(models.EmulatorRoot, pi.Pak.Name+".pak")
		}

		err = utils.Unzip(tmp, pakDestination, pi.Pak, false)
		if err != nil {
			return nil, -1, err
		}

		info := database.InstallParams{
			DisplayName:  pi.Pak.StorefrontName,
			Name:         pi.Pak.Name,
			Version:      pi.Pak.Version,
			Type:         models.PakTypeMap[pi.Pak.PakType],
			CanUninstall: int64(1),
		}

		ctx := context.Background()
		err = database.DBQ().Install(ctx, info)
		if err != nil {
			// TODO wtf do I do here?
		}
	}

	return nil, code, nil
}
