package nearmap

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestParcelGeometryPolygon(t *testing.T) {
	lat := 40.60989724
	lng := -73.95637593
	rect, err := ParcelGeometryPolygon(context.Background(), lat, lng)
	if err != nil {
		t.Error(err)
	}

	var sb strings.Builder
	for _, p := range rect {
		sb.WriteString(fmt.Sprintf("%f,%f,", p.Lng, p.Lat))
	}

	sb.WriteString(fmt.Sprintf("%f,%f", rect[0].Lng, rect[0].Lat))
	t.Log(sb.String())
}
