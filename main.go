package main

import (
	"context"
	"fmt"
	"os"

	"code.sajari.com/docconv/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	res, err := docconv.ConvertPath("YOUR_PPT_PATH")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(res)
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer client.Close()
	
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("Check for plagiarism by only analysing text,give a score out of 100 without additional text %s", res)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}