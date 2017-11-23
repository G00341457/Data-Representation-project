// forms.go
package main

import (
    "fmt"
    "html/template"
    "net/http"
    "math/rand"
    "regexp"
)

func ElizaResponse(input string) string {
    
    if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
    
        return "Why don’t you tell me more about your father?"
    }
    
    if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
        
            return "Why don’t you tell me more about your mother?"
        }

        r := regexp.MustCompile(`I need (.*)`)
        
        if matched := r.MatchString(input); matched {
        
            return r.ReplaceAllString(input, "Why do you need $1?")
        }    

    re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "Hi $1, I am Eliza?")
    }   
    
    
    
    answers := []string{
    
        "I’m not sure what you’re trying to say. Could you explain it to me?",    
        "How does that make you feel?",   
        "Why do you say that?",
        "Why dont you tell me about it?",
        "Can I help you with that?",
        "Very Interesting",
        "Tell me more",
        "Use adjectives",
    }
    
    return answers[rand.Intn(len(answers))]
    
}

func main() {
    tmpl := template.Must(template.ParseFiles("eliza.html"))
    tmpl2 := template.Must(template.ParseFiles("echomessage.html"))

    
    http.HandleFunc("/eliza.html", func(w http.ResponseWriter, r *http.Request) {
            
            tmpl.Execute(w, nil)
            
    })
    
    http.HandleFunc("/echomessage.html", func(w http.ResponseWriter, r *http.Request) {
            
        r.ParseForm()
        usermessage := r.Form.Get("message")
        fmt.Println(usermessage)
        
        response := ElizaResponse(usermessage)
        fmt.Println(response)
        
        tmpl2.Execute(w, struct {Message string}{response})
    })

    http.ListenAndServe(":8080", nil)
}
