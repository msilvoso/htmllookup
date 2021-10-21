package htmllookup

import "testing"

func Test_searchableHtmlPage_itemsJson(t *testing.T) {
	h := New()
	err := h.LoadData("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	err = h.headerJson()
	if err != nil {
		t.Fatal(err)
	}
	err = h.itemsJson()
	if err != nil {
		t.Fatal(err)
	}
	should := `[{"Age":"22","Index":"1","Movie":"Seventh Heaven, Street Angel and Sunrise: A Song of Two Humans","Name":"Janet Gaynor","Year":"1928","normalized_search_column":"1192822JanetGaynorSeventhHeaven,StreetAngelandSunrise:ASongofTwoHumans"},{"Age":"37","Index":"2","Movie":"Coquette","Name":"Mary Pickford","Year":"1929","normalized_search_column":"2192937MaryPickfordCoquette"},{"Age":"28","Index":"3","Movie":"The Divorcee","Name":"Norma Shearer","Year":"1930","normalized_search_column":"3193028NormaShearerTheDivorcee"},{"Age":"63","Index":"4","Movie":"Min and Bill","Name":"Marie Dressler","Year":"1931","normalized_search_column":"4193163MarieDresslerMinandBill"},{"Age":"32","Index":"5","Movie":"The Sin of Madelon Claudet","Name":"Helen Hayes","Year":"1932","normalized_search_column":"5193232HelenHayesTheSinofMadelonClaudet"},{"Age":"26","Index":"6","Movie":"Morning Glory","Name":"Katharine Hepburn","Year":"1933","normalized_search_column":"6193326KatharineHepburnMorningGlory"},{"Age":"31","Index":"7","Movie":"It Happened One Night","Name":"Claudette Colbert","Year":"1934","normalized_search_column":"7193431ClaudetteColbertItHappenedOneNight"},{"Age":"27","Index":"8","Movie":"Dangerous","Name":"Bette Davis","Year":"1935","normalized_search_column":"8193527BetteDavisDangerous"},{"Age":"27","Index":"9","Movie":"The Great Ziegfeld","Name":"Luise Rainer","Year":"1936","normalized_search_column":"9193627LuiseRainerTheGreatZiegfeld"},{"Age":"28","Index":"10","Movie":"The Good Earth","Name":"Luise Rainer","Year":"1937","normalized_search_column":"10193728LuiseRainerTheGoodEarth"},{"Age":"30","Index":"11","Movie":"Jezebel","Name":"Bette Davis","Year":"1938","normalized_search_column":"11193830BetteDavisJezebel"},{"Age":"26","Index":"12","Movie":"Gone with the Wind","Name":"Vivien Leigh","Year":"1939","normalized_search_column":"12193926VivienLeighGonewiththeWind"},{"Age":"29","Index":"13","Movie":"Kitty Foyle","Name":"Ginger Rogers","Year":"1940","normalized_search_column":"13194029GingerRogersKittyFoyle"},{"Age":"24","Index":"14","Movie":"Suspicion","Name":"Joan Fontaine","Year":"1941","normalized_search_column":"14194124JoanFontaineSuspicion"},{"Age":"38","Index":"15","Movie":"Mrs. Miniver","Name":"Greer Garson","Year":"1942","normalized_search_column":"15194238GreerGarsonMrs.Miniver"},{"Age":"25","Index":"16","Movie":"The Song of Bernadette","Name":"Jennifer Jones","Year":"1943","normalized_search_column":"16194325JenniferJonesTheSongofBernadette"},{"Age":"29","Index":"17","Movie":"Gaslight","Name":"Ingrid Bergman","Year":"1944","normalized_search_column":"17194429IngridBergmanGaslight"},{"Age":"40","Index":"18","Movie":"Mildred Pierce","Name":"Joan Crawford","Year":"1945","normalized_search_column":"18194540JoanCrawfordMildredPierce"},{"Age":"30","Index":"19","Movie":"To Each His Own","Name":"Olivia de Havilland","Year":"1946","normalized_search_column":"19194630OliviadeHavillandToEachHisOwn"},{"Age":"35","Index":"20","Movie":"The Farmer's Daughter","Name":"Loretta Young","Year":"1947","normalized_search_column":"20194735LorettaYoungTheFarmer'sDaughter"},{"Age":"32","Index":"21","Movie":"Johnny Belinda","Name":"Jane Wyman","Year":"1948","normalized_search_column":"21194832JaneWymanJohnnyBelinda"},{"Age":"33","Index":"22","Movie":"The Heiress","Name":"Olivia de Havilland","Year":"1949","normalized_search_column":"22194933OliviadeHavillandTheHeiress"},{"Age":"29","Index":"23","Movie":"Born Yesterday","Name":"Judy Holliday","Year":"1950","normalized_search_column":"23195029JudyHollidayBornYesterday"},{"Age":"38","Index":"24","Movie":"A Streetcar Named Desire","Name":"Vivien Leigh","Year":"1951","normalized_search_column":"24195138VivienLeighAStreetcarNamedDesire"},{"Age":"54","Index":"25","Movie":"Come Back, Little Sheba","Name":"Shirley Booth","Year":"1952","normalized_search_column":"25195254ShirleyBoothComeBack,LittleSheba"},{"Age":"24","Index":"26","Movie":"Roman Holiday","Name":"Audrey Hepburn","Year":"1953","normalized_search_column":"26195324AudreyHepburnRomanHoliday"},{"Age":"25","Index":"27","Movie":"The Country Girl","Name":"Grace Kelly","Year":"1954","normalized_search_column":"27195425GraceKellyTheCountryGirl"},{"Age":"48","Index":"28","Movie":"The Rose Tattoo","Name":"Anna Magnani","Year":"1955","normalized_search_column":"28195548AnnaMagnaniTheRoseTattoo"},{"Age":"41","Index":"29","Movie":"Anastasia","Name":"Ingrid Bergman","Year":"1956","normalized_search_column":"29195641IngridBergmanAnastasia"},{"Age":"28","Index":"30","Movie":"The Three Faces of Eve","Name":"Joanne Woodward","Year":"1957","normalized_search_column":"30195728JoanneWoodwardTheThreeFacesofEve"},{"Age":"41","Index":"31","Movie":"I Want to Live!","Name":"Susan Hayward","Year":"1958","normalized_search_column":"31195841SusanHaywardIWanttoLive!"},{"Age":"39","Index":"32","Movie":"Room at the Top","Name":"Simone Signoret","Year":"1959","normalized_search_column":"32195939SimoneSignoretRoomattheTop"},{"Age":"29","Index":"33","Movie":"BUtterfield 8","Name":"Elizabeth Taylor","Year":"1960","normalized_search_column":"33196029ElizabethTaylorBUtterfield8"},{"Age":"27","Index":"34","Movie":"Two Women","Name":"Sophia Loren","Year":"1961","normalized_search_column":"34196127SophiaLorenTwoWomen"},{"Age":"31","Index":"35","Movie":"The Miracle Worker","Name":"Anne Bancroft","Year":"1962","normalized_search_column":"35196231AnneBancroftTheMiracleWorker"},{"Age":"31","Index":"36","Movie":"Hud","Name":"Patricia Neal","Year":"1963","normalized_search_column":"36196331PatriciaNealHud"},{"Age":"29","Index":"37","Movie":"Mary Poppins","Name":"Julie Andrews","Year":"1964","normalized_search_column":"37196429JulieAndrewsMaryPoppins"},{"Age":"25","Index":"38","Movie":"Darling","Name":"Julie Christie","Year":"1965","normalized_search_column":"38196525JulieChristieDarling"},{"Age":"35","Index":"39","Movie":"Who's Afraid of Virginia Woolf?","Name":"Elizabeth Taylor","Year":"1966","normalized_search_column":"39196635ElizabethTaylorWho'sAfraidofVirginiaWoolf?"},{"Age":"60","Index":"40","Movie":"Guess Who's Coming to Dinner","Name":"Katharine Hepburn","Year":"1967","normalized_search_column":"40196760KatharineHepburnGuessWho'sComingtoDinner"},{"Age":"61","Index":"41","Movie":"The Lion in Winter","Name":"Katharine Hepburn","Year":"1968","normalized_search_column":"41196861KatharineHepburnTheLioninWinter"},{"Age":"26","Index":"42","Movie":"Funny Girl","Name":"Barbra Streisand","Year":"1969","normalized_search_column":"42196926BarbraStreisandFunnyGirl"},{"Age":"35","Index":"43","Movie":"The Prime of Miss Jean Brodie","Name":"Maggie Smith","Year":"1970","normalized_search_column":"43197035MaggieSmithThePrimeofMissJeanBrodie"},{"Age":"34","Index":"44","Movie":"Women in Love","Name":"Glenda Jackson","Year":"1971","normalized_search_column":"44197134GlendaJacksonWomeninLove"},{"Age":"34","Index":"45","Movie":"Klute","Name":"Jane Fonda","Year":"1972","normalized_search_column":"45197234JaneFondaKlute"},{"Age":"27","Index":"46","Movie":"Cabaret","Name":"Liza Minnelli","Year":"1973","normalized_search_column":"46197327LizaMinnelliCabaret"},{"Age":"37","Index":"47","Movie":"A Touch of Class","Name":"Glenda Jackson","Year":"1974","normalized_search_column":"47197437GlendaJacksonATouchofClass"},{"Age":"42","Index":"48","Movie":"Alice Doesn't Live Here Anymore","Name":"Ellen Burstyn","Year":"1975","normalized_search_column":"48197542EllenBurstynAliceDoesn'tLiveHereAnymore"},{"Age":"41","Index":"49","Movie":"One Flew Over the Cuckoo's Nest","Name":"Louise Fletcher","Year":"1976","normalized_search_column":"49197641LouiseFletcherOneFlewOvertheCuckoo'sNest"},{"Age":"36","Index":"50","Movie":"Network","Name":"Faye Dunaway","Year":"1977","normalized_search_column":"50197736FayeDunawayNetwork"},{"Age":"32","Index":"51","Movie":"Annie Hall","Name":"Diane Keaton","Year":"1978","normalized_search_column":"51197832DianeKeatonAnnieHall"},{"Age":"41","Index":"52","Movie":"Coming Home","Name":"Jane Fonda","Year":"1979","normalized_search_column":"52197941JaneFondaComingHome"},{"Age":"33","Index":"53","Movie":"Norma Rae","Name":"Sally Field","Year":"1980","normalized_search_column":"53198033SallyFieldNormaRae"},{"Age":"31","Index":"54","Movie":"Coal Miner's Daughter","Name":"Sissy Spacek","Year":"1981","normalized_search_column":"54198131SissySpacekCoalMiner'sDaughter"},{"Age":"74","Index":"55","Movie":"On Golden Pond","Name":"Katharine Hepburn","Year":"1982","normalized_search_column":"55198274KatharineHepburnOnGoldenPond"},{"Age":"33","Index":"56","Movie":"Sophie's Choice","Name":"Meryl Streep","Year":"1983","normalized_search_column":"56198333MerylStreepSophie'sChoice"},{"Age":"49","Index":"57","Movie":"Terms of Endearment","Name":"Shirley MacLaine","Year":"1984","normalized_search_column":"57198449ShirleyMacLaineTermsofEndearment"},{"Age":"38","Index":"58","Movie":"Places in the Heart","Name":"Sally Field","Year":"1985","normalized_search_column":"58198538SallyFieldPlacesintheHeart"},{"Age":"61","Index":"59","Movie":"The Trip to Bountiful","Name":"Geraldine Page","Year":"1986","normalized_search_column":"59198661GeraldinePageTheTriptoBountiful"},{"Age":"21","Index":"60","Movie":"Children of a Lesser God","Name":"Marlee Matlin","Year":"1987","normalized_search_column":"60198721MarleeMatlinChildrenofaLesserGod"},{"Age":"41","Index":"61","Movie":"Moonstruck","Name":"Cher","Year":"1988","normalized_search_column":"61198841CherMoonstruck"},{"Age":"26","Index":"62","Movie":"The Accused","Name":"Jodie Foster","Year":"1989","normalized_search_column":"62198926JodieFosterTheAccused"},{"Age":"80","Index":"63","Movie":"Driving Miss Daisy","Name":"Jessica Tandy","Year":"1990","normalized_search_column":"63199080JessicaTandyDrivingMissDaisy"},{"Age":"42","Index":"64","Movie":"Misery","Name":"Kathy Bates","Year":"1991","normalized_search_column":"64199142KathyBatesMisery"},{"Age":"29","Index":"65","Movie":"The Silence of the Lambs","Name":"Jodie Foster","Year":"1992","normalized_search_column":"65199229JodieFosterTheSilenceoftheLambs"},{"Age":"33","Index":"66","Movie":"Howards End","Name":"Emma Thompson","Year":"1993","normalized_search_column":"66199333EmmaThompsonHowardsEnd"},{"Age":"36","Index":"67","Movie":"The Piano","Name":"Holly Hunter","Year":"1994","normalized_search_column":"67199436HollyHunterThePiano"},{"Age":"45","Index":"68","Movie":"Blue Sky","Name":"Jessica Lange","Year":"1995","normalized_search_column":"68199545JessicaLangeBlueSky"},{"Age":"49","Index":"69","Movie":"Dead Man Walking","Name":"Susan Sarandon","Year":"1996","normalized_search_column":"69199649SusanSarandonDeadManWalking"},{"Age":"39","Index":"70","Movie":"Fargo","Name":"Frances McDormand","Year":"1997","normalized_search_column":"70199739FrancesMcDormandFargo"},{"Age":"34","Index":"71","Movie":"As Good as It Gets","Name":"Helen Hunt","Year":"1998","normalized_search_column":"71199834HelenHuntAsGoodasItGets"},{"Age":"26","Index":"72","Movie":"Shakespeare in Love","Name":"Gwyneth Paltrow","Year":"1999","normalized_search_column":"72199926GwynethPaltrowShakespeareinLove"},{"Age":"25","Index":"73","Movie":"Boys Don't Cry","Name":"Hilary Swank","Year":"2000","normalized_search_column":"73200025HilarySwankBoysDon'tCry"},{"Age":"33","Index":"74","Movie":"Erin Brockovich","Name":"Julia Roberts","Year":"2001","normalized_search_column":"74200133JuliaRobertsErinBrockovich"},{"Age":"35","Index":"75","Movie":"Monster's Ball","Name":"Halle Berry","Year":"2002","normalized_search_column":"75200235HalleBerryMonster'sBall"},{"Age":"35","Index":"76","Movie":"The Hours","Name":"Nicole Kidman","Year":"2003","normalized_search_column":"76200335NicoleKidmanTheHours"},{"Age":"28","Index":"77","Movie":"Monster","Name":"Charlize Theron","Year":"2004","normalized_search_column":"77200428CharlizeTheronMonster"},{"Age":"30","Index":"78","Movie":"Million Dollar Baby","Name":"Hilary Swank","Year":"2005","normalized_search_column":"78200530HilarySwankMillionDollarBaby"},{"Age":"29","Index":"79","Movie":"Walk the Line","Name":"Reese Witherspoon","Year":"2006","normalized_search_column":"79200629ReeseWitherspoonWalktheLine"},{"Age":"61","Index":"80","Movie":"The Queen","Name":"Helen Mirren","Year":"2007","normalized_search_column":"80200761HelenMirrenTheQueen"},{"Age":"32","Index":"81","Movie":"La Vie en rose","Name":"Marion Cotillard","Year":"2008","normalized_search_column":"81200832MarionCotillardLaVieenrose"},{"Age":"33","Index":"82","Movie":"The Reader","Name":"Kate Winslet","Year":"2009","normalized_search_column":"82200933KateWinsletTheReader"},{"Age":"45","Index":"83","Movie":"The Blind Side","Name":"Sandra Bullock","Year":"2010","normalized_search_column":"83201045SandraBullockTheBlindSide"},{"Age":"29","Index":"84","Movie":"Black Swan","Name":"Natalie Portman","Year":"2011","normalized_search_column":"84201129NataliePortmanBlackSwan"},{"Age":"62","Index":"85","Movie":"The Iron Lady","Name":"Meryl Streep","Year":"2012","normalized_search_column":"85201262MerylStreepTheIronLady"},{"Age":"22","Index":"86","Movie":"Silver Linings Playbook","Name":"Jennifer Lawrence","Year":"2013","normalized_search_column":"86201322JenniferLawrenceSilverLiningsPlaybook"},{"Age":"44","Index":"87","Movie":"Blue Jasmine","Name":"Cate Blanchett","Year":"2014","normalized_search_column":"87201444CateBlanchettBlueJasmine"},{"Age":"54","Index":"88","Movie":"Still Alice","Name":"Julianne Moore","Year":"2015","normalized_search_column":"88201554JulianneMooreStillAlice"},{"Age":"26","Index":"89","Movie":"Room","Name":"Brie Larson","Year":"2016","normalized_search_column":"89201626BrieLarsonRoom"}]`
	if h.ItemsJson != should {
		t.Errorf("itemsjson is not correct: \nis:\n%s\nshould be:\n%s", h.ItemsJson, should)
	}
}