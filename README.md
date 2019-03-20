### My custom weather service

Some example API calls, or from where are we going to gather the weather information :)

API: `https://twcservice.mybluemix.net/rest-api/`

Search:
```
curl -X GET --header 'Accept: application/json' 'https://twcservice.mybluemix.net/api/weather/v3/location/search?query=Cacheuta&locationType=city&language=es-AR'

```

Godoy Cruz: https://twcservice.mybluemix.net/api/weather/v1/geocode/-32.903/-68.897/forecast/hourly/48hour.json?language=es-AR&units=m
```
curl -X GET --header 'Accept: application/json' 'https://twcservice.mybluemix.net/api/weather/v1/geocode/-32.903/-68.897/forecast/hourly/48hour.json?language=es-AR&units=m'
```

Potrerillos: ARXX0340:1:AR
Godoy Cruz: ARMA0036:1:AR
Cacheuta: ARXX0024:1:AR

Cacheuta 48 hours by hour: https://twcservice.mybluemix.net/api/weather/v1/location/ARXX0024%3A1%3AAR/forecast/hourly/48hour.json
Cacheuta 10 days by day: `curl -X GET --header 'Accept: application/json' 'https://twcservice.mybluemix.net/api/weather/v1/location/ARXX0024%3A1%3AAR/forecast/daily/10day.json?language=es-AR&units=m'`
```
curl -X GET --header 'Accept: application/json' 'https://twcservice.mybluemix.net/api/weather/v1/location/ARXX0024%3A1%3AAR/forecast/hourly/48hour.json'
```

Example response:
```
{
  "metadata": {
    "language": "en-US",
    "transaction_id": "1552260729515:-562530502",
    "version": "1",
    "location_id": "ARXX0024:1:AR",
    "units": "e",
    "expire_time_gmt": 1552261190,
    "status_code": 200
  },
  "forecasts": [
    {
      "class": "fod_short_range_hourly",
      "expire_time_gmt": 1552261190,
      "fcst_valid": 1552262400,
      "fcst_valid_local": "2019-03-10T21:00:00-0300",
      "num": 1,
      "day_ind": "N",
```

Example requests to our weatherapi:
```
curl -X POST \
  http://localhost:3000/locations/create \
  -H 'Authorization: Open sesame' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \  -d '{
    "location_id": "Cacheuta", \
    "user_id": "pepe@mail.com" \
}'


curl -X POST \
  http://localhost:3000/locations/search \
  -H 'Authorization: Open sesame' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Cacheuta",
}'

curl -X DELETE \
  http://localhost:3000/locations/1 \
  -H 'Authorization: Open sesame' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json'
```
