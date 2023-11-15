# Lambo-Rizz-Bot-Go
Just a twitch bot that tells you your Rizz levels

## How to compile and run?

Clone the repository: 
```https://github.com/lamborghinigamer1/lambo-rizz-bot-go.git```

Change directory into the repository ```cd lambo-rizz-bot-go```

Run the command:
```go get lambo-rizz-bot-go/api```


Generate access tokens:
[Generate a twitch OAuth token for bots (access token)](https://twitchtokengenerator.com/)



Create a file named config.json

```json
{
    "Nickname": "lamborizzbot",
    "OAuth": "YourAccessToken",
    "Channels": ["mrborghini_", "lamborizzbot"]
}
```


Then compile the executable:
```go build```

Congratulations you should now have a lambo-rizz-bot-go executable. If you're Linux just run the application with: ```./lambo-rizz-bot-go```


