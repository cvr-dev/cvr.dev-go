package cvr

type Produktionsenhed struct {
	Aarsbeskaeftigelse      []ProduktionsenhedAarsbeskaeftigelse     `json:"aarsbeskaeftigelse"`
	Attributter             []ProduktionsenhedAttribut               `json:"attributter"`
	Beliggenhedsadresse     []ProduktionsenhedAdresse                `json:"beliggenhedsadresse"`
	Bibranche1              []ProduktionsenhedBranche                `json:"bibranche1"`
	Bibranche2              []ProduktionsenhedBranche                `json:"bibranche2"`
	Bibranche3              []ProduktionsenhedBranche                `json:"bibranche3"`
	BrancheAnsvarskode      *int                                     `json:"brancheAnsvarskode"`
	DataAdgang              int                                      `json:"dataAdgang"`
	DeltagerRelation        []ProduktionsenhedDeltagerRelation       `json:"deltagerRelation"`
	ElektroniskPost         []ProduktionsenhedKontaktoplysning       `json:"elektroniskPost"`
	Enhedsnummer            int                                      `json:"enhedsnummer"`
	Enhedstype              string                                   `json:"enhedstype"`
	FejlBeskrivelse         *string                                  `json:"fejlBeskrivelse"`
	FejlRegistreret         bool                                     `json:"fejlRegistreret"`
	FejlVedIndlaesning      bool                                     `json:"fejlVedIndlaesning"`
	Hovedbranche            []ProduktionsenhedBranche                `json:"hovedbranche"`
	Kvartalsbeskaeftigelse  []ProduktionsenhedKvartalsbeskaeftigelse `json:"kvartalsbeskaeftigelse"`
	Livsforloeb             []ProduktionsenhedLivsforloeb            `json:"livsforloeb"`
	NaermesteFremtidigeDato *string                                  `json:"naermesteFremtidigeDato"`
	Navne                   []ProduktionsenhedNavn                   `json:"navne"`
	PNummer                 int                                      `json:"pNummer"`
	Postadresse             []ProduktionsenhedAdresse                `json:"postadresse"`
	Metadata                ProduktionsenhedMetadata                 `json:"produktionsEnhedMetadata"`
	Reklamebeskyttet        bool                                     `json:"reklamebeskyttet"`
	SamtID                  int                                      `json:"samtId"`
	SidstIndlaest
	SidstOpdateret
	TelefaxNummer       []VirksomhedKontaktoplysning          `json:"telefaxNummer"`
	TelefonNummer       []VirksomhedKontaktoplysning          `json:"telefonNummer"`
	Virksomhedsrelation []ProduktionsenhedVirksomhedsrelation `json:"virksomhedsrelation"`

	RawJSON []byte
}

type ProduktionsenhedAttribut struct {
	Sekvensnummer int                              `json:"sekvensnr"`
	Type          string                           `json:"type"`
	Vaerditype    string                           `json:"vaerditype"`
	Vaerdier      []ProduktionsenhedAttributVaerdi `json:"vaerdier"`
}

type ProduktionsenhedAttributVaerdi struct {
	Vaerdi  string `json:"vaerdi"`
	Periode `json:"periode"`
	SidstOpdateret
}

type ProduktionsenhedAarsbeskaeftigelse struct {
	Aar                             int      `json:"aar"`
	AntalAarsvaerk                  *float32 `json:"antalAarsvaerk"`
	AntalAnsatte                    *int     `json:"antalAnsatte"`
	AntalInklusivEjere              *int     `json:"antalInklusivEjere"`
	IntervalKodeAntalAarsvaerk      string   `json:"intervalKodeAntalAarsvaerk"`
	IntervalKodeAntalAnsatte        string   `json:"intervalKodeAntalAnsatte"`
	IntervalKodeAntalInklusiveEjere string   `json:"intervalKodeAntalInklusivEjere"`
	SidstOpdateret
}

type ProduktionsenhedKvartalsbeskaeftigelse struct {
	Aar                             int      `json:"aar"`
	Kvartal                         int      `json:"kvartal"`
	AntalAarsvaerk                  *float32 `json:"antalAarsvaerk"`
	AntalAnsatte                    *int     `json:"antalAnsatte"`
	AntalInklusivEjere              *int     `json:"antalInklusivEjere"`
	IntervalKodeAntalAarsvaerk      *string  `json:"intervalKodeAntalAarsvaerk"`
	IntervalKodeAntalAnsatte        *string  `json:"intervalKodeAntalAnsatte"`
	IntervalKodeAntalInklusiveEjere *string  `json:"intervalKodeAntalInklusivEjere"`
	SidstOpdateret
}

