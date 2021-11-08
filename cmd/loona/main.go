// Package loona
// everything about loona bot
package loona

import (
	"context"
	"fmt"
	"io/fs"
	"regexp"
	"strings"

	"github.com/parthpower/loonabot/cmd/loona/static"
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
	Session *session.Session
}

type discordmsg struct {
	*gateway.MessageCreateEvent
	*Loona
	extraArgs []string
}

func NewBot(token string) (*Loona, error) {
	s, err := session.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	l := &Loona{Session: s}
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
	m.do(m.matchSimpleRegex("i love loona"), m.actionSendMsg("https://www.youtube.com/watch?v=Jw_7hfGpYbo"))
	m.do(m.matchSimpleRegex("apple"), m.actionSendMsg("fuck apple", "https://gfycat.com/bowedmammothcrocodile"))
	m.do(m.matchSimpleRegex("chuu"), m.actionSendMsg("ten knee sue "))
	m.do(m.matchSimpleRegex("chuu"), m.actionSendMsgNoTTS("https://gfycat.com/honestweirddromaeosaur"))
	m.do(m.matchPrefix(".yt"), m.actionYTSearch())
	m.do(m.matchPrefix(".translate"), m.actionTranslate())
}

func (m *discordmsg) matchPrefix(prefix string) func() bool {
	return func() bool {
		re := regexp.MustCompile(prefix + ` (.*)`)
		matches := re.FindStringSubmatch(m.Content)
		if len(matches) < 2 {
			return false
		}
		m.extraArgs = append(m.extraArgs, matches[1])
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
		if m.extraArgs == nil || len(m.extraArgs) < 1 {
			return fmt.Errorf("m.extraArgs not populated")
		}
		term := m.extraArgs[0]
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
		if m.extraArgs == nil || len(m.extraArgs) < 1 {
			return fmt.Errorf("m.extraArgs not populated")
		}
		translated, err := gtranslate.TranslateWithParams(m.extraArgs[0], gtranslate.TranslationParams{
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
