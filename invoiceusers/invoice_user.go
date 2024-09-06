package invoiceusers

import (
	"strings"
	"time"

	"github.com/nmpatzios/go-utils/resterrors"
)

type BranchStore struct {
	ID                      string `json:"id"`
	AA                      int    `json:"aa"`
	Eponimia                string `json:"eponimia"`
	Address                 string `json:"address"`
	AddressNumber           string `json:"address_number"`
	City                    string `json:"city"`
	Postcode                string `json:"postcode"`
	Drastiriotita           string `json:"drastiriotita"`
	IncludeToInvoiceCounter bool   `json:"include_to_invoice_counter"`
}

type Associated struct {
	ID       string `json:"id"`
	Accepted bool   `json:"accepted"`
	DEclined bool   `json:"declined"`
}

type Subscription struct {
	ID   string  `json:"id"`
	Cost float64 `json:"cost"`
}

type AADEMyData struct {
	AADEUserID       string `json:"aade_user_id"`
	SubscriptionKey  string `json:"subscription_key"`
	AutoSendToMydata bool   `json:"auto_send_to_mydata"`
}
type TaxisNetDetails struct {
	TaxisNetAFM      string           `json:"taxisnet_afm"`
	TaxisNetPassword string           `json:"taxisnet_password"`
	ForologikoProfil ForologikoProfil `json:"forologiko_profil"`
}
type BankDetails struct {
	BankName string `json:"bank_name"`
	Bic      string `json:"bic"`
	Iban     string `json:"iban"`
	IsNew    bool   `json:"is_new"`
}

type ForologikoProfil struct {
	AsfalistikiKatigoria map[string]interface{} `json:"asfalistiki_katigoria"`
	GenikiKatigoria      struct {
		Kanonikos     bool `json:"kanonikos"`
		Meiomenos     bool `json:"meiomenos"`
		Ypermeiomenos bool `json:"ypermeiomenos"`
	} `json:"geniki_katigoria"`

	GenikiKatigoriaUK struct {
		Kanonikos     bool `json:"kanonikos"`
		Meiomenos     bool `json:"meiomenos"`
		Ypermeiomenos bool `json:"ypermeiomenos"`
	} `json:"geniki_katigoria_uk"`
	EidikiKatigoriaNisia struct {
		Kanonikos     bool `json:"kanonikos"`
		Meiomenos     bool `json:"meiomenos"`
		Ypermeiomenos bool `json:"ypermeiomenos"`
	} `json:"eidiki_katigoria_nisia"`

	KatigoriaFpa struct {
		Apalasomenos bool `json:"apalasomenos"`
	} `json:"katigoria_fpa"`
	ParakratoumenoiForoi struct {
		ForosEleutheronErgolavon      bool `json:"foros_eleutheron_ergolavon"`
		ForosErgolavon                bool `json:"foros_ergolavon"`
		ProkatavliteosForos           bool `json:"prokatavliteos_foros"`
		ProkatavliteosForosDs         bool `json:"prokatavliteos_foros_ds"`
		ProkatavliteosForosMichanikon bool `json:"prokatavliteos_foros_michanikon"`
		ParoxiYpiresion               bool `json:"paroxi_ypiresion"`
	} `json:"parakratoumenoi_foroi"`
	ParakratoumenoiForoiUK struct {
		ParakratoumenosForos0  bool `json:"parakratoumenos_foros_0"`
		ParakratoumenosForos20 bool `json:"parakratoumenos_foros_20"`
		ParakratoumenosForos30 bool `json:"parakratoumenos_foros30"`
	} `json:"parakratoumenoi_foroi_uk"`
	VatExemptionCategory   int  `json:"vatExemption"`
	EndokinotikesYpiresies bool `json:"endokinotikes_ypiresies"`
}
type FinishProfileSettings struct {
	DontRemindMeAgain bool `json:"dont_remind_me_again"`
	RemindMeNextLogin bool `json:"remind_me_next_login"`
}

