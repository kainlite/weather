package weatherapi

type Location struct {
	Id         string `form:"id" json:"id" binding:"required"`
	LocationId string `form:"location_id" json:"location_id" binding:"required"`
	UserId     string `form:"user_id" json:"user_id" binding:"required"`
	CreatedAt  string `form:"created_at" json:"created_at" binding:"required"`
}

type LocationKey struct {
	UserId     string `json:"user_id"`
	LocationId string `json:"location_id"`
}

// These structs were generated using:
// https://mholt.github.io/json-to-go/

type HourlyForecast struct {
	Metadata struct {
		Language      string `json:"language"`
		TransactionID string `json:"transaction_id"`
		Version       string `json:"version"`
		LocationID    string `json:"location_id"`
		Units         string `json:"units"`
		ExpireTimeGmt int    `json:"expire_time_gmt"`
		StatusCode    int    `json:"status_code"`
	} `json:"metadata"`
	Forecasts []struct {
		Class          string      `json:"class"`
		ExpireTimeGmt  int         `json:"expire_time_gmt"`
		FcstValid      int         `json:"fcst_valid"`
		FcstValidLocal string      `json:"fcst_valid_local"`
		Num            int         `json:"num"`
		DayInd         string      `json:"day_ind"`
		Temp           int         `json:"temp"`
		Dewpt          int         `json:"dewpt"`
		Hi             int         `json:"hi"`
		Wc             int         `json:"wc"`
		FeelsLike      int         `json:"feels_like"`
		IconExtd       int         `json:"icon_extd"`
		Wxman          string      `json:"wxman"`
		IconCode       int         `json:"icon_code"`
		Dow            string      `json:"dow"`
		Phrase12Char   string      `json:"phrase_12char"`
		Phrase22Char   string      `json:"phrase_22char"`
		Phrase32Char   string      `json:"phrase_32char"`
		SubphrasePt1   string      `json:"subphrase_pt1"`
		SubphrasePt2   string      `json:"subphrase_pt2"`
		SubphrasePt3   string      `json:"subphrase_pt3"`
		Pop            int         `json:"pop"`
		PrecipType     string      `json:"precip_type"`
		Qpf            int         `json:"qpf"`
		SnowQpf        int         `json:"snow_qpf"`
		Rh             int         `json:"rh"`
		Wspd           int         `json:"wspd"`
		Wdir           int         `json:"wdir"`
		WdirCardinal   string      `json:"wdir_cardinal"`
		Gust           interface{} `json:"gust"`
		Clds           int         `json:"clds"`
		Vis            int         `json:"vis"`
		Mslp           float64     `json:"mslp"`
		UvIndexRaw     int         `json:"uv_index_raw"`
		UvIndex        int         `json:"uv_index"`
		UvWarning      int         `json:"uv_warning"`
		UvDesc         string      `json:"uv_desc"`
		GolfIndex      interface{} `json:"golf_index"`
		GolfCategory   string      `json:"golf_category"`
		Severity       int         `json:"severity"`
	} `json:"forecasts"`
}

