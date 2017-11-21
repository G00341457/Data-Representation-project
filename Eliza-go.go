package main

import(
    import "regexp"
    
    
    struct Reponse{
    re =regexp.Rexexp
    answer [String]
    }
    
    //sample input :="Hello my name is blank"
    
    
    
    
    
    func subWords(topic String) string{
    
    reflections := map[String]string{
        "you": "me",
        "me": "you",
    }
    
       words := strings.split(topic, " ")
       
       for index, word := range words{
         val, ok := reflections[word]
         if ok{//the value was in the map
           
         }
         
         if val, ok := relections[word]; ok{//the
         //if we are here the word is in the map
         words[index] = val//replace you with me
       }
       
       return strings.Join(words, " ")
    
    }
    
    
    patternStr:="name is (.*)\s"
    answer := "hi %s, its nice to meet you"
         //hello, %s
         // hi there, %s
    
    
    re := rexep.MustCompile(patternstr)//fails early to find your mistake
    fmt.print()
    
    if. re.MatchString(userInput){
       fmt.println("there was a match")
       match:= re.matchSubStringMatch(userInput)
       fmt.println(match[0])
       sub := match[1]//not always going to be a name
       //pick answer from text file
       //if %s is in the input
       finalAnswer := fmt.Sprintf(answer,subject)
       fmt.println(FinalAnswer)
       }else{
       fmt.print("errorE)
       }
       
       
       //create a text file 
       //my name is (.*)
       //Hi there %s, hi %s