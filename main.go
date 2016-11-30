package main

const(
	FILENAME = "settings.json"
)

func main(){
	var setting Settings
	setting.ReadSettings(FILENAME)
}