type User struct {
	ID                     string                   `json:"id"`
	Name                   string                   `json:"name"`
	Username               string                   `json:"username"`
	Surname                string                   `json:"surname"`
	Email                  string                   `json:"email"`
	Password               string                   `json:"password"`
	Afm                    string                   `json:"afm"`
	CompanysHouseNumber    string                   `json:"companys_house_number"`
	UtrNumber              string                   `json:"utr_number"`
	DiakritikosTitlos      string                   `json:"diakritikos_titlos"`
	Drastiriotita          string                   `json:"drastiriotita"`
	DrastiriotitaEn        string                   `json:"drastiriotita_en"`
	Doy                    string                   `json:"doy"`
	DoyEN                  string                   `json:"doy_en"`
	Eponimia               string                   `json:"eponimia"`
	EponimiaEN             string                   `json:"eponimia_en"`
	Address                string                   `json:"address"`
	AddressEN              string                   `json:"address_en"`
	AddressNumber          string                   `json:"address_number"`
	City                   string                   `json:"city"`
	CityEN                 string                   `json:"city_en"`
	Area                   string                   `json:"area"`
	Postcode               string                   `json:"postcode"`
	PostcodeEN             string                   `json:"postcode_en"`
	CreatedAt              time.Time                `json:"created_at"`
	UpdatedAt              time.Time                `json:"updated_at"`
	LastLogin              time.Time                `json:"last_login"`
	ValidUntil             time.Time                `json:"valid_until"`
	Active                 bool                     `json:"active"`
	Role                   string                   `json:"role"`
	Token                  string                   `json:"token"`
	ActivationToken        string                   `json:"activation_token"`
	Hostname               string                   `json:"hostname"`
	CompanyEmail           string                   `json:"company_email"`
	AccounEmail            string                   `json:"account_email"`
	MobileNumber           string                   `json:"mobile_number"`
	PhoneNumber            string                   `json:"phone_number"`
	TaxisNetDetails        TaxisNetDetails          `json:"taxisnet_details"`
	AllDrastiriotites      []map[string]interface{} `json:"all_drastiriotites"`
	BankDetails            []BankDetails            `json:"bank_details"`
	AADEMyData             AADEMyData               `json:"aade_mydata"`
	Country                string                   `json:"country"`
	AssociatedClients      []Associated             `json:"associated_clients"`
	PronoiaCategory        float64                  `json:"pronoia_category"`
	DimotikosForosCategory float64                  `json:"dimotikos_foros_category"`
	EpikourikoCategory     float64                  `json:"epikouriko_category"`
	MyAccountant           Associated               `json:"my_accountant"`
	Subscription           Subscription             `json:"subscription"`
	BranchStores           []BranchStore            `json:"branch_stores"`
	Plugins                []int64                  `json:"plugins"`
	FinishProfileSettings  FinishProfileSettings    `json:"finish_profile_settings"`
	SaveParastatika        []string                 `json:"save_parastatika"`
	ForosDiamonis          ForoDiamonisType         `json:"forosdiamonis"`
	EisprakseisTelous0310  EisprakseisTelousType    `json:"eispraksei_telous_0310"`
	EisprakseisTelous1112  EisprakseisTelousType    `json:"eispraksei_telous_1112"`
	Geminumber             string                   `json:"gemi_number"`
	CanCreateUsers         CanCreateUsers           `json:"can_create_users"`
	SimplyProvider         SimplyProvider           `json:"simply_provider"`
	IlydaProvider          IlydaProvider            `json:"ilyda_provider"`
	VivaWalletAuth         VivaWalletAuth           `json:"viva_wallet_auth"`
	DsdcProvider           DsdcProvider             `json:"dsdc_provider"`
	BuyerIdentifier        string                   `json:"buyer_identifier"`
	IsSimplyProviderUser   bool                     `json:"is_simply_provider_user"`
	IsIlydaProviderUser    bool                     `json:"is_ilyda_provider_user"`
	ThermalPrinter         bool                     `json:"thermal_printer"`
	Language               string                   `json:"language"`
	Chartosimo             bool                     `json:"chartosimo"`
	POSTerminals           []POSTerminal            `json:"pos_terminals"`
}

