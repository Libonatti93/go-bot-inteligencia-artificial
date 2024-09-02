package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/go-resty/resty/v2"
)

const openAIAPIKey = "YOUR_OPENAI_API_KEY" // Substitua pela sua chave de API da OpenAI

// Função para enviar a solicitação para a API do OpenAI
func getResponseFromOpenAI(prompt string) (string, error) {
    client := resty.New()

    // Define o corpo da solicitação
    requestBody := map[string]interface{}{
        "model": "gpt-3.5-turbo", // Modelo GPT-3.5 ou GPT-4
        "messages": []map[string]string{
            {"role": "user", "content": prompt},
        },
        "max_tokens": 100,
    }

    // Envia a solicitação para a API da OpenAI
    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        SetHeader("Authorization", "Bearer "+openAIAPIKey).
        SetBody(requestBody).
        Post("https://api.openai.com/v1/chat/completions")

    if err != nil {
        return "", err
    }

    var result map[string]interface{}
    err = resp.Unmarshal(&result)
    if err != nil {
        return "", err
    }

    // Extrai a resposta do bot
    choices := result["choices"].([]interface{})
    firstChoice := choices[0].(map[string]interface{})
    message := firstChoice["message"].(map[string]interface{})
    content := message["content"].(string)

    return content, nil
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("AI Bot iniciado. Digite 'exit' para sair.")

    for {
        fmt.Print("Você: ")
        userInput, _ := reader.ReadString('\n')
        userInput = strings.TrimSpace(userInput)

        if strings.ToLower(userInput) == "exit" {
            fmt.Println("Encerrando o bot. Até mais!")
            break
        }

        response, err := getResponseFromOpenAI(userInput)
        if err != nil {
            log.Fatalf("Erro ao obter resposta da OpenAI: %v", err)
        }

        fmt.Printf("Bot: %s\n", response)
    }
}
