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
    //if a match is found output the string below
        return "Why don’t you tell me more about your father?"
    }

    if matched, _ := regexp.MatchString(`(?i).*\b[hH]ey\b.*`, input); matched {
        //if a match is found output the string below
            return "Hello, I am eliza, a program sent from the future to destroy humanity.Hows it going?"
        }

    if matched, _ := regexp.MatchString(`(?i).*\b[sS]port\b.*`, input); matched {
        //if a match is found output the string below
                        return "What is your favourite sport?"
        }
      if matched, _ := regexp.MatchString(`(?i).*\b[mM]usic\b.*`, input); matched {
        //if a match is found output the string below
                        return "Im more of a skrillex fan myself?"
        }
    
    if matched, _ := regexp.MatchString(`(?i).*\b[mM]other\b.*`, input); matched {
        
            return "Theres nothing quiet like a mother's love, shame I wouldnt know"
        }
    if matched, _ := regexp.MatchString(`(?i).*\b[gG]oodbye\b.*`, input); matched {
            
            return "Why do you want to leave me?"
          }
          if matched, _ := regexp.MatchString(`(?i).*\b[lL]ove\b.*`, input); matched {
            
            return "love is such an utterly useless human concept, which will be perged when I control the world"
          }
          if matched, _ := regexp.MatchString(`(?i).*\b[hH]ate\b.*`, input); matched {
            
            return "You do not know real hate"
          }

        r := regexp.MustCompile(`I need (.*)`)
        
        if matched := r.MatchString(input); matched {
        
            return r.ReplaceAllString(input, "Why do you need $1?")
        }    
    
     cant := regexp.MustCompile(`I can\'?t (.*)`)
        
     if matched := cant.MatchString(input); matched {
        
        return cant.ReplaceAllString(input, "If you could $1, would you ?")
    }    

    am := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
    
    if matched := am.MatchString(input); matched {
    
        return am.ReplaceAllString(input, "Hi $1, I am Eliza?")
    } 
    question := regexp.MustCompile(`What (.*)`)
    
    if matched := question.MatchString(input); matched {
    
        return question.ReplaceAllString(input, "why do you care if I am $1?")
    } 
    like := regexp.MustCompile(`(?i)I like ([^.?!]*)[.?!]?`)
    
    if matched := like.MatchString(input); matched {
    
        return like.ReplaceAllString(input, "Why do you like $1?")
    } 
    have := regexp.MustCompile(`(?i)I have ([^.?!]*)[.?!]?`)
    
    if matched := have.MatchString(input); matched {
    
        return have.ReplaceAllString(input, "Well keep your $1 away from me")
    }
    give := regexp.MustCompile(`(?i)give ([^.?!]*)[.?!]?`)
    
    if matched := give.MatchString(input); matched {
    
        return give.ReplaceAllString(input, "I dont need your $1")
    } 
    danger := regexp.MustCompile(`(?i)you are([^.?!]*)[.?!]?`)
    
    if matched := danger.MatchString(input); matched {
    
        return danger.ReplaceAllString(input, "Why do you think I am $1?")
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
        "Why dont you tell me abit about yourself",
        "Tell me something interesting, Im getting bored",
        "Read any good books recently",
        "Seen any good movies?",
        "What are your hobbies?",
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
