package nearmap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"roofix/pkg/apiaccess"
	"roofix/pkg/util/http"
	"roofix/pkg/util/log"
)

type Rollup struct {
	Desc  string      `json:"description"`
	Value interface{} `json:"value"`
}

type Response struct {
	SurveyId   string             `json:"surveyId"`
	SurveyDate string             `json:"surveyDate"`
	Link       string             `json:"link"`
	ExtentMinX float64            `json:"extentMinX"`
	ExtentMaxX float64            `json:"extentMaxX"`
	ExtentMinY float64            `json:"extentMinY"`
	ExtentMaxY float64            `json:"extentMaxY"`
	Rollups    map[string]*Rollup `json:"rollups"`
}

type NearmapAoi struct {
	Aoi *Geometry `json:"aoi"`
}

func addressRollups(ctx context.Context, streetAddress, city, state, zip, country, until string) (*Response, error) {
	// get api token detail
	host, key, err := apiToken(ctx)
	if err != nil {
		return nil, err
	}

	// call API
	// search parcel by address
	ep := fmt.Sprintf("rollups.json?order=latest&bulk=false&packs=roof_char&apikey=%s", key)
	ep = fmt.Sprintf(
		"%s&streetAddress=%s&city=%s&state=%s&zip=%s&country=%s&until=%s",
		ep,
		url.QueryEscape(streetAddress),
		url.QueryEscape(city),
		url.QueryEscape(state),
		url.QueryEscape(zip),
		url.QueryEscape(country),
		until,
	)
	b, err := http.Get(
		fmt.Sprintf("%s/%s", host, ep),
		http.WithContentTypeJson(),
	)
	if err != nil {
		return nil, err
	}

	// parse response
	var res Response
	if err := json.Unmarshal(b, &res); err != nil {
		log.Error(log.Wrap(err, "api response UNMARSHAL error"))
		return nil, nil
	}

	// done
	return &res, nil
}

func geometryRollups(ctx context.Context, geo *Geometry) (*Response, error) {
	// get api token detail
	host, key, err := apiToken(ctx)
	if err != nil {
		return nil, err
	}

	// call API
	ep := fmt.Sprintf("rollups.json?order=latest&bulk=false&packs=roof_char&apikey=%s", key)
	b, err := http.Post(
		fmt.Sprintf("%s/%s", host, ep),
		NearmapAoi{Aoi: geo},
		http.WithContentTypeJson(),
	)
	if err != nil {
		return nil, err
	}

	// parse response
	var res Response
	if err = json.Unmarshal(b, &res); err != nil {
		return nil, errors.New("failed to parse Nearmap API response")
	}

	// done
	return &res, nil
}

func apiToken(ctx context.Context) (host, key string, err error) {
	api, err := apiaccess.GetKey(ctx, ProviderName)
	if err != nil {
		return host, key, log.Wrap(err, "api access errot")
	}

	host = api.URL
	key = api.Key
	return host, key, nil
}
