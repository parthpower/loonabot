// Package loona
// everything about loona bot
package loona

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"regexp"
	"strings"

	"github.com/parthpower/loonabot/cmd/loona/static"
	"github.com/parthpower/loonabot/pkg/insta"
	"github.com/parthpower/loonabot/pkg/myscraper"

	"github.com/bregydoc/gtranslate"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

type Loona struct {
	Session     *session.Session
	InstaCookie string
}

type discordmsg struct {
	*gateway.MessageCreateEvent
	*Loona
	// TODO: get better way to pass info from match to action
	// this isn't thread safe
	extraArgs interface{}
}

func NewBot(token string, instacookie string) (*Loona, error) {
	s, err := session.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	l := &Loona{Session: s, InstaCookie: instacookie}
	l.Session.AddIntents(gateway.IntentGuildMessages)
	l.Session.AddHandler(l.handler)
	return l, nil
}

func (l *Loona) Start(ctx context.Context) error {
	logrus.Info("bot started")
	return l.Session.Open(ctx)
}

func (l *Loona) Stop() error {
	return l.Session.Close()
}

func (l *Loona) handler(msg *gateway.MessageCreateEvent) {
	m := &discordmsg{msg, l, []string{}}
	if msg.Author.Bot {
		return
	}
	m.do(m.matchNewMember, m.actionSendFileMsg("welcome to LOONA", static.WelcomeToLoona))
	m.do(m.matchSimpleRegex("test welcome to LOONA"), m.actionSendFileMsg("welcome to loona", static.WelcomeToLoona))
	m.do(m.matchSimpleRegex("haseul"), m.actionSendFileMsg("haseul omma", static.Haseul))
	m.do(m.matchSimpleRegex("hyunjin"), m.actionSendFileMsg("i am little meow meow", static.ImCat))
	m.do(m.matchSimpleRegex("kim lip"), m.actionSendFileMsg("KIMBERLY LIPPINGTON!", static.KimLip))
	m.do(m.matchSimpleRegex("i love loona"), m.actionSendMsgNoTTS("https://www.youtube.com/watch?v=xdvzjQzLYE0"))
	m.do(m.matchSimpleRegex("i .+ love loona"), m.actionSendMsgNoTTS("https://www.youtube.com/watch?v=Jw_7hfGpYbo"))
	m.do(m.matchSimpleRegex("apple"), m.actionSendMsg("fuck apple"))
	m.do(m.matchSimpleRegex("apple"), m.actionSendMsgNoTTS("https://gfycat.com/bowedmammothcrocodile"))
	m.do(m.matchSimpleRegex("chuu"), m.actionSendMsg("ten knee sue "))
	m.do(m.matchSimpleRegex("chuu"), m.actionSendMsgNoTTS("https://gfycat.com/honestweirddromaeosaur"))
	m.do(m.matchSimpleRegex("heejin"), m.actionSendFileMsg("jay map elle heejin", static.Heejin))
	m.do(m.matchSimpleRegex("aeong"), m.actionSendMsg("aeong"))
	m.do(m.matchSimpleRegex("aeong"), m.actionSendMsgNoTTS("https://gfycat.com/innocentsphericalarrowworm"))
	m.do(m.matchSimpleRegex("vivi "), m.actionSendMsg("jinjah moosaw, yeah baw may"))
	m.do(m.matchSimpleRegex("vivi "), m.actionSendMsgNoTTS("https://gfycat.com/decimalunimportantbeaver"))
	m.do(m.matchSimpleRegex("go ?wo+n"), m.actionSendMsgNoTTS("https://cdn.discordapp.com/attachments/819626268726263828/941467694090321940/best_member.mp4"))
	m.do(m.matchPrefix(".yt"), m.actionYTSearch())
	m.do(m.matchPrefix(".translate"), m.actionTranslate())
	m.do(m.matchInstagramURL(), m.actionSendInstagramContent())
}

func (m *discordmsg) matchInstagramURL() func() bool {
	return func() bool {
		re := regexp.MustCompile(`https?:\/\/(?:www\.)?instagram\.com\/(:?.*\/)?(:?tv|p|reel)\/([A-Za-z0-9_-]+)`)

		matches := re.FindAllStringSubmatch(m.Content, -1)
		if len(matches) < 1 {
			return false
		}
		for _, match := range matches {
			if len(match) > 1 {
				m.extraArgs = append(m.extraArgs.([]string), match[0])
			}
		}
		return len(m.extraArgs.([]string)) > 0
	}
}

func (m *discordmsg) matchPrefix(prefix string) func() bool {
	return func() bool {
		re := regexp.MustCompile(prefix + ` (.*)`)
		matches := re.FindStringSubmatch(m.Content)
		if len(matches) < 2 {
			return false
		}
		m.extraArgs = []string{matches[1]}
		return true
	}
}

func (m *discordmsg) matchNewMember() bool {
	return m.Type == discord.GuildMemberJoinMessage
}

func (m *discordmsg) matchSimpleRegex(match string) func() bool {
	return func() bool {
		if match, err := regexp.MatchString(`.*`+match+`.*`, strings.ToLower(m.Content)); match && err == nil {
			return true
		}
		return false
	}
}

