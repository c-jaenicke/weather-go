package weather

type ForecastResponse struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []List `json:"list"`
	City    City   `json:"city"`
}

type Forecast struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}
type Rain struct {
	ThreeH float64 `json:"3h"`
}
type Sys struct {
	Pod string `json:"pod"`
}
type List struct {
	Dt         int       `json:"dt"`
	Main       Forecast  `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`
	Rain       Rain      `json:"rain,omitempty"`
	Sys        Sys       `json:"sys"`
	DtTxt      string    `json:"dt_txt"`
}
type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type City struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

// ApiResponse struct symbolising standard response of API
type ApiResponse struct {
	Lat            float64  `json:"lat"`
	Lon            float64  `json:"lon"`
	Timezone       string   `json:"timezone"`
	TimezoneOffset int      `json:"timezone_offset"`
	Current        Current  `json:"current"`
	Hourly         []Hourly `json:"hourly"`
}

// Weather struct symbolising a weather
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Current struct symbolising current weather at a location
type Current struct {
	Dt         int       `json:"dt"`
	Sunrise    int       `json:"sunrise"`
	Sunset     int       `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        int       `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
}

type Snow struct {
	OneH float64 `json:"1h"`
}

// Hourly struct symbolising the forecast for an hour
type Hourly struct {
	Dt         int       `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        int       `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
	Snow       Snow      `json:"snow,omitempty"`
}

type GeocodeResponse []struct {
	Name       string     `json:"name"`
	LocalNames LocalNames `json:"local_names"`
	Lat        float64    `json:"lat"`
	Lon        float64    `json:"lon"`
	Country    string     `json:"country"`
}

type LocalNames struct {
	Fi string `json:"fi"`
	Rm string `json:"rm"`
	Nl string `json:"nl"`
	Mg string `json:"mg"`
	Lb string `json:"lb"`
	So string `json:"so"`
	Qu string `json:"qu"`
	Fa string `json:"fa"`
	De string `json:"de"`
	Kk string `json:"kk"`
	Fy string `json:"fy"`
	It string `json:"it"`
	Et string `json:"et"`
	Ga string `json:"ga"`
	Ms string `json:"ms"`
	Ie string `json:"ie"`
	Vi string `json:"vi"`
	Da string `json:"da"`
	Br string `json:"br"`
	Sv string `json:"sv"`
	Sl string `json:"sl"`
	Ht string `json:"ht"`
	Am string `json:"am"`
	Sr string `json:"sr"`
	Mr string `json:"mr"`
	Uz string `json:"uz"`
	Bs string `json:"bs"`
	Ro string `json:"ro"`
	Ky string `json:"ky"`
	Af string `json:"af"`
	Sh string `json:"sh"`
	Sq string `json:"sq"`
	Yo string `json:"yo"`
	Kv string `json:"kv"`
	Hu string `json:"hu"`
	Ug string `json:"ug"`
	Os string `json:"os"`
	ID string `json:"id"`
	Ty string `json:"ty"`
	Cs string `json:"cs"`
	Ja string `json:"ja"`
	An string `json:"an"`
	Jv string `json:"jv"`
	Sw string `json:"sw"`
	Sk string `json:"sk"`
	Oc string `json:"oc"`
	Ln string `json:"ln"`
	Lt string `json:"lt"`
	Ps string `json:"ps"`
	Ar string `json:"ar"`
	Ur string `json:"ur"`
	Fo string `json:"fo"`
	Zu string `json:"zu"`
	Iu string `json:"iu"`
	Tg string `json:"tg"`
	Se string `json:"se"`
	Si string `json:"si"`
	Az string `json:"az"`
	Eu string `json:"eu"`
	Eo string `json:"eo"`
	Mt string `json:"mt"`
	He string `json:"he"`
	My string `json:"my"`
	Te string `json:"te"`
	Tl string `json:"tl"`
	Es string `json:"es"`
	Bo string `json:"bo"`
	Pl string `json:"pl"`
	Lv string `json:"lv"`
	El string `json:"el"`
	Hy string `json:"hy"`
	Co string `json:"co"`
	Kn string `json:"kn"`
	Tt string `json:"tt"`
	Bg string `json:"bg"`
	Gl string `json:"gl"`
	La string `json:"la"`
	Gv string `json:"gv"`
	Is string `json:"is"`
	Yi string `json:"yi"`
	Wo string `json:"wo"`
	Kw string `json:"kw"`
	Ka string `json:"ka"`
	Ia string `json:"ia"`
	Cu string `json:"cu"`
	En string `json:"en"`
	Uk string `json:"uk"`
	Na string `json:"na"`
	Be string `json:"be"`
	Hr string `json:"hr"`
	Mk string `json:"mk"`
	Nn string `json:"nn"`
	Th string `json:"th"`
	Ko string `json:"ko"`
	Mn string `json:"mn"`
	Su string `json:"su"`
	Bi string `json:"bi"`
	Ml string `json:"ml"`
	Cv string `json:"cv"`
	No string `json:"no"`
	Fr string `json:"fr"`
	Io string `json:"io"`
	Ta string `json:"ta"`
	Gd string `json:"gd"`
	Pt string `json:"pt"`
	Lg string `json:"lg"`
	Ru string `json:"ru"`
	Cy string `json:"cy"`
	Ab string `json:"ab"`
	Tr string `json:"tr"`
	Gn string `json:"gn"`
	Ba string `json:"ba"`
	Ca string `json:"ca"`
	Ku string `json:"ku"`
	Hi string `json:"hi"`
	Mi string `json:"mi"`
	Li string `json:"li"`
	Bn string `json:"bn"`
	Zh string `json:"zh"`
	Sc string `json:"sc"`
}
