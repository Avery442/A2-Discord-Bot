# A2 Discord Bot

This bot is a recreation of the Orion Drift A2 bot that used to serve information about the current players online in each fleet.

## Requirements

Before running the bot, ensure you have the following environment variables set:

* `TOKEN`: Your Discord bot token.
* `A2_API_KEY`: API key for accessing fleet data.

## Setup

1. Clone this repository:

   ```bash
   git clone https://github.com/Avery442/A2-Discord-Bot.git
   cd A2-Discord-Bot
   ```

2. Create a `.env` file in the root directory with the following content:

   ```env
   TOKEN=your_discord_bot_token
   A2_API_KEY=your_a2_api_key
   ```

3. Install the required Go dependencies:

   ```bash
   go mod tidy
   ```

4. Run the bot:

   ```bash
   go run main.go
   ```

## Example Response

![Bot Response Example](assets/Howmanyspacemonke_example.png)

## Contributing

Feel free to fork the repository, submit issues, and create pull requests. Contributions are welcome!

