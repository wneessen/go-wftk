package geoip

import (
	"fmt"
	"github.com/oschwald/maxminddb-golang"
	"net"
)

// GeoIP structs
type LookupCountry struct {
	GeoNameId int               `maxminddb:"geoname_id"`
	Names     map[string]string `maxminddb:"names"`
	IsoCode   string            `maxminddb:"iso_code"`
}

type LookupContinent struct {
	GeoNameId int               `maxminddb:"geoname_id"`
	Names     map[string]string `maxminddb:"names"`
	Code      string            `maxminddb:"code"`
}

type LookupLocation struct {
	AccuracyRadius int     `maxminddb:"accuracy_radius"`
	Latitude       float64 `maxminddb:"latitude"`
	Longitude      float64 `maxminddb:"longitude"`
	MetroCode      int     `maxminddb:"metro_code"`
	TimeZone       string  `maxminddb:"time_zone"`
}

type LookupCity struct {
	GeoNameId int               `maxminddb:"geoname_id"`
	Names     map[string]string `maxminddb:"names"`
}

type LookupPostal struct {
	Code string `maxminddb:"code"`
}

type IpLookup struct {
	City         LookupCity      `maxminddb:"city"`
	Continent    LookupContinent `maxminddb:"continent"`
	Country      LookupCountry   `maxminddb:"country"`
	Location     LookupLocation  `maxminddb:"location"`
	Postal       LookupPostal    `maxminddb:"postal"`
	RegCountry   LookupCountry   `maxminddb:"registered_country"`
	Subdivisions []LookupCountry `maxminddb:"subdivisions"`
}

// Return structs
type GeoRespCity struct {
	Name    string
	ZipCode string
}

type GeoRespContinent struct {
	Code string
	Name string
}

type GeoRespCountry struct {
	IsoCode string
	Name    string
}

type GeoRespLocation struct {
	Latitude  float64
	Longitude float64
	TimeZone  string
}

type GeoResponeObj struct {
	City         GeoRespCity
	Continent    GeoRespContinent
	Country      GeoRespCountry
	Location     GeoRespLocation
	RegCountry   GeoRespCountry
	Subdivisions []GeoRespCountry
}

// Return GeoIP information for a given IP address
// Return language is english (use GeoIpWithLang() for localized version)
func GeoIp(i string, f string) (GeoResponeObj, error) {
	return GeoIpWithLang(i, f, "en")
}

// Return localized GeoIP information for a given IP address
func GeoIpWithLang(i string, f string, l string) (GeoResponeObj, error) {
	// Open GeoIP db
	geoDb, err := maxminddb.Open(f)
	if err != nil {
		return GeoResponeObj{}, err
	}
	// Close GeoIP db
	defer func() {
		if err := geoDb.Close(); err != nil {
			fmt.Printf("Error while closing GeoIP database: %v", err)
		}
	}()

	// Loopup IP data
	ipAddr := net.ParseIP(i)
	lookupData := IpLookup{}
	if err := geoDb.Lookup(ipAddr, &lookupData); err != nil {
		return GeoResponeObj{}, err
	}

	// Prepare and return repsonse object
	geoResp := GeoResponeObj{
		City: GeoRespCity{
			Name:    lookupData.City.Names[l],
			ZipCode: lookupData.Postal.Code,
		},
		Continent: GeoRespContinent{
			Code: lookupData.Continent.Code,
			Name: lookupData.Continent.Names[l],
		},
		Country: GeoRespCountry{
			IsoCode: lookupData.Country.IsoCode,
			Name:    lookupData.Country.Names[l],
		},
		Location: GeoRespLocation{
			Latitude:  lookupData.Location.Latitude,
			Longitude: lookupData.Location.Longitude,
			TimeZone:  lookupData.Location.TimeZone,
		},
		RegCountry: GeoRespCountry{
			IsoCode: lookupData.RegCountry.IsoCode,
			Name:    lookupData.RegCountry.Names[l],
		},
	}
	for _, subDiv := range lookupData.Subdivisions {
		geoResp.Subdivisions = append(geoResp.Subdivisions,
			GeoRespCountry{
				Name:    subDiv.Names[l],
				IsoCode: subDiv.IsoCode,
			},
		)
	}
	return geoResp, nil
}
