package pedant

import (
	"code.google.com/p/go.exp/inotify"
	"encoding/json"
	"io/ioutil"
	"log"
)

type State struct {
	json  []byte
	valid bool
}


func errHandler(err error) {

	if err != nil {
		log.Print(err)
	}

}

func validateJSON(jdata []byte) bool {
	var jmap map[string]interface{}
	return json.Unmarshal(jdata, &jmap) == nil
}

func TrackFile(filename string, status chan<- *State) {

	watcher, err := inotify.NewWatcher()
	errHandler(err)

	err = watcher.Watch(filename)
	errHandler(err)

	for {
		select {
		case ev := <-watcher.Event:
			if ev.Mask == 2 {
				log.Print("file changed")
				jsondata, err := ioutil.ReadFile(filename)
				errHandler(err)

				status <- &State{json: jsondata, valid: validateJSON(jsondata)}
			}
		}
	}

}