type DailyForecast struct {
	Metadata struct {
		Language      string `json:"language"`
		TransactionID string `json:"transaction_id"`
		Version       string `json:"version"`
		LocationID    string `json:"location_id"`
		Units         string `json:"units"`
		ExpireTimeGmt int    `json:"expire_time_gmt"`
		StatusCode    int    `json:"status_code"`
	} `json:"metadata"`
	Forecasts []struct {
		Class          string      `json:"class"`
		ExpireTimeGmt  int         `json:"expire_time_gmt"`
		FcstValid      int         `json:"fcst_valid"`
		FcstValidLocal string      `json:"fcst_valid_local"`
		Num            int         `json:"num"`
		MaxTemp        interface{} `json:"max_temp"`
		MinTemp        int         `json:"min_temp"`
		Torcon         interface{} `json:"torcon"`
		Stormcon       interface{} `json:"stormcon"`
		Blurb          interface{} `json:"blurb"`
		BlurbAuthor    interface{} `json:"blurb_author"`
		LunarPhaseDay  int         `json:"lunar_phase_day"`
		Dow            string      `json:"dow"`
		LunarPhase     string      `json:"lunar_phase"`
		LunarPhaseCode string      `json:"lunar_phase_code"`
		Sunrise        string      `json:"sunrise"`
		Sunset         string      `json:"sunset"`
		Moonrise       string      `json:"moonrise"`
		Moonset        string      `json:"moonset"`
		QualifierCode  interface{} `json:"qualifier_code"`
		Qualifier      interface{} `json:"qualifier"`
		Narrative      string      `json:"narrative"`
		Qpf            int         `json:"qpf"`
		SnowQpf        int         `json:"snow_qpf"`
		SnowRange      string      `json:"snow_range"`
		SnowPhrase     string      `json:"snow_phrase"`
		SnowCode       string      `json:"snow_code"`
		Night          struct {
			FcstValid          int         `json:"fcst_valid"`
			FcstValidLocal     string      `json:"fcst_valid_local"`
			DayInd             string      `json:"day_ind"`
			ThunderEnum        int         `json:"thunder_enum"`
			DaypartName        string      `json:"daypart_name"`
			LongDaypartName    string      `json:"long_daypart_name"`
			AltDaypartName     string      `json:"alt_daypart_name"`
			ThunderEnumPhrase  interface{} `json:"thunder_enum_phrase"`
			Num                int         `json:"num"`
			Temp               int         `json:"temp"`
			Hi                 int         `json:"hi"`
			Wc                 int         `json:"wc"`
			Pop                int         `json:"pop"`
			IconExtd           int         `json:"icon_extd"`
			IconCode           int         `json:"icon_code"`
			Wxman              string      `json:"wxman"`
			Phrase12Char       string      `json:"phrase_12char"`
			Phrase22Char       string      `json:"phrase_22char"`
			Phrase32Char       string      `json:"phrase_32char"`
			SubphrasePt1       string      `json:"subphrase_pt1"`
			SubphrasePt2       string      `json:"subphrase_pt2"`
			SubphrasePt3       string      `json:"subphrase_pt3"`
			PrecipType         string      `json:"precip_type"`
			Rh                 int         `json:"rh"`
			Wspd               int         `json:"wspd"`
			Wdir               int         `json:"wdir"`
			WdirCardinal       string      `json:"wdir_cardinal"`
			Clds               int         `json:"clds"`
			PopPhrase          string      `json:"pop_phrase"`
			TempPhrase         string      `json:"temp_phrase"`
			AccumulationPhrase string      `json:"accumulation_phrase"`
			WindPhrase         string      `json:"wind_phrase"`
			Shortcast          string      `json:"shortcast"`
			Narrative          string      `json:"narrative"`
			Qpf                int         `json:"qpf"`
			SnowQpf            int         `json:"snow_qpf"`
			SnowRange          string      `json:"snow_range"`
			SnowPhrase         string      `json:"snow_phrase"`
			SnowCode           string      `json:"snow_code"`
			VocalKey           string      `json:"vocal_key"`
			QualifierCode      interface{} `json:"qualifier_code"`
			Qualifier          interface{} `json:"qualifier"`
			UvIndexRaw         int         `json:"uv_index_raw"`
			UvIndex            int         `json:"uv_index"`
			UvWarning          int         `json:"uv_warning"`
			UvDesc             string      `json:"uv_desc"`
			GolfIndex          interface{} `json:"golf_index"`
			GolfCategory       string      `json:"golf_category"`
		} `json:"night"`
		Day struct {
			FcstValid          int         `json:"fcst_valid"`
			FcstValidLocal     string      `json:"fcst_valid_local"`
			DayInd             string      `json:"day_ind"`
			ThunderEnum        int         `json:"thunder_enum"`
			DaypartName        string      `json:"daypart_name"`
			LongDaypartName    string      `json:"long_daypart_name"`
			AltDaypartName     string      `json:"alt_daypart_name"`
			ThunderEnumPhrase  interface{} `json:"thunder_enum_phrase"`
			Num                int         `json:"num"`
			Temp               int         `json:"temp"`
			Hi                 int         `json:"hi"`
			Wc                 int         `json:"wc"`
			Pop                int         `json:"pop"`
			IconExtd           int         `json:"icon_extd"`
			IconCode           int         `json:"icon_code"`
			Wxman              string      `json:"wxman"`
			Phrase12Char       string      `json:"phrase_12char"`
			Phrase22Char       string      `json:"phrase_22char"`
			Phrase32Char       string      `json:"phrase_32char"`
			SubphrasePt1       string      `json:"subphrase_pt1"`
			SubphrasePt2       string      `json:"subphrase_pt2"`
			SubphrasePt3       string      `json:"subphrase_pt3"`
			PrecipType         string      `json:"precip_type"`
			Rh                 int         `json:"rh"`
			Wspd               int         `json:"wspd"`
			Wdir               int         `json:"wdir"`
			WdirCardinal       string      `json:"wdir_cardinal"`
			Clds               int         `json:"clds"`
			PopPhrase          string      `json:"pop_phrase"`
			TempPhrase         string      `json:"temp_phrase"`
			AccumulationPhrase string      `json:"accumulation_phrase"`
			WindPhrase         string      `json:"wind_phrase"`
			Shortcast          string      `json:"shortcast"`
			Narrative          string      `json:"narrative"`
			Qpf                int         `json:"qpf"`
			SnowQpf            int         `json:"snow_qpf"`
			SnowRange          string      `json:"snow_range"`
			SnowPhrase         string      `json:"snow_phrase"`
			SnowCode           string      `json:"snow_code"`
			VocalKey           string      `json:"vocal_key"`
			QualifierCode      interface{} `json:"qualifier_code"`
			Qualifier          interface{} `json:"qualifier"`
			UvIndexRaw         float64     `json:"uv_index_raw"`
			UvIndex            int         `json:"uv_index"`
			UvWarning          int         `json:"uv_warning"`
			UvDesc             string      `json:"uv_desc"`
			GolfIndex          int         `json:"golf_index"`
			GolfCategory       string      `json:"golf_category"`
		} `json:"day,omitempty"`
	} `json:"forecasts"`
}
