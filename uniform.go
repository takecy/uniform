package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/takecy/uniform/aws"
	"github.com/takecy/uniform/cli"
)

var usage = `Usage: uniform [options...]
Options:
  -k  AWS access key or environment variables(UNIFORM_AWS_KEY).
  -s  AWS secret key or environment variables(UNIFORM_AWS_SECRET).
  -r  AWS region (ex: ap-northeast-1) or environment variables(UNIFORM_AWS_REGION).
  -t  AWS instance Tags for list targets. (ex: Environment=Production,Name=1a01,Type=API).
  -a  Alive API URL (default: "/alive").
  -p  Alive API URL's port (default: 80).
  -v  Alive API response(json) version property (default: "version").
`

var (
	k = flag.String("k", "", "")
	s = flag.String("s", "", "")
	r = flag.String("r", "", "")
	t = flag.String("t", "", "")
	a = flag.String("a", "/alive", "")
	p = flag.Int("p", 80, "")
	v = flag.String("v", "version", "")
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	if *k == "" {
		awsKey := os.Getenv("UNIFORM_AWS_KEY")
		k = &awsKey
	}
	if *s == "" {
		awsSecret := os.Getenv("UNIFORM_AWS_SECRET")
		s = &awsSecret
	}
	if *r == "" {
		awsRegion := os.Getenv("UNIFORM_AWS_REGION")
		r = &awsRegion
	}

	if *k == "" || *s == "" || *r == "" {
		fmt.Fprint(os.Stderr, "AWS setting not enought")
		os.Exit(1)
	}

	tags := strings.Split(*t, ",")
	tagM := make(map[string]string, len(tags))
	for _, tag := range tags {
		tt := strings.Split(tag, "=")
		if len(tt) != 2 {
			fmt.Fprintf(os.Stderr, "bad.tag.formt %V", tt)
			continue
		}
		tagM[tt[0]] = tt[1]
	}

	ai := aws.NewAWS(*k, *s, *r, tagM)
	ci := cli.NewCli(*a, *v, *p)

	err := run(ai, ci)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func usageAndExit(message string) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func run(a *aws.AWS, c *cli.Client) (err error) {
	addresses, err := a.ListAddresses()
	if err != nil {
		return
	}
	if len(addresses) == 0 {
		return errors.New("not target instance.")
	}

	incorrectAddresses, err := c.Alive(addresses)
	if err != nil {
		return
	}

	if len(incorrectAddresses) > 0 {
		return errors.New(fmt.Sprintf("exists incorrect addresses %v", incorrectAddresses))
	}

	return
}
