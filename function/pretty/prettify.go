package pretty

import (
	"encoding/json"
	"strings"

	"github.com/bbjj040471/mejson"
	"github.com/bbjj040471/transporter/function"
	"github.com/bbjj040471/transporter/log"
	"github.com/bbjj040471/transporter/message"
)

const (
	DefaultIndent = 2
)

var (
	DefaultPrettifier = &Prettify{Spaces: DefaultIndent}
)

func init() {
	function.Add(
		"pretty",
		func() function.Function {
			return DefaultPrettifier
		},
	)
}

type Prettify struct {
	Spaces int `json:"spaces"`
}

func (p *Prettify) Apply(msg message.Msg) (message.Msg, error) {
	d, _ := mejson.Unmarshal(msg.Data())
	b, _ := json.Marshal(d)
	if p.Spaces > 0 {
		b, _ = json.MarshalIndent(d, "", strings.Repeat(" ", p.Spaces))
	}
	log.Infof("\n%s", string(b))
	return msg, nil
}
