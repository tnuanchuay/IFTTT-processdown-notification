# IFTTT-processdown-notification
Notify when your important process is down via IFTTT with many services.
### Require
- Golang

### Available
- Windows

### Installation
1. create event in https://ifttt.com using maker service
2. `$ go get github.com/tspn/IFTTT-processdown-notification`
3. copy settings.json.example to settings.json and setup 
4. run `$ IFTTT-processdown-notification.exe`

### Setting & Example
````json
{
  "ifttt-maker-key" : "<YOUR-MAKER-TOKEN>",
  "interval-time" : 1000,
  "processes-watcher" : [
    {
      "name":"<<Process 1 Name>>",
      "event":["When Process Down what even will trigged"]
    },
    {
      "name":"nginx.exe",
      "event":["facebook"]
    }
  ]
}
````