type ProduktionsenhedMaanedsbeskaeftigelse struct {
	Aar                        int      `json:"aar"`
	Maaned                     int      `json:"maaned"`
	AntalAarsvaerk             *float32 `json:"antalAarsvaerk"`
	AntalAnsatte               *int     `json:"antalAnsatte"`
	IntervalKodeAntalAarsvaerk *string  `json:"intervalKodeAntalAarsvaerk"`
	IntervalKodeAntalAnsatte   *string  `json:"intervalKodeAntalAnsatte"`
	SidstOpdateret
}

type ProduktionsenhedVirksomhedsrelation struct {
	CVRNummer int `json:"cvrNummer"`
	Periode   `json:"periode"`
	SidstOpdateret
}

type ProduktionsenhedMetadata struct {
	NyesteAarsbeskaeftigelse        *ProduktionsenhedAarsbeskaeftigelse     `json:"nyesteAarsbeskaeftigelse"`
	NyesteBeliggenhedsadresse       *ProduktionsenhedAdresse                `json:"nyesteBeliggenhedsadresse"`
	NyesteErstMaanedsbeskaeftigelse *ProduktionsenhedMaanedsbeskaeftigelse  `json:"nyesteErstMaanedsbeskaeftigelse"`
	NyesteBibranche1                *ProduktionsenhedBranche                `json:"nyesteBibranche1"`
	NyesteBibranche2                *ProduktionsenhedBranche                `json:"nyesteBibranche2"`
	NyesteBibranche3                *ProduktionsenhedBranche                `json:"nyesteBibranche3"`
	NyesteCVRNummerRelation         *int                                    `json:"nyesteCvrNummerRelation"`
	NyesteHovedbranche              *ProduktionsenhedBranche                `json:"nyesteHovedbranche"`
	NyesteKontaktoplysninger        []string                                `json:"nyesteKontaktoplysninger"`
	NyesteKvartalsbeskaeftigelse    *ProduktionsenhedKvartalsbeskaeftigelse `json:"nyesteKvartalsbeskaeftigelse"`
	NyesteNavn                      ProduktionsenhedNavn                    `json:"nyesteNavn"`
	SammensatStatus                 string                                  `json:"sammensatStatus"`
}

type ProduktionsenhedNavn struct {
	Navn    string `json:"navn"`
	Periode `json:"periode"`
	SidstOpdateret
}

type ProduktionsenhedLivsforloeb struct {
	Periode `json:"periode"`
	SidstOpdateret
}

type ProduktionsenhedBranche struct {
	Branchekode  string `json:"branchekode"`
	Branchetekst string `json:"branchetekst"`
	Periode      `json:"periode"`
	SidstOpdateret
}

type ProduktionsenhedKontaktoplysning struct {
	Hemmelig         bool   `json:"hemmelig"`
	Kontaktoplysning string `json:"kontaktoplysning"`
	Periode          `json:"periode"`
	SidstOpdateret
}

type ProduktionsenhedKommune struct {
	Kommunekode int    `json:"kommuneKode"`
	KommuneNavn string `json:"kommuneNavn"`
	Periode     `json:"periode"`
	SidstOpdateret
}

type ProduktionsenhedAdresse struct {
	BogstavFra   *string                 `json:"bogstavFra"`
	BogstavTil   *string                 `json:"bogstavTil"`
	Bynavn       *string                 `json:"bynavn"`
	COnavn       *string                 `json:"conavn"`
	Etage        *string                 `json:"etage"`
	Fritekst     *string                 `json:"fritekst"`
	HusnummerFra *int                    `json:"husnummerFra"`
	AdresseID    *string                 `json:"adresseId"`
	HusnummerTil *int                    `json:"husnummerTil"`
	Kommune      ProduktionsenhedKommune `json:"kommune"`
	Landekode    *string                 `json:"landekode"`
	Periode      `json:"periode"`
	Postboks     *string `json:"postboks"`
	Postdistrikt *string `json:"postdistrikt"`
	Postnummer   int     `json:"postnummer"`
	Sidedoer     *string `json:"sidedoer"`
	SidstOpdateret
	SidstValideret *string `json:"sidstValideret"`
	Vejkode        *int    `json:"vejkode"`
	Vejnavn        *string `json:"vejnavn"`
}

type ProduktionsenhedDeltagerRelation struct {
}