type POSTerminal struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Selected bool   `json:"selected"`
}

type VivaWalletAuth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	SourceCode   string `json:"source_code"`
}

type SimplyProvider struct {
	ApiKey        string    `json:"api_key"`
	CreatedAt     time.Time `json:"created_at"`
	ActivatedAt   time.Time `json:"activated_at"`
	Active        bool      `json:"active"`
	Parastatika   []string  `json:"parastatika"`
	ExpiredAt     time.Time `json:"expired_at"`
	ProviderFee   float64   `json:"provider_fee"`
	EmailSent     bool      `json:"email_sent"`
	EmailSendDate time.Time `json:"email_send_date"`
}

type DsdcProvider struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"created_at"`
	ActivatedAt   time.Time `json:"activated_at"`
	Active        bool      `json:"active"`
	Parastatika   []string  `json:"parastatika"`
	ExpiredAt     time.Time `json:"expired_at"`
	ProviderFee   float64   `json:"provider_fee"`
	EmailSent     bool      `json:"email_sent"`
	EmailSendDate time.Time `json:"email_send_date"`
}

type IlydaProvider struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"created_at"`
	ActivatedAt   time.Time `json:"activated_at"`
	Active        bool      `json:"active"`
	Parastatika   []string  `json:"parastatika"`
	ExpiredAt     time.Time `json:"expired_at"`
	ProviderFee   float64   `json:"provider_fee"`
	EmailSent     bool      `json:"email_sent"`
	EmailSendDate time.Time `json:"email_send_date"`
	IlydaFees     string    `json:"ilyda_fees"`
}

type SendEmailToUsers struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

type CanCreateUsers struct {
	CanCreate bool     `json:"can_create"`
	Users     []string `json:"users"`
}

type ChangePassword struct {
	UserID           string `json:"user_id"`
	CurrrentPassword string `json:"current_password"`
	NewPassword      string `json:"new_password"`
}

type Clients struct {
	Clients []string `json:"clients"`
}

type ElasticForgotPassword struct {
	Index         string          `json:"_index"`
	Type          string          `json:"_type"`
	ID            string          `json:"_id"`
	Score         float64         `json:"_score"`
	Source        SendGoldenEmail `json:"_source"`
	Authenticated string          `json:"type"`
}

type SendGoldenEmail struct {
	CustomerName    string    `json:"customer_name"`
	Email           string    `json:"email"`
	Subject         string    `json:"subject"`
	UserEmail       string    `json:"user_email"`
	CustomerMessage string    `json:"customer_message"`
	CompanyAddress  string    `json:"company_address"`
	ContactDetails  string    `json:"contact_details"`
	CompanyNickName string    `json:"company_nick_name"`
	ActivationToken string    `json:"activation_token"`
	ValidUntil      time.Time `json:"valid_until"`
	VatNumber       string    `json:"vat_number"`
}

type ActivateUser struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ForoDiamonis struct {
	Hotels050       ForoDiamonisType `json:"holets050"`
	Hotels150       ForoDiamonisType `json:"holets150"`
	Hotels300       ForoDiamonisType `json:"holets300"`
	Hotels400       ForoDiamonisType `json:"holets400"`
	RentedHotels050 ForoDiamonisType `json:"rented_holets050"`
}

type ForoDiamonisType struct {
	Type  int64   `json:"type"`
	Value float64 `json:"value"`
}

