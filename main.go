package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type ThumbnailRequest struct {
    VideoURL   string `json:"video_url"`
    StartTime  string `json:"start_time"`
    Duration   string `json:"duration"`
    Size       string `json:"size"`
}

func generateThumbnail(apiToken, videoURL, outputFileName string) {
    // Include the output file name in the API URL
    apiUrl := fmt.Sprintf("https://api.apyhub.com/generate/video-thumbnail/url?output=%s", outputFileName)
    
    payload := ThumbnailRequest{
        VideoURL:  videoURL,
        StartTime: "0",
        Duration:  "2",
        Size:      "400x300",
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        fmt.Println("Error marshalling data:", err)
        return
    }

    req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("apy-token", apiToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request to the API:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response:", string(body))
}

func main() {
    apiToken := "APY0BOODK2plpXgxRjezmBOXqID51DGpFq8QnHJeBQrrzuIBc25UIglN93bbwvnkBWlUia1" // Replace with your APY token
    videoURL := "https://assets.apyhub.com/samples/sample.mp4"
	outputFileName := "test-sample.mp4" // Specify the desired output file name

    generateThumbnail(apiToken, videoURL, outputFileName)
}


