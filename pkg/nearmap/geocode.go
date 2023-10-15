package nearmap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"roofix/ent/schema"
	"roofix/pkg/apiaccess"
	"roofix/pkg/util/http"
	"roofix/pkg/util/log"
)

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type NearGridResponse struct {
	Results []struct {
		Id         int      `json:"id"`
		Geometry   Geometry `json:"geometry"`
		Properties struct {
			Headline string `json:"headline"`
			Path     string `json:"path"`
			Fields   struct {
				OgcFid                 int     `json:"ogc_fid"`
				Geoid                  string  `json:"geoid"`
				Parcelnumb             string  `json:"parcelnumb"`
				ParcelnumbNoFormatting string  `json:"parcelnumb_no_formatting"`
				AltParcelnumb1         string  `json:"alt_parcelnumb1"`
				Owner                  string  `json:"owner"`
				Mailadd                string  `json:"mailadd"`
				MailCity               string  `json:"mail_city"`
				MailState              string  `json:"mail_state2"`
				MailZip                string  `json:"mail_zip"`
				County                 string  `json:"county"`
				AddressSource          string  `json:"address_source"`
				Legaldesc              string  `json:"legaldesc"`
				Lat                    string  `json:"lat"`
				Lon                    string  `json:"lon"`
				CensusTract            string  `json:"census_tract"`
				CensusBlock            string  `json:"census_block"`
				CensusBlockgroup       string  `json:"census_blockgroup"`
				LlLastRefresh          string  `json:"ll_last_refresh"`
				Gisacre                float64 `json:"gisacre"`
				LlGisacre              float64 `json:"ll_gisacre"`
				LlGissqft              int     `json:"ll_gissqft"`
				Path                   string  `json:"path"`
				Txaccttype             string  `json:"txaccttype"`
				NeighborhoodCode       string  `json:"neighborhood_code"`
			} `json:"fields"`
			Context struct {
				Headline string `json:"headline"`
				Name     string `json:"name"`
				Path     string `json:"path"`
				Active   bool   `json:"active"`
			} `json:"context"`
			LlUuid string  `json:"ll_uuid"`
			Score  float64 `json:"score"`
		} `json:"properties"`
	} `json:"results"`
}

// ParcelGeometryPolygon are obtained from the regrid.com
func ParcelGeometryPolygon(ctx context.Context, lat, lng float64) ([]schema.Point, error) {
	geometry, err := parcelGeometry(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	return parcelGeometryPoints(geometry), nil
}

func parcelGeometry(ctx context.Context, lat, lng float64) (*Geometry, error) {
	api, err := apiaccess.GetKey(ctx, "ReGrid")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/search.json?token=%s&lat=%f&lon=%f&strict=1&limit=1", api.URL, api.Key, lat, lng)
	b, err := http.Get(url)
	if err != nil {
		return nil, log.Wrap(err, "error on api call")
	}

	var res NearGridResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, log.Wrap(err, "failed to Unmarshal")
	}

	if len(res.Results) == 0 {
		return nil, errors.New(fmt.Sprintf("%s returned nothing for parcel search: lat=%f lon=%f", api.URL, lat, lng))
	}

	geometry := res.Results[0].Geometry
	return &geometry, nil
}

func parcelGeometryPoints(geo *Geometry) []schema.Point {
	if geo == nil {
		return nil
	}

	var bounds []schema.Point
	for _, v := range geo.Coordinates[0] {
		bounds = append(bounds, schema.Point{
			Lng: v[0],
			Lat: v[1],
		})
	}

	return bounds
}

func rectangle(ne, sw schema.Point) []schema.Point {
	// (𝑥1,𝑦1),(𝑥1,𝑦2),(𝑥2,𝑦2),(𝑥2,𝑦1).
	return []schema.Point{
		{Lat: ne.Lat, Lng: ne.Lng},
		{Lat: ne.Lat, Lng: sw.Lng},
		{Lat: sw.Lat, Lng: sw.Lng},
		{Lat: sw.Lat, Lng: ne.Lng},
	}
}
