# http-health-checker

A simple http health checker that sends a notification to your telegram when your site is not reachable.

<p align="center">
  <img width="400px" src="https://user-images.githubusercontent.com/4256471/73611224-7824c580-45be-11ea-98f9-d7478e27118b.png">
</p>

## Getting started

### Telegram bot

You need to create a telegram bot to receive the notification for not reachable links.

1. Open your telegram and talk to `@BotFather`
2. Create your bot and save your token (it will be used as `TELEGRAM_TOKEN`)
3. Start a simple conversation with your new created bot (just one word is enought, this will give permissions to bot talk back to you)
4. Talk to `@JsonDumpBot` and get your user id (it will be used as `CHAT_ID`)

### Running

1. Set the `CHAT_ID` and `TELEGRAM_TOKEN` environment variables
2. Run the app and call the /?link=<your-website-link-here>

## Tip

You can deploy it to google functions and set up the cloud scheduler to call it once per minute.
