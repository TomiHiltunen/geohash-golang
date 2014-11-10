package geohash

import "testing"

type geohashTest struct {
	input  string
	output *BoundingBox
}

func TestDecode(t *testing.T) {
	var tests = []geohashTest{
		geohashTest{
			"d",
			&BoundingBox{
				LatLng{0, -90},
				LatLng{45, -45},
				LatLng{22.5, -67.5},
			},
		},
		geohashTest{
			"dr",
			&BoundingBox{
				LatLng{39.375, -78.75},
				LatLng{45, -67.5},
				LatLng{42.1875, -73.125},
			},
		},
		geohashTest{
			"dr1",
			&BoundingBox{
				LatLng{39.375, -77.34375},
				LatLng{40.78125, -75.9375},
				LatLng{40.078125, -76.640625},
			},
		},
		geohashTest{
			"dr12",
			&BoundingBox{
				LatLng{39.375, -76.9921875},
				LatLng{39.55078125, -76.640625},
				LatLng{39.462890625, -76.81640625},
			},
		},
	}

	for _, test := range tests {
		box := Decode(test.input)
		if !equalBoundingBoxes(test.output, box) {
			t.Errorf("expected bounding box %v, got %v", test.output, box)
		}
	}
}

func equalBoundingBoxes(b1, b2 *BoundingBox) bool {
	return b1.ne == b2.ne &&
		b1.sw == b2.sw &&
		b1.center == b1.center
}

type encodeTest struct {
	latlng  LatLng
	geohash string
}

func TestEncode(t *testing.T) {
	var tests = []encodeTest{
		encodeTest{
			LatLng{39.55078125, -76.640625},
			"dr12zzzzzzzz",
		},
		encodeTest{
			LatLng{39.5507, -76.6406},
			"dr18bpbp88fe",
		},
		encodeTest{
			LatLng{39.55, -76.64},
			"dr18bpb7qw65",
		},
		encodeTest{
			LatLng{39, -76},
			"dqcvyedrrwut",
		},
	}

	for _, test := range tests {
		geohash := Encode(test.latlng.lat, test.latlng.lng)
		if test.geohash != geohash {
			t.Errorf("expectd %s, got %s", test.geohash, geohash)
		}
	}

	for prec := range []int{3, 4, 5, 6, 7, 8} {
		for _, test := range tests {
			geohash := EncodeWithPrecision(test.latlng.lat, test.latlng.lng, prec)
			if len(geohash) != prec {
				t.Errorf("expected len %d, got %d", prec, len(geohash))
			}
			if test.geohash[0:prec] != geohash {
				t.Errorf("expectd %s, got %s", test.geohash, geohash)
			}
		}
	}
}
