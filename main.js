const propertyId = process.env.PROPERTY_ID

// Imports the Google Analytics Data API client library.
const { BetaAnalyticsDataClient } = require('@google-analytics/data');

// Using a default constructor instructs the client to use the credentials
// specified in GOOGLE_APPLICATION_CREDENTIALS environment variable.
const analyticsDataClient = new BetaAnalyticsDataClient();

// Runs a simple report.
async function runReport () {
    const [response] = await analyticsDataClient.runReport({
        property: `properties/${propertyId}`,
        dateRanges: [
            {
                startDate: '2020-03-31',
                endDate: 'today',
            },
        ],
        dimensions: [
            // {
            //     name: 'city',
            // },
            // {
            //     name: 'pageTitle'
            // },
            {
                name: 'pagePath'
                // name: 'pagePathPlusQueryString'
            },
            // {
            //     name: 'eventName'
            // }
            // {
            //     name: 'hostName'
            // },
            // {
            //     name: 'fullPageUrl'
            // }
        ],
        metrics: [
            // {
            //     name: 'activeUsers',
            // },
            // {
            //     name: 'eventCount'
            // },
            // {
            //     name: 'eventCountPerUser'
            // },
            {
                name: 'screenPageViews'
            },
            // {
            //     name: 'transactions',
            // },
            {
                name: 'totalUsers',
                invisible: false
            },
            // {
            //     name: 'transactions'
            // },
        ],
        offset: 0,
        limit: 10
        // keepEmptyRows: true,
        // metricFilter: {
        //     "filter": {
        //         object (Filter)
        //       }
        // }
    });

    // console.log('Report result:');

    // response.rows.forEach(row => {
    //     console.log(row.dimensionValues[0], row.metricValues[0]);
    // });


    const util = require('util')

    console.log("===========================")
    console.log(util.inspect(response, false, null, true /* enable colors */))
    console.log("===========================")
}

runReport();
