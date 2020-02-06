package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

const (
	ponyURL = "https://theponyapi.com/api/v1/pony/random"
	helpMsg = `Please choose any opsfun command:
/opsfun help - for help :trollface:
/opsfun magic
/opsfun this-is-fine
/opsfun burn-money
/opsfun this-is-not-fine
/opsfun everything-is-on-fire
/opsfun johnny
/opsfun crazy-cat
/opsfun pony
/opsfun princess
/opsfun honk
/opsfun meow
/opsfun oompa
`
)

func Handle(req handler.Request) (handler.Response, error) {

	var query *url.Values
	if len(req.Body) > 0 {
		q, err := url.ParseQuery(string(req.Body))
		if err != nil {
			return handler.Response{}, err
		}
		query = &q
	}

	command := query.Get("command")
	arg := query.Get("text")

	if command != "/opsfun" || len(arg) == 0 {
		return handler.Response{
			Body:       []byte(generateStandardSlashResponse(helpMsg, "ephemeral", "OpsFun", "")),
			StatusCode: http.StatusOK,
			Header: map[string][]string{
				"Content-Type": []string{"application/json"},
			},
		}, nil
	}

	message := processCommand(arg)

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
		Header: map[string][]string{
			"Content-Type": []string{"application/json"},
		},
	}, nil
}

func processCommand(arg string) string {
	command := strings.Split(arg, " ")

	switch command[0] {
	case "magic":
		gif := "in response to `/opsfun magic`\n\nhttps://thumbs.gfycat.com/AmbitiousUnselfishHake-size_restricted.gif"
		return generateStandardSlashResponse(gif, "in_channel", "Magic", "")
	case "this-is-fine":
		gif := "in response to `/opsfun this-is-fine`\n\nhttps://thumbs.gfycat.com/AngryGoldenIsabellinewheatear-size_restricted.gif"
		return generateStandardSlashResponse(gif, "in_channel", "This is Fine", "")
	case "burn-money":
		gif := "in response to `/opsfun burn-money`\n\nhttps://thumbs.gfycat.com/EllipticalEsteemedBackswimmer-size_restricted.gif"
		return generateStandardSlashResponse(gif, "in_channel", "Burn!", "")
	case "this-is-not-fine":
		gif := "in response to `/opsfun this-is-not-fine`\n\nhttps://storage.googleapis.com/this-is-fine-images/this_is_not_fine.png"
		return generateStandardSlashResponse(gif, "in_channel", "This is Not Fine", "")
	case "everything-is-on-fire":
		gif := "in response to `/opsfun everything-is-on-fire`\n\nhttps://media.giphy.com/media/13HgwGsXF0aiGY/giphy.gif"
		return generateStandardSlashResponse(gif, "in_channel", "Fire!", "")
	case "johnny":
		gif := "in response to `/opsfun johnny`\n\nhttps://media.giphy.com/media/RoajqIorBfSE/giphy.gif"
		return generateStandardSlashResponse(gif, "in_channel", "Here is Johnny!", "")
	case "waiting":
		gif := "in response to `/opsfun waiting`\n\nhttps://thumbs.gfycat.com/SkeletalIllustriousDonkey-size_restricted.gif"
		return generateStandardSlashResponse(gif, "in_channel", "zzzZZzz", "")
	case "oompa":
		gif := "in response to `/opsfun oompa`\n\nhttps://thumbs.gfycat.com/TallOrnateDuckbillcat-size_restricted.gif"
		return generateStandardSlashResponse(gif, "in_channel", "Oompa Loompa", "")
	case "hugops":
		gif := "in response to `/opsfun hugops`\n\nhttps://thumbs.gfycat.com/DishonestHideousBlacklab-size_restricted.gif\n#hugops"
		return generateStandardSlashResponse(gif, "in_channel", "Hug Ops", "")
	case "crazy-cat":
		gif := "in response to `/opsfun crazy-cat`\n\nhttps://thumbs.gfycat.com/FancyFastGarpike-size_restricted.gif"
		return generateStandardSlashResponse(gif, "in_channel", "Cat", "")
	case "pony":
		ponyImage, ponyThumb := getPony()
		if ponyImage == "" {
			ponyImage = "Couldn't find a pony matching that query."
			ponyThumb = "https://pones.theponyapi.com/file/ponies/6247559689404416/thumb_small"
		}
		gif := fmt.Sprintf("in response to `/opsfun %s`\n\n%s", command[0], ponyImage)
		return generateStandardSlashResponse(gif, "in_channel", "Pony", ponyThumb)
	case "princess":
		princessImage, princessThumb := getImage("princess")
		if princessImage == "" {
			princessImage = "Couldn't find a princess matching that query."
			princessThumb = "https://pones.theponyapi.com/file/ponies/6247559689404416/thumb_small"
		}

		gif := fmt.Sprintf("in response to `/opsfun %s`\n\n%s", command[0], princessImage)
		return generateStandardSlashResponse(gif, "in_channel", "Princess", princessThumb)
	case "honk":
		honkImage, _ := getImage("goose")
		if honkImage == "" {
			honkImage = "Couldn't find a honk matching that query."
		}

		gif := fmt.Sprintf("in response to `/opsfun %s`\n\n%s", command[0], honkImage)
		return generateStandardSlashResponse(gif, "in_channel", "HONK", "")
	case "meow":
		honkImage, _ := getImage("cat")
		if honkImage == "" {
			honkImage = "Couldn't find a cat matching that query."
		}

		gif := fmt.Sprintf("in response to `/opsfun %s`\n\n%s", command[0], honkImage)
		return generateStandardSlashResponse(gif, "in_channel", "Meow", "")
	case "help", "--help", "-h":
		return generateStandardSlashResponse(helpMsg, "ephemeral", "OpsFun", "")
	default:
		return generateStandardSlashResponse(helpMsg, "ephemeral", "OpsFun", "")
	}

	return ""
}

type MMSlashResponse struct {
	ResponseType string `json:"response_type,omitempty"`
	Username     string `json:"username,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	Channel      string `json:"channel,omitempty"`
	Text         string `json:"text,omitempty"`
	GotoLocation string `json:"goto_location,omitempty"`
}

func generateStandardSlashResponse(text, respType, username, iconURL string) string {
	response := MMSlashResponse{
		ResponseType: respType,
		Text:         text,
		GotoLocation: "",
		Username:     username,
		IconURL:      iconURL,
	}

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Unable to marshal response")
		return ""
	}
	return string(b)
}

func getPony() (ponyImage, ponyThumb string) {
	for i := 0; i < 5; i++ {
		image, thumb, err := readPony("")
		if err != nil {
			continue
		}
		return image, thumb
	}

	return "", ""
}

// Only the properties we actually use.
type ponyResult struct {
	Pony ponyResultPony `json:"pony"`
}

type ponyResultPony struct {
	Representations ponyRepresentations `json:"representations"`
}

type ponyRepresentations struct {
	Small      string `json:"small"`
	ThumbSmall string `json:"thumbSmall"`
}

func readPony(ponyName string) (ponyImage string, ponyThumb string, err error) {
	client := http.Client{}
	uri := ponyURL + "?q=" + url.QueryEscape(ponyName)
	resp, err := client.Get(uri)
	if err != nil {
		return "", "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("no pony found")
	}
	var a ponyResult
	if err = json.NewDecoder(resp.Body).Decode(&a); err != nil {
		return "", "", fmt.Errorf("failed to decode response: %v", err)
	}
	return a.Pony.Representations.Small, a.Pony.Representations.ThumbSmall, nil
}