func (m *discordmsg) actionSendMsgNoTTS(messages ...string) func() error {
	return func() error {
		var err error
		for _, msg := range messages {
			_, e := m.Loona.Session.SendMessageComplex(m.ChannelID, api.SendMessageData{
				Content: msg,
			})
			if e != nil {
				err = multierror.Append(err, e)
			}
		}
		return err
	}
}

func (m *discordmsg) actionSendMsg(messages ...string) func() error {
	return func() error {
		var err error
		for _, msg := range messages {
			_, e := m.Loona.Session.SendMessageComplex(m.ChannelID, api.SendMessageData{
				Content: msg,
				TTS:     true,
			})
			if e != nil {
				err = multierror.Append(err, e)
			}
		}
		return err
	}
}

func (m *discordmsg) actionSendFileMsg(msg string, fileFn func() (fs.File, error)) func() error {
	fn := func() error {
		f, err := fileFn()
		if err != nil {
			return err
		}
		name, _ := f.Stat()

		_, err = m.Loona.Session.SendMessageComplex(m.ChannelID, api.SendMessageData{
			Content: msg,
			TTS:     true,
			Files:   []sendpart.File{{Name: name.Name(), Reader: f}},
		})
		return err
	}
	return fn
}

func (m *discordmsg) actionYTSearch() func() error {
	return func() error {
		if m.extraArgs == nil || len(m.extraArgs.([]string)) < 1 {
			return fmt.Errorf("m.extraArgs not populated")
		}
		term := m.extraArgs.([]string)[0]
		results, err := myscraper.Search(term)
		if err != nil {
			return err
		}
		msg := ""

		if len(results) == 0 {
			msg = "no results! sorry but here's Hula Hoop https://www.youtube.com/watch?v=tnpUv1Vj5MA"
		} else {
			msg = results[0].URL
		}

		_, err = m.Session.SendMessage(m.ChannelID, msg)
		return err
	}
}

func (m *discordmsg) actionTranslate() func() error {
	return func() error {
		if m.extraArgs == nil || len(m.extraArgs.([]string)) < 1 {
			return fmt.Errorf("m.extraArgs not populated")
		}
		translated, err := gtranslate.TranslateWithParams(m.extraArgs.([]string)[0], gtranslate.TranslationParams{
			From: "auto",
			To:   "en",
		})
		if err != nil {
			return err
		}
		_, err = m.Session.SendMessage(m.ChannelID, translated)
		return err
	}
}

func (m *discordmsg) actionSendInstagramContent() func() error {
	return func() error {
		for _, u := range m.extraArgs.([]string) {
			media, err := insta.GetMediaFromUrl(u, m.InstaCookie)
			if err != nil {
				logrus.Errorf("GetPostByUrl: %v", err)
				continue
			}
			urls := media.DownloadURL
			logrus.Infof("instagram urls: %v", urls)
			f, closer, err := getSendPartFile(urls...)
			if err != nil || f == nil || len(f) == 0 || closer == nil {
				logrus.Error("failed getSendPartFile: " + err.Error())
				continue
			}
			defer closer()
			_, err = m.Loona.Session.SendMessageComplex(m.ChannelID, api.SendMessageData{
				Content: media.Caption,
				Files:   f,
			})
			if err != nil {
				logrus.Error("failed to send msg: " + err.Error())
			}
		}
		return nil
	}
}

func getSendPartFile(urls ...string) ([]sendpart.File, func(), error) {
	if len(urls) < 1 {
		return nil, nil, fmt.Errorf("empty urls")
	}
	f := []sendpart.File{}
	re := regexp.MustCompile(`\.mp4`)
	pipes := []io.ReadCloser{}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			logrus.Error("failed to fetch: " + url + " error: " + err.Error())
			continue
		}
		// defer resp.Body.Close()
		pipes = append(pipes, resp.Body)
		isMp4 := re.MatchString(url)
		name := "img.jpg"
		if isMp4 {
			name = "vid.mp4"
		}
		f = append(f, sendpart.File{Reader: resp.Body, Name: name})
	}
	closer := func() {
		for _, i := range pipes {
			i.Close()
		}
	}
	return f, closer, nil
}

func (m *discordmsg) actionDownloadTempFile(urls ...string) func() error {
	return func() error {
		if len(urls) < 1 {
			return nil
		}
		f := []sendpart.File{}
		re := regexp.MustCompile(`\.mp4`)
		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				return nil
			}
			defer resp.Body.Close()
			isMp4 := re.MatchString(url)
			name := "img.jpg"
			if isMp4 {
				name = "vid.mp4"
			}
			f = append(f, sendpart.File{Reader: resp.Body, Name: name})
		}
		_, err := m.Loona.Session.SendMessageComplex(m.ChannelID, api.SendMessageData{
			Content: "um",
			Files:   f,
		})
		return err
	}
}

func (m *discordmsg) do(match func() bool, action func() error) (bool, error) {
	if match() {
		// l.Info("matched")
		err := action()
		if err != nil {
			logrus.Error(err)
		}
		return true, err
	}
	return false, nil
}
