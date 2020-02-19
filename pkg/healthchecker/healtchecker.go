package healthchecker

import (
	"github.com/yanzay/tbot/v2"
	"net/http"
	"os"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	link := r.FormValue("link")

	if link == "" {
		http.Error(w, `{"success": false,"message":"Link parameter not is present"}`, http.StatusUnprocessableEntity)
		return
	}

	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := c.Head(link)

	if err != nil {
		http.Error(w, `{"success": false,"message":"Unable to reach `+link+`"}`, http.StatusNotFound)
		return
	}

	if response.StatusCode != 200 {
		bot := tbot.NewClient(os.Getenv("TELEGRAM_TOKEN"), http.DefaultClient, "https://api.telegram.org")
		_, botErr := bot.SendMessage(os.Getenv("CHAT_ID"), "❗Failure: "+link, tbot.OptParseModeMarkdown)
		if botErr != nil {
			http.Error(w, `{"success": false,"message":"Unable to send message to telegram bot"}`, http.StatusNotFound)
			return
		}

		http.Error(w, `{"success": false,"message":"Unable to reach `+link+`"}`, http.StatusNotFound)
		return
	}

	w.Write([]byte(`{"success": true,"message":"` + link + ` is responding with status 200"}`))
}
