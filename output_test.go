package twilio

var output = map[string]string{
	// Message: sample of the MMS response
	"message": `
{
	"account_sid": "AC5ef8732a3c49700934481addd5ce1659",
	"api_version": "2010-04-01",
	"body": "Jenny please?! I love you <3",
	"num_segments": "1",
	"num_media": "1",
	"date_created": "Wed, 18 Aug 2010 20:01:40 +0000",
	"date_sent": null,
	"date_updated": "Wed, 18 Aug 2010 20:01:40 +0000",
	"direction": "outbound-api",
	"from": "+14158141829",
	"price": null,
	"sid": "MM90c6fc909d8504d45ecdb3a3d5b3556e",
	"status": "queued",
	"to": "+15558675309",
	"uri": "/2010-04-01/Accounts/AC5ef8732a3c49700934481addd5ce1659/Messages/MM90c6fc909d8504d45ecdb3a3d5b3556e.json"
}`,

	// Message: sample of the error response
	"message_error": `
{
	"status": 400,
	"message": "A 'From' phone number is required.",
	"code": 21603,
	"more_info": "http:\/\/www.twilio.com\/docs\/errors\/21603"
}`,

	// Message: List of the messages response
	"message_list": `
{
	"start": 0,
	"total": 261,
	"num_pages": 6,
	"page": 0,
	"page_size": 50,
	"end": 49,
	"uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json",
	"first_page_uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json?Page=0&PageSize=50",
	"last_page_uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json?Page=5&PageSize=50",
	"next_page_uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json?Page=1&PageSize=50",
	"previous_page_uri": null,
	"messages": [
		{
			"account_sid": "ACc51860f991f74032b73fdc58841d39fa",
			"api_version": "2010-04-01",
			"body": "Hey Jenny why aren't you returning my calls?",
			"num_segments": "1",
			"num_media": "0",
			"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
			"date_sent": "Mon, 16 Aug 2010 03:45:03 +0000",
			"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
			"direction": "outbound-api",
			"from": "+14158141829",
			"price": "-0.02000",
			"sid": "SM800f449d0399ed014aae2bcc0cc2f2ec",
			"status": "sent",
			"to": "+15558675309",
			"uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec.json"
		}
	]
}`,
}
