# wapp_start_page_links

This is a simple web application written in Go using the Fiber web framework. The purpose of this application is to provide a customizable start page with links to your favorite websites.



## Getting Started

1. Clone the repository to your local machine.
2. Make sure you have [Go](https://golang.org/) installed.
3. Create a `.env` file in the project root with the following content:
   ```plaintext
   PORT=8585
   HOST=0.0.0.0
   CONFIG_FILE=./config.yaml
   ```
   Adjust the values as needed.
4. Run the application:
   ```bash
   go run main.go
   ```

## Usage

- Visit [http://localhost:8585](http://localhost:8585) in your web browser to access the start page.
- Click on the "Edit List" button to manage your website links.
- Customize the settings by navigating to the "Settings" page.

## Dependencies

- [Fiber](https://github.com/gofiber/fiber)
- [Godotenv](https://github.com/joho/godotenv)
- [Embed](https://pkg.go.dev/embed)

