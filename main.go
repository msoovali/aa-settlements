package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/msoovali/aa-settlements/adapter/archiver"
	"github.com/msoovali/aa-settlements/adapter/googledrive"
	"github.com/msoovali/aa-settlements/adapter/googlesheet"
	"github.com/msoovali/aa-settlements/adapter/translations"
	"github.com/msoovali/aa-settlements/config"
	"github.com/msoovali/aa-settlements/domain"
	"github.com/msoovali/aa-settlements/domain/clock"
	"github.com/msoovali/aa-settlements/domain/localizer"
	"github.com/msoovali/aa-settlements/domain/port"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	log.Println("Next month settlement creation started!")
	conf, err := config.LoadConfiguration("config.json")
	if err != nil {
		log.Fatalf("Unable to read config.json file: %v", err)
	}
	ctx := context.Background()
	credentials, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read credentials.json file: %v", err)
	}
	googleApiConfig, err := google.JWTConfigFromJSON(credentials, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	googleApiClient := googleApiConfig.Client(ctx)

	ports, localizer, err := createPortsAndLocalizer(googleApiClient, conf)
	if err != nil {
		log.Fatalf("Failed to create ports: %v", err)
	}
	service, err := domain.New(ports, conf, clock.RealClock{}, localizer)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	err = service.CreateNextMonthSettlements()
	if err != nil {
		log.Fatalf("Failed to create next month settlements: %v", err)
	}
	log.Println("Next month settlements prepared successfully!")
}

func createPortsAndLocalizer(googleApiClient *http.Client, config *config.Config) (*port.Ports, *localizer.Localizer, error) {
	driveSrv, err := drive.NewService(context.Background(), option.WithHTTPClient(googleApiClient))
	if err != nil {
		return nil, nil, err
	}
	driveService := googledrive.NewDriveGateway(driveSrv, config.RootFolderId)
	sheetSrv, err := sheets.NewService(context.Background(), option.WithHTTPClient(googleApiClient))
	if err != nil {
		return nil, nil, err
	}
	sheetService := googlesheet.NewSheetGateway(sheetSrv)
	translationsService := translations.NewTranslationsGateway()
	localizer, err := translationsService.Get(config.Locale)
	archiverService := archiver.New(driveService, *config, localizer)

	return &port.Ports{
		DrivePort:        driveService,
		SheetPort:        sheetService,
		TranslationsPort: translationsService,
		ArchiverPort:     archiverService,
	}, localizer, err
}
