package imdb

import "testing"

func TestTitle(t *testing.T) {
	_, err := NewTitle(client, "wrong")
	if err != ErrInvalidID {
		t.Errorf("NewTitle(wrong) = %v; want ErrInvalidId", err)
	}

	for _, tt := range []struct {
		ID   string
		want Title
	}{
		{
			ID: "tt0073845",
			want: Title{
				ID:          "tt0073845",
				URL:         "https://www.imdb.com/title/tt0073845",
				Name:        "L'uomo che sfidò l'organizzazione",
				Type:        "Movie",
				Year:        1975,
				Rating:      "5.3",
				RatingCount: 25,
				Duration:    "1h27m",
				Directors: []Name{
					Name{ID: "nm0340894", URL: "https://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
				},
				Writers: []Name{
					Name{ID: "nm0340894", URL: "https://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
					Name{ID: "nm0739308", URL: "https://www.imdb.com/name/nm0739308", FullName: "Rafael Romero Marchent"},
				},
				Actors: []Name{
					Name{ID: "nm0001886", URL: "https://www.imdb.com/name/nm0001886", FullName: "Howard Ross"},
					Name{ID: "nm0775810", URL: "https://www.imdb.com/name/nm0775810", FullName: "Karin Schubert"},
					Name{ID: "nm0000963", URL: "https://www.imdb.com/name/nm0000963", FullName: "Stephen Boyd"},
				},
				Genres:        []string{"Crime", "Drama"},
				Languages:     []string{"Italian"},
				Nationalities: []string{"Italy", "France", "Spain"},
				Description:   "An airport employee switches a pack of drugs for baking soda and absconds to Barcelona, meanwhile the drug-runners are on his trail.",
				Poster:        Media{ID: "rm3600132097", TitleID: "tt0073845", URL: "https://www.imdb.com/title/tt0073845/mediaviewer/rm3600132097", ContentURL: "https://m.media-amazon.com/images/M/MV5BMzJmM2E0ZTEtYmJhMy00Nzc2LTgyMDUtODViNzQyZDczODc0XkEyXkFqcGdeQXVyMTQ3Njg3MQ@@._V1_.jpg"},
			},
		},
		{
			ID: "tt0437804",
			want: Title{
				ID:          "tt0437803", // ID redirect.
				URL:         "https://www.imdb.com/title/tt0437803",
				Name:        "Alien Siege",
				Type:        "Movie",
				Year:        2005,
				Rating:      "3.5",
				RatingCount: 799,
				Duration:    "1h30m",
				Directors: []Name{
					Name{ID: "nm0821100", URL: "https://www.imdb.com/name/nm0821100", FullName: "Robert Stadd"},
				},
				Writers: []Name{
					Name{ID: "nm1305705", URL: "https://www.imdb.com/name/nm1305705", FullName: "Bill Lundy"},
					Name{ID: "nm0757507", URL: "https://www.imdb.com/name/nm0757507", FullName: "Paul Salamoff"},
					Name{ID: "nm0821100", URL: "https://www.imdb.com/name/nm0821100", FullName: "Robert Stadd"},
				},
				Actors: []Name{
					Name{ID: "nm0424635", URL: "https://www.imdb.com/name/nm0424635", FullName: "Brad Johnson"},
					Name{ID: "nm1658541", URL: "https://www.imdb.com/name/nm1658541", FullName: "Erin Ross"},
					Name{ID: "nm0250169", URL: "https://www.imdb.com/name/nm0250169", FullName: "Lilas Lane"},
				},
				Genres:        []string{"Action", "Adventure", "Drama"},
				Languages:     []string{"English"},
				Nationalities: []string{"United States"},
				Description:   "An advanced race of aliens descends upon Earth with a single goal - the blood of eight million humans to save their own dying planet.",
				Poster:        Media{ID: "rm804357633", TitleID: "tt0437803", URL: "https://www.imdb.com/title/tt0437803/mediaviewer/rm804357633", ContentURL: "https://m.media-amazon.com/images/M/MV5BZTEwZDBmNTgtMTg3OC00MDJiLTkzZTgtZGYwYzcwMWMwZTBjXkEyXkFqcGdeQXVyMTY5Nzc4MDY@._V1_.jpg"},
			},
		},
		{
			ID: "tt0290988",
			want: Title{
				Name:          "Trailer Park Boys",
				SeasonCount:   12,
				Duration:      "30m",
				RatingCount:   51412,
				Rating:        "8.5",
				Year:          2004,
				ID:            "tt0290988",
				Type:          "TVSeries",
				URL:           "https://www.imdb.com/title/tt0290988",
				Writers:       []Name{Name{ID: "nm0154104", URL: "https://www.imdb.com/name/nm0154104", FullName: "Mike Clattenburg"}},
				Actors:        []Name{Name{ID: "nm1033215", URL: "https://www.imdb.com/name/nm1033215", FullName: "John Paul Tremblay"}, Name{ID: "nm1036211", URL: "https://www.imdb.com/name/nm1036211", FullName: "Robb Wells"}, Name{ID: "nm1034266", URL: "https://www.imdb.com/name/nm1034266", FullName: "Mike Smith"}},
				Genres:        []string{"Comedy", "Crime"},
				Languages:     []string{"English"},
				Nationalities: []string{"Canada"},
				Description:   "Three petty felons have a documentary made about their life in a trailer park.",
				Poster:        Media{ID: "rm2293178112", TitleID: "tt0290988", URL: "https://www.imdb.com/title/tt0290988/mediaviewer/rm2293178112", ContentURL: "https://m.media-amazon.com/images/M/MV5BOTA0NTAwMTc1MF5BMl5BanBnXkFtZTgwODk2NjY0ODE@._V1_.jpg"},
			},
		},
		{
			ID: "tt1179034",
			want: Title{
				ID:          "tt1179034",
				URL:         "https://www.imdb.com/title/tt1179034",
				Name:        "From Paris with Love",
				Type:        "Movie",
				Year:        2010,
				Rating:      "6.4",
				RatingCount: 107802,
				Duration:    "1h32m",
				Directors: []Name{
					Name{ID: "nm0603628", URL: "https://www.imdb.com/name/nm0603628", FullName: "Pierre Morel"},
				},
				Writers: []Name{
					Name{ID: "nm0367867", URL: "https://www.imdb.com/name/nm0367867", FullName: "Adi Hasak"},
					Name{ID: "nm0000108", URL: "https://www.imdb.com/name/nm0000108", FullName: "Luc Besson"},
				},
				Actors: []Name{
					Name{ID: "nm0000237", URL: "https://www.imdb.com/name/nm0000237", FullName: "John Travolta"},
					Name{ID: "nm0001667", URL: "https://www.imdb.com/name/nm0001667", FullName: "Jonathan Rhys Meyers"},
					Name{ID: "nm0810738", URL: "https://www.imdb.com/name/nm0810738", FullName: "Kasia Smutniak"},
				},
				Genres:        []string{"Action", "Crime", "Thriller"},
				Languages:     []string{"English", "French", "Mandarin", "German"},
				Nationalities: []string{"France", "United Kingdom", "United States"},
				Description:   "In Paris, a young employee in the office of the US Ambassador hooks up with an American spy looking to stop a terrorist attack in the city.",
				Poster:        Media{ID: "rm30481408", TitleID: "tt1179034", URL: "https://www.imdb.com/title/tt1179034/mediaviewer/rm30481408", ContentURL: "https://m.media-amazon.com/images/M/MV5BODAwMDFjNjktMWY2Mi00MmVhLWI0MjYtNzg4OTI0NzA5YzBjXkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_.jpg"},
			},
		},
		{
			ID: "tt0133093",
			want: Title{
				ID:          "tt0133093",
				URL:         "https://www.imdb.com/title/tt0133093",
				Name:        "The Matrix",
				Type:        "Movie",
				Year:        1999,
				Rating:      "8.7",
				RatingCount: 1536356,
				Duration:    "2h16m",
				Directors: []Name{
					Name{ID: "nm0905154", URL: "https://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
					Name{ID: "nm0905152", URL: "https://www.imdb.com/name/nm0905152", FullName: "Lilly Wachowski"},
				},
				Writers: []Name{
					Name{ID: "nm0905152", URL: "https://www.imdb.com/name/nm0905152", FullName: "Lilly Wachowski"},
					Name{ID: "nm0905154", URL: "https://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
				},
				Actors: []Name{
					Name{ID: "nm0000206", URL: "https://www.imdb.com/name/nm0000206", FullName: "Keanu Reeves"},
					Name{ID: "nm0000401", URL: "https://www.imdb.com/name/nm0000401", FullName: "Laurence Fishburne"},
					Name{ID: "nm0005251", URL: "https://www.imdb.com/name/nm0005251", FullName: "Carrie-Anne Moss"},
				},
				Genres:        []string{"Action", "Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"United States", "Australia"},
				Description:   "When a beautiful stranger leads computer hacker Neo to a forbidding underworld, he discovers the shocking truth--the life he knows is the elaborate deception of an evil cyber-intelligence.",
				Poster:        Media{ID: "rm525547776", TitleID: "tt0133093", URL: "https://www.imdb.com/title/tt0133093/mediaviewer/rm525547776", ContentURL: "https://m.media-amazon.com/images/M/MV5BNzQzOTk3OTAtNDQ0Zi00ZTVkLWI0MTEtMDllZjNkYzNjNTc4L2ltYWdlXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_.jpg"},
			},
		},
		{
			ID: "tt21831910",
			want: Title{
				ID:            "tt21831910",
				URL:           "https://www.imdb.com/title/tt21831910",
				Name:          "The Decameron",
				Type:          "TVSeries",
				Year:          2024,
				RatingCount:   4319,
				SeasonCount:   1,
				Writers:       []Name{Name{ID: "nm6767020", URL: "https://www.imdb.com/name/nm6767020", FullName: "Kathleen Jordan"}},
				Actors:        []Name{Name{ID: "nm0149447", URL: "https://www.imdb.com/name/nm0149447", FullName: "Amar Chadha-Patel"}, Name{ID: "nm3355536", URL: "https://www.imdb.com/name/nm3355536", FullName: "Leila Farzad"}, Name{ID: "nm6103482", URL: "https://www.imdb.com/name/nm6103482", FullName: "Lou Gala"}},
				Rating:        "6.3",
				Genres:        []string{"Comedy", "Drama", "History"},
				Nationalities: []string{"United States"},
				Languages:     []string{"English"},
				Poster:        Media{ID: "rm1531016193", TitleID: "tt21831910", URL: "https://www.imdb.com/title/tt21831910/mediaviewer/rm1531016193", ContentURL: "https://m.media-amazon.com/images/M/MV5BNTI2YjliODQtOGU2OS00ZWI0LWFiZmMtNzdmNWZhN2FhNWFmXkEyXkFqcGc@._V1_.jpg"},
				Description:   "Set in 1348 during the Black Death, the deadliest pandemic in human history, explores the timely themes of class systems, power struggles, and survival in a time of pandemic.",
			},
		},
		{
			ID: "tt0291830",
			want: Title{
				ID:       "tt0291830",
				URL:      "https://www.imdb.com/title/tt0291830",
				Name:     "Corps à coeur",
				Type:     "Movie",
				Year:     1981,
				Duration: "26m",
				Directors: []Name{
					Name{ID: "nm1029036", URL: "https://www.imdb.com/name/nm1029036", FullName: "Jean-François Després"},
					Name{ID: "nm1003289", URL: "https://www.imdb.com/name/nm1003289", FullName: "Alain Godon"},
				},
				Genres:        []string{"Documentary", "Short"},
				Nationalities: []string{"Canada"},
			},
		},
		{
			ID: "tt1965639",
			want: Title{
				ID:            "tt1965639",
				URL:           "https://www.imdb.com/title/tt1965639",
				Name:          "El clima y las cuatro estaciones",
				Type:          "TVSeries",
				Year:          1994,
				Genres:        []string{"Documentary"},
				Languages:     []string{"Spanish"},
				Nationalities: []string{"Spain"},
			},
		},
		{
			ID: "tt0086677",
			want: Title{
				ID:          "tt0086677",
				URL:         "https://www.imdb.com/title/tt0086677",
				Name:        "Brothers",
				Type:        "TVSeries",
				Year:        1984,
				SeasonCount: 5,
				Rating:      "8.0",
				RatingCount: 267,
				Duration:    "22m",
				Writers: []Name{
					Name{ID: "nm0515953", URL: "https://www.imdb.com/name/nm0515953", FullName: "David Lloyd"},
					Name{ID: "nm0031246", URL: "https://www.imdb.com/name/nm0031246", FullName: "Greg Antonacci"},
				},
				Actors: []Name{
					Name{ID: "nm0907107", URL: "https://www.imdb.com/name/nm0907107", FullName: "Robert Walden"},
					Name{ID: "nm0716618", URL: "https://www.imdb.com/name/nm0716618", FullName: "Paul Regina"},
					Name{ID: "nm0535933", URL: "https://www.imdb.com/name/nm0535933", FullName: "Brandon Maggart"},
				},
				Genres:        []string{"Comedy"},
				Languages:     []string{"English"},
				Nationalities: []string{"United States"},
				Description:   "When their youngest brother comes out as gay, two conservative men support him and help him navigate being openly gay in 1980s Philadelphia.",
				Poster:        Media{ID: "rm1328617472", TitleID: "tt0086677", URL: "https://www.imdb.com/title/tt0086677/mediaviewer/rm1328617472", ContentURL: "https://m.media-amazon.com/images/M/MV5BNDM3N2IxY2UtMDc4MS00ZDM5LWJjODUtMjQ5ODg2YWIxMTE4XkEyXkFqcGdeQXVyMDYxMTUwNg@@._V1_.jpg"},
			},
		},
		{
			ID: "tt1371159",
			want: Title{
				ID:          "tt1371159",
				URL:         "https://www.imdb.com/title/tt1371159",
				Name:        "Iron Man 2",
				Type:        "VideoGame",
				Year:        2010,
				Rating:      "6.1",
				RatingCount: 641,
				Directors: []Name{
					Name{ID: "nm4157448", URL: "https://www.imdb.com/name/nm4157448", FullName: "Michael McCormick"},
					Name{ID: "nm4157551", URL: "https://www.imdb.com/name/nm4157551", FullName: "Robert Taylor"},
				},
				Writers: []Name{
					Name{ID: "nm3881781", URL: "https://www.imdb.com/name/nm3881781", FullName: "Matt Fraction"},
					Name{ID: "nm1411347", URL: "https://www.imdb.com/name/nm1411347", FullName: "Don Heck"},
				},
				Actors: []Name{
					Name{ID: "nm0000332", URL: "https://www.imdb.com/name/nm0000332", FullName: "Don Cheadle"},
					Name{ID: "nm0519680", URL: "https://www.imdb.com/name/nm0519680", FullName: "Eric Loomis"},
					Name{ID: "nm0000168", URL: "https://www.imdb.com/name/nm0000168", FullName: "Samuel L. Jackson"},
				},
				Genres:        []string{"Action", "Adventure", "Crime"},
				Languages:     []string{"English"},
				Nationalities: []string{"United States"},
				Description:   "Video game based upon the film of the same name.",
				Poster:        Media{ID: "rm2383169025", TitleID: "tt1371159", URL: "https://www.imdb.com/title/tt1371159/mediaviewer/rm2383169025", ContentURL: "https://m.media-amazon.com/images/M/MV5BYTMzOTJmNDgtNmE4My00MzQ5LTg4YmMtOWRmNjJiMzUzZTg0XkEyXkFqcGdeQXVyMTA0MTM5NjI2._V1_.jpg"},
			},
		},
		{
			ID: "tt0423866",
			want: Title{
				ID:          "tt0423866",
				URL:         "https://www.imdb.com/title/tt0423866",
				Name:        "Bin-jip",
				Type:        "Movie",
				Year:        2004,
				Rating:      "7.9",
				RatingCount: 45493,
				Duration:    "1h28m",
				Directors: []Name{
					Name{ID: "nm1104118", URL: "https://www.imdb.com/name/nm1104118", FullName: "Kim Ki-duk"},
				},
				Writers: []Name{
					Name{ID: "nm15621750", URL: "https://www.imdb.com/name/nm15621750", FullName: "Kim Ki-Duk"},
					Name{ID: "nm1104118", URL: "https://www.imdb.com/name/nm1104118", FullName: "Kim Ki-duk"},
				},
				Actors: []Name{
					{ID: "nm1165901", URL: "https://www.imdb.com/name/nm1165901", FullName: "Lee Seung-yun"},
					{ID: "nm1030819", URL: "https://www.imdb.com/name/nm1030819", FullName: "Jae Hee"},
					{ID: "nm1891528", URL: "https://www.imdb.com/name/nm1891528", FullName: "Hyuk-ho Kwon"},
				},
				Genres:        []string{"Crime", "Drama", "Romance"},
				Languages:     []string{"Korean"},
				Nationalities: []string{"South Korea", "Japan"},
				Description:   "A transient young man breaks into empty homes to partake of the vacationing residents' lives for a few days.",
				Poster:        Media{ID: "rm880057600", TitleID: "tt0423866", URL: "https://www.imdb.com/title/tt0423866/mediaviewer/rm880057600", ContentURL: "https://m.media-amazon.com/images/M/MV5BMTM1ODIwNzM5OV5BMl5BanBnXkFtZTcwNjk5MDkyMQ@@._V1_.jpg"},
			},
		},
		{
			ID: "tt0900059",
			want: Title{
				ID:       "tt0900059",
				URL:      "https://www.imdb.com/title/tt0900059",
				Name:     "Mayans",
				ParentID: "tt0883792",
				Type:     "TVEpisode",
				Year:     2002,
				Writers: []Name{
					Name{ID: "nm2422546", URL: "https://www.imdb.com/name/nm2422546", FullName: "Nick Greenaway"},
					Name{ID: "nm0646855", URL: "https://www.imdb.com/name/nm0646855", FullName: "Harley Oliver"},
				},
				Actors: []Name{
					Name{ID: "nm2426039", URL: "https://www.imdb.com/name/nm2426039", FullName: "Jo Rush"},
					Name{ID: "nm2423121", URL: "https://www.imdb.com/name/nm2423121", FullName: "Matt Tomaszewski"},
				},
				Genres: []string{"Documentary"},
			},
		},
		{
			ID: "tt0732901",
			want: Title{
				ID:            "tt0732901",
				URL:           "https://www.imdb.com/title/tt0732901",
				Name:          "Closer to the Heart",
				Type:          "TVEpisode",
				ParentID:      "tt0290988",
				Year:          2003,
				Rating:        "9.1",
				RatingCount:   916,
				Duration:      "23m",
				Directors:     []Name{{ID: "nm0154104", URL: "https://www.imdb.com/name/nm0154104", FullName: "Mike Clattenburg"}},
				Writers:       []Name{Name{ID: "nm0154104", URL: "https://www.imdb.com/name/nm0154104", FullName: "Mike Clattenburg"}},
				Actors:        []Name{Name{ID: "nm1033215", URL: "https://www.imdb.com/name/nm1033215", FullName: "John Paul Tremblay"}, Name{ID: "nm1036211", URL: "https://www.imdb.com/name/nm1036211", FullName: "Robb Wells"}, Name{ID: "nm0243082", URL: "https://www.imdb.com/name/nm0243082", FullName: "John Dunsworth"}},
				Genres:        []string{"Comedy", "Crime"},
				Languages:     []string{"English"},
				Nationalities: []string{"Canada"},
				Description:   "When Randy and Lahey rip the boys off on buying Rush tickets, Ricky \"borrows\" guitarist Alex Lifeson for access to the concert.",
				Poster:        Media{ID: "rm1259893761", TitleID: "tt0732901", URL: "https://www.imdb.com/title/tt0732901/mediaviewer/rm1259893761", ContentURL: "https://m.media-amazon.com/images/M/MV5BNjY4YjJhZDQtM2U3ZS00ZThkLWIyZmMtZTY3MmU3N2QwOTVmXkEyXkFqcGc@._V1_.jpg"},
				Season:        3,
				Episode:       5,
			},
		},
		{
			ID: "tt6856242",
			want: Title{
				ID:          "tt6856242",
				URL:         "https://www.imdb.com/title/tt6856242",
				Name:        "The King's Man",
				Type:        "Movie",
				Year:        2021,
				Rating:      "6.3",
				RatingCount: 115757,
				Duration:    "2h11m",
				Directors: []Name{
					Name{ID: "nm0891216", URL: "https://www.imdb.com/name/nm0891216", FullName: "Matthew Vaughn"},
				},
				Writers: []Name{
					Name{ID: "nm0891216", URL: "https://www.imdb.com/name/nm0891216", FullName: "Matthew Vaughn"},
					Name{ID: "nm2244980", URL: "https://www.imdb.com/name/nm2244980", FullName: "Karl Gajdusek"},
					Name{ID: "nm2092839", URL: "https://www.imdb.com/name/nm2092839", FullName: "Mark Millar"},
				},
				Actors: []Name{
					Name{ID: "nm0000146", URL: "https://www.imdb.com/name/nm0000146", FullName: "Ralph Fiennes"},
					Name{ID: "nm2605345", URL: "https://www.imdb.com/name/nm2605345", FullName: "Gemma Arterton"},
					Name{ID: "nm0406975", URL: "https://www.imdb.com/name/nm0406975", FullName: "Rhys Ifans"},
				},
				Genres:        []string{"Action", "Adventure", "Thriller"},
				Languages:     []string{"English", "Latin", "German", "French", "Russian"},
				Nationalities: []string{"United Kingdom", "United States"},
				Description:   "In the early years of the 20th century, the Kingsman agency is formed to stand against a cabal plotting a war to wipe out millions.",
				Poster:        Media{ID: "rm1942154753", TitleID: "tt6856242", URL: "https://www.imdb.com/title/tt6856242/mediaviewer/rm1942154753", ContentURL: "https://m.media-amazon.com/images/M/MV5BMDEzZDY2ZDktNTlmOS00NThjLThkNTEtMjE5MzI5NWEwZmRjXkEyXkFqcGdeQXVyMDA4NzMyOA@@._V1_.jpg"},
			},
		},
	} {
		got, err := NewTitle(client, tt.ID)
		if err != nil {
			t.Errorf("NewTitle(%s) error: %v", tt.ID, err)
		} else {
			if got.RatingCount > tt.want.RatingCount && tt.want.RatingCount > 0 {
				tt.want.RatingCount = got.RatingCount
			}
			if err := diffStruct(*got, tt.want); err != nil {
				t.Errorf("NewTitle(%s): %v", tt.ID, err)
			}
		}
	}
}
