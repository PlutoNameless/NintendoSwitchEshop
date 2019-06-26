package nsEshop

import (
	"time"
)

const JP_GET_GAMES_CURRENT = "https://www.nintendo.co.jp/data/software/xml-system/switch-onsale.xml"
const JP_GET_GAMES_COMING = "https://www.nintendo.co.jp/data/software/xml-system/switch-coming.xml"

type GameJapan struct {
	LinkURL          string `xml:"LinkURL" json:"link_url"`
	LinkTarget       string `xml:"LinkTarget" json:"link_target"`
	ScreenshotImgURL string `xml:"ScreenshotImgURL,omitempty" json:"screenshot_img_url,omitempty"`
	ThumbVariation   string `xml:"ThumbVariation"`
	ComingThumb      string `xml:"ComingThumb" json:"coming_thumb"`
	TitleName        string `xml:"TitleName" json:"title_name"`
	TitleNameRuby    string `xml:"TitleNameRuby" json:"title_name_ruby"`
	SoftType         string `xml:"SoftType" json:"soft_type"`
	SalesDateStr     string `xml:"SalesDateStr" json:"sales_date_str"`
	MakerName        string `xml:"MakerName" json:"maker_name"`
	Hard             string `xml:"Hard" json:"hard"`
	Memo             string `xml:"Memo" json:"memo"`
	Nsuid            string `xml:"Nsuid,omitempty" json:"nsuid,omitempty"`
}

type GameJapanXML struct {
	LastUpdate string      `xml:"last_update,attr"`
	Games      []GameJapan `xml:"TitleInfo"`
}

type GameEurope struct {
	FsID                             string      `json:"fs_id"`
	ChangeDate                       time.Time   `json:"change_date"`
	URL                              string      `json:"url"`
	Type                             string      `json:"type"`
	DatesReleasedDts                 []time.Time `json:"dates_released_dts"`
	ClubNintendo                     bool        `json:"club_nintendo"`
	PrettyDateS                      string      `json:"pretty_date_s"`
	PlayModeTvModeB                  bool        `json:"play_mode_tv_mode_b"`
	PlayModeHandheldModeB            bool        `json:"play_mode_handheld_mode_b"`
	ProductCodeTxt                   []string    `json:"product_code_txt"`
	ImageURLSqS                      string      `json:"image_url_sq_s"`
	PgS                              string      `json:"pg_s"`
	GiftFinderDetailPageImageURLS    string      `json:"gift_finder_detail_page_image_url_s"`
	CompatibleController             []string    `json:"compatible_controller"`
	ImageURL                         string      `json:"image_url"`
	OriginallyForT                   string      `json:"originally_for_t"`
	PaidSubscriptionRequiredB        bool        `json:"paid_subscription_required_b"`
	CloudSavesB                      bool        `json:"cloud_saves_b"`
	DigitalVersionB                  bool        `json:"digital_version_b"`
	TitleExtrasTxt                   []string    `json:"title_extras_txt"`
	ImageURLH2X1S                    string      `json:"image_url_h2x1_s"`
	SystemType                       []string    `json:"system_type"`
	AgeRatingSortingI                int         `json:"age_rating_sorting_i"`
	GameCategoriesTxt                []string    `json:"game_categories_txt"`
	PlayModeTabletopModeB            bool        `json:"play_mode_tabletop_mode_b,omitempty"`
	Publisher                        string      `json:"publisher"`
	ProductCodeSs                    []string    `json:"product_code_ss"`
	Excerpt                          string      `json:"excerpt"`
	NsuidTxt                         []string    `json:"nsuid_txt"`
	DateFrom                         time.Time   `json:"date_from"`
	LanguageAvailability             []string    `json:"language_availability"`
	PriceHasDiscountB                bool        `json:"price_has_discount_b"`
	PriceDiscountPercentageF         float64     `json:"price_discount_percentage_f"`
	Title                            string      `json:"title"`
	SortingTitle                     string      `json:"sorting_title"`
	CopyrightS                       string      `json:"copyright_s"`
	GiftFinderCarouselImageURLS      string      `json:"gift_finder_carousel_image_url_s"`
	WishlistEmailSquareImageURLS     string      `json:"wishlist_email_square_image_url_s"`
	PlayersTo                        int         `json:"players_to"`
	WishlistEmailBanner640WImageURLS string      `json:"wishlist_email_banner640w_image_url_s"`
	VoiceChatB                       bool        `json:"voice_chat_b"`
	PlayableOnTxt                    []string    `json:"playable_on_txt"`
	PrettyGameCategoriesTxt          []string    `json:"pretty_game_categories_txt"`
	GiftFinderWishlistImageURLS      string      `json:"gift_finder_wishlist_image_url_s"`
	GameCategory                     []string    `json:"game_category"`
	SystemNamesTxt                   []string    `json:"system_names_txt"`
	PrettyAgeratingS                 string      `json:"pretty_agerating_s"`
	PlayersFrom                      int         `json:"players_from"`
	AgeRatingType                    string      `json:"age_rating_type"`
	PriceSortingF                    float64     `json:"price_sorting_f"`
	PriceLowestF                     float64     `json:"price_lowest_f"`
	AgeRatingValue                   string      `json:"age_rating_value"`
	PhysicalVersionB                 bool        `json:"physical_version_b"`
	WishlistEmailBanner460WImageURLS string      `json:"wishlist_email_banner460w_image_url_s"`
	HitsI                            int         `json:"hits_i"`
	Version                          int64       `json:"_version_"`
	Internet                         bool        `json:"internet,omitempty"`
}

type GamesEuropeJSON struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		Params struct {
			Q     string `json:"q"`
			Start string `json:"start"`
			Fq    string `json:"fq"`
			Sort  string `json:"sort"`
			Rows  string `json:"rows"`
			Wt    string `json:"wt"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response struct {
		NumFound int          `json:"numFound"`
		Start    int          `json:"start"`
		Games    []GameEurope `json:"docs"`
	} `json:"response"`
}

type Eshop struct {
	Code     string `json:"code"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
	Region   string `json:"region"`
}

type Region int

const (
	AMERICAS Region = iota
	EUROPE
	ASIA
)

type PriceResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Personalized bool   `json:"personalized"`
	Country      string `json:"country"`
	Prices       []struct {
		TitleID      int64  `json:"title_id"`
		SalesStatus  string `json:"sales_status"`
		RegularPrice struct {
			Amount        string    `json:"amount"`
			Currency      string    `json:"currency"`
			RawValue      string    `json:"raw_value"`
			StartDatetime time.Time `json:"start_datetime"`
			EndDatetime   time.Time `json:"end_datetime"`
		} `json:"regular_price"`
		DiscountPrice struct {
			Amount        string    `json:"amount"`
			Currency      string    `json:"currency"`
			RawValue      string    `json:"raw_value"`
			StartDatetime time.Time `json:"start_datetime"`
			EndDatetime   time.Time `json:"end_datetime"`
		} `json:"discount_price"`
	} `json:"prices"`
}

const PRICE_LIST_LIMIT = 50
