package main

type Character struct {
	BasicInfo struct {
		CharacterName string `yaml:"name"`
		XID           int    `yaml:"xid"`
		Sex           string `yaml:"sex"`
		Height        string `yaml:"height"`
		Weight        int    `yaml:"weight"`
		Hair          string `yaml:"hair"`
		Eyes          string `yaml:"eyes"`
		Cat           int    `yaml:"cat"`
		Psyche        int    `yaml:"psyche"`
		Sin           struct {
			Current int `yaml:"current"`
			Max     int `yaml:"max"`
		} `yaml:"sin"`
		Bursts struct {
			Current int `yaml:"current"`
			Max     int `yaml:"max"`
		} `yaml:"bursts"`
		SinMarks struct {
			SinName     string `yaml:"sin_name"`
			Description string `yaml:"description"`
		} `yaml:"sin_marks"`
	} `yaml:"basic_info"`
	CharacterStats struct {
		Skills struct {
			Force         int `yaml:"force"`
			Coordination  int `yaml:"coordination"`
			Interfacing   int `yaml:"interfacing"`
			Surveillance  int `yaml:"surveillance"`
			Authority     int `yaml:"authority"`
			Covert        int `yaml:"covert"`
			Investigation int `yaml:"investigation"`
			Negotiation   int `yaml:"negotiation"`
			Connection    int `yaml:"connection"`
		} `yaml:"skills"`
		Advancement struct {
			Experience int `yaml:"experience"`
			Advances   struct {
				Description []string `yaml:"description"`
			}
		}
	} `yaml:"character_stats"`
}
