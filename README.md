# Video Thumbnail Generator

This project provides a simple Go application that generates thumbnails for videos by interacting with an external API. It sends a request to the API with details about the video for which a thumbnail is needed, including the video URL, start time, duration, and size of the thumbnail.

## Prerequisites

Before running this program, ensure you have the following installed:
- Go (version 1.15 or later recommended)
- An API token from the service provider (in this case, `apyhub.com`)

## Steps to Run the Program

1. **Obtain an API Token**: First, you need to obtain an API token from `apyhub.com`. Replace the `apiToken` variable's value in the `main.go` file with your actual API token.

2. **Install Go**: If you haven't already, [install Go](https://golang.org/doc/install) on your machine.

3. **Configure the Program**: Open the `main.go` file and set the `videoURL` and `outputFileName` variables to your desired video URL and output file name, respectively.

4. **Run the Program**: Open a terminal or command prompt, navigate to the project directory, and run the following command: `go run main.go`

## Note

This program is for educational purposes and demonstrates basic API interaction using Go. Ensure you comply with the API provider's usage policies and guidelines.