type EisprakseisTelous0310 struct {
	Hotels150                                    EisprakseisTelousType `json:"hotels_150"`
	Hotels300                                    EisprakseisTelousType `json:"hotels_300"`
	Hotels700                                    EisprakseisTelousType `json:"hotels_700"`
	Hotels1000                                   EisprakseisTelousType `json:"hotels_1000"`
	InvoiceEEDD                                  EisprakseisTelousType `json:"invoice_eedd"`
	InvoiceBB                                    EisprakseisTelousType `json:"invoice_bb"`
	InvoiceMonBB                                 EisprakseisTelousType `json:"invoice_mon_bb"`
	InvoiceAKTEE                                 EisprakseisTelousType `json:"invoice_aktee"`
	InvoiceVraxichoniaMisthosi                   EisprakseisTelousType `json:"invoice_vraxichonia_misthosi"`
	InvoiceVraxichoniaMisthosiMonKatLessThan80TM EisprakseisTelousType `json:"invoice_vraxichonia_misthosi_mon_kat_less_than_80_tm"`
	InvoiceVraxichoniaMisthosiMonKatMoreThan80TM EisprakseisTelousType `json:"invoice_vraxichonia_misthosi_mon_kat_more_than_80_tm"`
}

type EisprakseisTelous1112 struct {
	Hotels050                                    EisprakseisTelousType `json:"hotels_050"`
	Hotels150                                    EisprakseisTelousType `json:"hotels_150"`
	Hotels300                                    EisprakseisTelousType `json:"hotels_300"`
	Hotels400                                    EisprakseisTelousType `json:"hotels_400"`
	InvoiceEEDD                                  EisprakseisTelousType `json:"invoice_eedd"`
	InvoiceBB                                    EisprakseisTelousType `json:"invoice_bb"`
	InvoiceMonBB                                 EisprakseisTelousType `json:"invoice_mon_bb"`
	InvoiceAKTEE                                 EisprakseisTelousType `json:"invoice_aktee"`
	InvoiceVraxichoniaMisthosi                   EisprakseisTelousType `json:"invoice_vraxichonia_misthosi"`
	InvoiceVraxichoniaMisthosiMonKatLessThan80TM EisprakseisTelousType `json:"invoice_vraxichonia_misthosi_mon_kat_less_than_80_tm"`
	InvoiceVraxichoniaMisthosiMonKatMoreThan80TM EisprakseisTelousType `json:"invoice_vraxichonia_misthosi_mon_kat_more_than_80_tm"`
}

type EisprakseisTelousType struct {
	Type  int64   `json:"type"`
	Value float64 `json:"value"`
}

func (u *User) ToPublicUser(users []User) map[string]interface{} {
	allUsers := make([]string, 0)
	for _, user := range users {
		allUsers = append(allUsers, user.Eponimia)
	}
	return map[string]interface{}{
		"aade_user_id":       u.AADEMyData.AADEUserID,
		"aade_password":      u.AADEMyData.SubscriptionKey,
		"id":                 u.ID,
		"email":              u.Email,
		"afm":                u.Afm,
		"diakritikos_titlos": u.DiakritikosTitlos,
		"drastiriotita":      u.Drastiriotita,
		"doy":                u.Doy,
		"eponimia":           u.Eponimia,
		"created_users":      allUsers,
	}
}

func (user *User) MydataRegisterAPIValidate() resterrors.RestErr {
	user.Afm = strings.TrimSpace(user.Afm)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return resterrors.NewBadRequestError("invalid email address")
	}
	if user.Afm == "" {
		return resterrors.NewBadRequestError("invalid afm")
	}
	if user.Password == "" {
		return resterrors.NewBadRequestError("invalid password")
	}
	if user.AADEMyData.AADEUserID == "" {
		return resterrors.NewBadRequestError("invalid aade_user_id")
	}
	if user.AADEMyData.SubscriptionKey == "" {
		return resterrors.NewBadRequestError("invalid subscription_key")
	}

	return nil
}
