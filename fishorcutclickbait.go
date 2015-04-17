package clickbait

import (
	"regexp"
	"strings"
)

var cbRegexp = regexp.MustCompile(`(?i)` + strings.Join([]string{
	`character are you`,
	`before you die`,
	`you probably didn['’]t`,
	`in your life`,
	`reasons you should`,
	`\d+ things that`,
	`\d+ most important`,
	`probably didn['’]t know`,
	`things you probably`,
	`\d+ signs you`,
	`of all time`,
	`\d+ reasons you`,
	`didn[']t know about`,
	`of the most`,
	`in real life`,
	`game of thrones`,
	`things that happen`,
	`\d+ (?:dogs|cats|animals) who`,
	`\d+ things you`,
	`is this the`,
	`\d+ signs you['’]re`,
	`things you didn['’]t`,
	`you should be`,
	`will make your`,
	`ways to eat`,
	`the \d+ most`,
	`a lot of`,
	`\d+ reasons why`,
	`will restore your`,
	`see what happens`,
	`what happens next`,
	`can['’]t miss`,
	`see how`,
	`asked \d+`,
	`life changed forever`,
	`never guess`,
	`blown away`,
	`blow your`,
	`when you see`,
}, `|`))

func IsClickbait(s string) bool {
	return cbRegexp.MatchString(s)
}
