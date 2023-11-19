package config

type Config struct {
	LocationUrl string
	NextUrl     *string
	PreviousUrl *string
}

func ReadConfig () Config {
	conf := Config{
		LocationUrl: "https://pokeapi.co/api/v2/location/",
		NextUrl: 	 nil,
		PreviousUrl: nil,
	}

	return conf
}

func (c *Config) SetNextUrl(url *string) {
	c.NextUrl = url
}

func (c *Config) SetPreviousUrl(url *string) {
	c.PreviousUrl = url
}