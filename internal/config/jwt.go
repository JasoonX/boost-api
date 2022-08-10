package config

type JwtIssuer interface {
	Issuer() string
	SecretKey() []byte
}

type jwtIssuer struct {
	issuer    string
	secretKey []byte
}

type yamlJwtConfig struct {
	Issuer    string `yaml:"issuer"`
	SecretKey string `yaml:"secretKey"`
}

func NewJwtIssuer(issuer, secretKey string) JwtIssuer {
	return &jwtIssuer{
		issuer:    issuer,
		secretKey: []byte(secretKey),
	}
}

func (i *jwtIssuer) Issuer() string {
	return i.issuer
}

func (i *jwtIssuer) SecretKey() []byte {
	return i.secretKey
}
