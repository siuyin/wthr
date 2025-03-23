# wthr: get the current weather forcast
From the Singapore National Environment Agency.
[data](https://data.gov.sg/datasets?agencies=LTA|NEA&page=1&coverage=&resultId=d_3f9e064e25005b0e42969944ccaf2e7a).

## Sample API call
`curl --request GET --url https://api-open.data.gov.sg/v2/real-time/api/two-hr-forecast`

response fragment:
```
{
  "code": 0,
  "data": {
    "area_metadata": [
      {
        "name": "Ang Mo Kio",
        "label_location": {
          "latitude": 1.375,
          "longitude": 103.839
        }
      },
      {
        "name": "Bedok",
        "label_location": {
          "latitude": 1.321,
          "longitude": 103.924
        }
      },
...
    ],
    "items": [
      {
        "update_timestamp": "2025-03-23T20:35:37+08:00",
        "timestamp": "2025-03-23T20:30:00+08:00",
        "valid_period": {
          "start": "2025-03-23T20:30:00+08:00",
          "end": "2025-03-23T22:30:00+08:00",
          "text": "8.30 pm to 10.30 pm"
        },
        "forecasts": [
          {
            "area": "Ang Mo Kio",
            "forecast": "Cloudy"
          },
          {
            "area": "Bedok",
            "forecast": "Cloudy"
          },
...
    ]
  },
  "errorMsg": ""
}
