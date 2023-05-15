package soap

type options struct {
	Encoding          string // 目前只支持XML
	Authentication    string // 目前只支持NTLM
	WithAccountDomain bool   // account是否携带domain信息
}

type Option func(c *options)

func WithAccountDomain() Option {
	return func(c *options) {
		c.WithAccountDomain = true
	}
}
