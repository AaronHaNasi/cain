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
				Description    []string `yaml:"description"`
				AdvancesNumber struct {
					Current int `yaml:"current"`
					Max     int `yaml:"max"`
				} `yaml:"advances_number"`
			} `yaml:"advances"`
		}
		Agendas []struct {
			AgendaName string   `yaml:"name"`
			Abilities  []string `yaml:"abilities"`
		} `yaml:"agendas"`
		Blasphemies []struct {
			BlasphemyName string `yaml:"name"`
			Abilities     []struct {
				AbilityName string   `yaml:"name"`
				AbilityType string   `yaml:"type"`
				Cost        string   `yaml:"cost"`
				Keywords    []string `yaml:"keywords"`
				Description string   `yaml:"description"`
			} `yaml:"abilities"`
		} `yaml:"blasphemies"`
	} `yaml:"character_stats"`
	Other struct {
		Notes         string `yaml:"notes"`
		Questionnaire struct {
			Question string `yaml:"question"`
			Answer   string `yaml:"answer"`
		} `yaml:"questionnaire"`
	} `yaml:"character"`
}
