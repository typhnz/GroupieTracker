package main

//localhost:8080/ enter this to connect the web site
import (
	"HangmanWeb"
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

var Data HangmanWeb.Dataform // We create Data with type Dataform
var Online bool              //We create this global variable who say if the player is online or not
var wg sync.WaitGroup        //We use this value for create a secondary chanel for AddWord

func main() {
	//We check if the player is online
	resp, _ := http.Get("https://github.com/GurvanN22/HangMan")
	if resp != nil {
		Online = true
	}
	Data.WaitingScreen = true //We set the Waiting Screen on true
	Inisialistion()
}

func Inisialistion() {
	Port := "8080"                                          //We choose port 8080
	fmt.Println("The serveur start on port " + Port + " 🔥") //We print this when the server is online
	styles := http.FileServer(http.Dir("template/css"))
	http.Handle("/styles/", http.StripPrefix("/styles", styles)) //We link the css with http.Handle
	http.HandleFunc("/", MainPage)                               //We create the main page , the only function who use a template
	http.HandleFunc("/letterInput", LetterInput)                 //We create letter input , we run this function when a letter is click on the main page
	http.HandleFunc("/buttonLevel", ButtonLevel)                 //We create this function for the choose of the difficulty and the initialisation of starts values
	http.HandleFunc("/Save", Save)
	http.HandleFunc("/Load", Load)
	http.HandleFunc("/AddWord", AddWord)
	http.HandleFunc("/Reset", Reset)
	http.ListenAndServe(":"+Port, nil) //We start the server
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/index.html"))             //We link the template and the html file
	Data.DisplayEmpty = HangmanWeb.SpaceMaker(HangmanWeb.ArrayStrToStr(Data.Empty)) // We use SpaceMaker for the text Display, (SpaceMaker create a space between every letters for more accecibility )
	tmpl.Execute(w, Data)                                                           //We execute the template and put the Data in
}

func ButtonLevel(w http.ResponseWriter, r *http.Request) {
	NameDB := r.FormValue("buttonLevel")
	Data = HangmanWeb.Dataform{}
	Data = HangmanWeb.InitialiatonOfStartsValues(NameDB, Online)
	MainPage(w, r)
}

func LetterInput(w http.ResponseWriter, r *http.Request) {
	if Data.GameStart == true { //We check if the game start
		Data.LastInput = r.FormValue("letter")                                       //We get the clicked letter
		Data = HangmanWeb.ValueTest(Data)                                            //We use ValueTest for make all test on the letter and update the Data
		Data.Button = HangmanWeb.SuppLetterInDictionary(Data.LastInput, Data.Button) //We delete the clicked Letter , so the button disapear
		if Data.LenghtWord == Data.TrueValues && Data.LenghtWord != 0 {              //We check if the player win the game
			Data.ForDefinition = HangmanWeb.ArrayStrToStr(Data.Word) //We make a the word without space for the definition
			Data.GameStart = false                                   //We stop the game
			Data.Win = true                                          //We display the Win page
			MainPage(w, r)
		} else if Data.Mistakes == 10 { //We check if the palyer loss
			Data.ForDefinition = HangmanWeb.ArrayStrToStr(Data.Word)
			Data.Lose = true
			Data.GameStart = false
			MainPage(w, r)
		} else {
			MainPage(w, r)
		}
	} else {
		MainPage(w, r)
	}

}

func Save(w http.ResponseWriter, r *http.Request) {
	if Data.GameStart == true { //We check if the player is in the game
		HangmanWeb.SaveProgration(Data) //We use SaveProgration in the save.go file
	}
	MainPage(w, r)
}

func Load(w http.ResponseWriter, r *http.Request) {
	Data = HangmanWeb.ExtractSave(Data)
	Data.OnlineConnection = Online
	Data.WaitingScreen = false
	MainPage(w, r)
}

func Reset(w http.ResponseWriter, r *http.Request) {
	Data = HangmanWeb.Dataform{}
	Data.OnlineConnection = Online
	Data.WaitingScreen = true
	MainPage(w, r)
}

func AddWord(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("AddWord")
	wg.Add(1)                           //We create a secondary chanel
	go HangmanWeb.AddNewWord(word, &wg) //We run AddNewWord in this chanel because the function is slow
	wg.Wait()                           //We stop the chanel
	Reset(w, r)
}
