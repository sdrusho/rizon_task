# rizon_task
feedback app with back end


Back end Details

Swagger is used for every api details documentation

Jwt token is generated and validate by desired api

Email service working with google smtp. if needed other smtp server can be used only credentials in .env file need to change 
no need to change in email service.

For sending mail pkg/templates is used

For user check userId is UUID and checked in back end so that one user can not request multiple time

Rate limiting is used on memory sliding window algorithm, it can be configured like per minute/hour how many request can be used (server.go)

Deep link is used for singup due to restriction for ip it will tested manually though real domain is not available.

Slack mock is used in a way only real client update can work everything , no need to change in service


Mobile app details

Deep link is used for sign in and jwt token is used for feedback post request

Onbording should be shown only once per user and implemented

Expo is used for navigation both restart and deep link

Bottom sheet is used for feedback

auth layout used for feedback and User provider used for sink for every feedback sheet

different generic components are used like Themed text, view,input, Review logo

constant are used for theme but not properly used

Re direct in IOS app-store and Google play store are also implemented from app













Supports Google App Engine
Google App Engine does not allow the use of raw net/http clients. Instead you need to use the urlfetch package. This library allows you to pass the HTTP client that it should use:

ctx := appengine.NewContext(r)
httpCl := urlfetch.Client(ctx)
slackCl := slack.New(token, slack.WithClient(httpCl))
Google API style library
All of the APIs in this library resemble that of google.golang.org/api, and is very predictable across all APIs.

Full support for context.Context
The API is designed from the ground up to integrate well with context.Context.

Mock Server
A server that mocks the Slack API calls is available, which should help you integrate your applications with Slack while you are still in development.

Note: many mock methods are still NOT properly implemented. Please see server/mockserver/mockserver.go and server/mockserver/response.go. PRs welcome!
