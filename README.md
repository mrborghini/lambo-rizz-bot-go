# Lambo-Rizz-Bot-Go
Just a twitch bot that tells you your Rizz levels

## Download and run binaries

[Get the binaries](https://github.com/lamborghinigamer1/lambo-rizz-bot-go/releases)

Change the config.json with your info.

[Generate a twitch OAuth token for bots (access token)](https://twitchtokengenerator.com/)

Nickname: is just the name of your bot. So the twitch.tv/name

OAuth: is the access token so this program can communicate between Twitch servers

A twitch token will look something like this: ```u1funyo1lo1ie12ponxa4soqp1332o```

Channels: Just add as many channels as you like. Just put all channels between double quotes and seperate them with commas

## How to compile and run?

[Install Go](https://go.dev/dl/)

Open a terminal

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

Congratulations you should now have a lambo-rizz-bot-go executable. If you're Linux just run the application with: ```./lambo-rizz-bot-go``` On windows ```.\lambo-rizz-bot-go.exe``` or double click it

## Troubleshooting

On linux if it doesn't run please run ```chmod +x lambo-rizz-bot-go```

Make 100% sure you have the generated an access token with reading chat privileges and send messages
