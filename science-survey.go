package science

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const VERSION = "0.1"

const BASE = "https://docs.google.com/forms/"
const SURVEY_POLITICAL = "d/e/1FAIpQLSfIc4B4XN5J06S3sT4pd-Ylehs1Ki75u9kYF3VlrHZaXYazvA/viewform"
const SURVEY_RELIGIOUS = "d/16a06KXHYdNjY9POJQ2hwL56t7rR4IYr0ce6fGUbuT6Q/viewform"

const PAGE = `
<html>
  <head>
  	<!-- %s -->
    <style>
        .intro {
            padding-top:10%%;
            padding-left:30%%;
            padding-right:30%%;
            font-size:larger;
        }
        .button {
          font-size: larger;
          text-decoration: none;
          background-color: #0095ff;
          color: #333333;
          padding: 2px 6px 2px 6px;
          border-top: 1px solid #CCCCCC;
          border-right: 1px solid #333333;
          border-bottom: 1px solid #333333;
          border-left: 1px solid #CCCCCC;
          box-shadow: inset 0 1px 0 #66bfff;
          border-radius: 10px;
        }
    </style>
  </head>
  <body>
    <table>
      <tr>
        <td class="intro">
          <center>
            Hi my name is Eli Farrer, and thank you for taking this survey. It will take about 1
            minute and it will really help me with my science fair project. Please do not discuss any part of the
            survey with anyone else to keep the survey unbiased. If you can please ask others to
            also take the survey, I would appreciate it. The more people I get the better my data will be.
            </center>
        </td>
      </tr>
      <tr>
        <td style="padding-top:5%%">
          <center>
            <a href="/survey" class="button">Begin Survey</a>
          </center>
        </td>
      </tr>
    </table>
  </body>
</html>
`

func init() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/survey", surveyHandler)
}

func surveyHandler(w http.ResponseWriter, r *http.Request) {
	val := rand.Int() % 2
	if val == 0 {
		http.Redirect(w, r, BASE+SURVEY_POLITICAL, http.StatusFound)
	} else if val == 1 {
		http.Redirect(w, r, BASE+SURVEY_RELIGIOUS, http.StatusFound)
	} else {
		panic("Random is broken")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, PAGE, VERSION)
}
