# google-analytics-101

[![Netlify Status](https://api.netlify.com/api/v1/badges/56ec4d05-5da2-460c-b2f2-4e82aa63596e/deploy-status)](https://app.netlify.com/sites/laughing-visvesvaraya-f032b7/deploys)

<!-- https://www.googleapis.com/analytics/v3/data/ga?ids=ga:12345&start-date=2008-10-01&end-date=2008-10-31&metrics=ga:sessions,ga:bounces -->

## REF

https://machiel.me/post/using-google-analytics-api-with-golang/

## Make credential for google

https://cloud.google.com/docs/authentication/production#cloud-console

## API Credential

https://console.cloud.google.com/apis/credentials?authuser=2&project=sale-page-325603

## ~~Golang UA~~ << Obsolete

### Step UA

1. Create google account
2. Create Organize of google analytic
3. Create googla analytics
4. [Enable Universal Analytics](https://support.google.com/analytics/answer/10269537?hl=en)
5. [Create Data Streams](https://support.google.com/analytics/answer/9304153?hl=en)
   1. [Other one](https://www.datadrivenu.com/understanding-data-streams-google-analytics-4/)

## TODO

- [migrate UA to GA4](https://developers.google.com/analytics/devguides/reporting/data/v1/migration-guide?authuser=2)

## JS (GA4)

Run report support only start GA4 to before current date

```sh
export GOOGLE_APPLICATION_CREDENTIALS="./key.json"
export PROPERTY_ID="xxxxxxxxx"
```

- [doc](https://googleapis.dev/nodejs/analytics-data/latest/index.html#installing-the-client-library)
- [Admin API](https://developers.google.com/analytics/devguides/config/admin/v1)
- [Manage API](https://developers.google.com/analytics/devguides/config/mgmt/v3/mgmtReference/management/webproperties/get)

## GTM

1. Add Trigger Click - All Elements
2. select `Some Clicks`
3. add condition at `Fire this trigger` with `Click ID + contains + your button ID`
4. Add tag on trigger
   1. add parameter with `page_path` + `{{Page Path}}`
5. Add Variable `Click ID`

> Don't forget publish GTM
