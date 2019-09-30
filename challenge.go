package challenge

// gophers is a list of GolangKazan GolangConf-2019 participants.
//
// Used in tests (see challenge_test.go).
var gophers = []gopher{
	// Participants (add yourself below):
	{
		name: "Alexander Kiryukhin",
		id:   "neonxp",
		post: "https://vk.com/feed?w=wall476865374_157",
	},
	{
		name: "Илья Шихалеев",
		id:   "ilya-shikhaleev",
		post: "https://vk.com/wall4220274_3221",
	},
	{
		name: "Sharapov Arkadiy",
		id:   "avsharapov",
		post: "https://www.instagram.com/p/B29dy8Pg9BT/",
	},
	{
		name: "Nikita Vanyasin",
		id:   "nikita-vanyasin",
		post: "https://www.facebook.com/nikita.vanyasin/posts/1312245865615823",
	},
	{
		name: "Delius Farkhullin",
		id:   "bontequero",
		post: "https://vk.com/bontequero?w=wall140207068_146",
	},
	{
		name: "Ilya Zaharov",
		id:   "zoh",
		post: "https://www.instagram.com/p/B2_vmNIB3At/",
	},
	{
		name: "Pavel Sukhodoev",
		id:   "xcyxoux",
		post: "https://twitter.com/GolangKazan/status/1177939214437994496", // Retweet
	},
	{
		name: "Vadim Gorbunov",
		id:   "vadimgorbunov",
		post: "https://twitter.com/vadimgorbunov/status/1178323050518974468",
	},
	{
		name: "Ivan Ryazanov",
		id:   "ivantaran",
		post: "https://www.instagram.com/p/B3AOS9TFtSI/",
	},
	{
		name: "Nur Kutlugallyamov",
		id:   "milQA",
		post: "https://vk.com/wall302234537_938",
	},
	{
		name: "Zufar Samigullin",
		id:   "profbis",
		post: "https://vk.com/profbis?w=wall13681269_3829",
	},
	{
		name: "Alik Khilazhev",
		id:   "alikhil",
		post: "https://vk.com/alikhil?w=wall122186426_669",
	},
	{
		name: "Rinat Baygildin",
		id:   "bayrinat",
		post: "https://www.instagram.com/p/B3B2PxAoXmf/?igshid=1o0caviaryfs7",
	},
	{
		name: "Ivan Nikiforov",
		id:   "niki4",
		post: "https://twitter.com/_niki4_/status/1178589754025238528",
	},
	{
		name: "Dmitry Merkushin",
		id:   "merkushin",
		post: "https://twitter.com/merkushin/status/1178595009140068352",
	},

	// Participants from Nizhny Novgorod (https://vk.com/golang_nizhny):
	{
		name:           "Aleksey Ryabchikov",
		id:             "aryabchi",
		post:           "https://vk.com/id27102990?w=wall27102990_3493",
		nizhnyNovgorod: true,
	},
	{
		name:           "Martynenko Andrey",
		id:             "lergor11",
		post:           "https://twitter.com/lergor11/status/1178569646473175040",
		nizhnyNovgorod: true,
	},
	{
		name:           "Daria Kolistratova",
		id:             "dakolistratova",
		post:           "https://vk.com/dakolistratova?w=wall18700317_3332",
		nizhnyNovgorod: true,
	},
	{
		name:           "Victor Ryabinin",
		id:             "mmvds",
		post:           "https://vk.com/vitek_ru?w=wall1966944_986",
		nizhnyNovgorod: true,
	},

	// Testers. Do not need to contain a valid URL in post,
	// but should use `tester: true` and can't win a prize.
	{
		name:   "Iskander Sharipov",
		id:     "quasilyte",
		post:   "https://todo.add.a.post/123",
		tester: true,
	},
	{
		name:   "Oleg Kovalov",
		id:     "cristaloleg",
		post:   "https://prodam.garaz",
		tester: true,
	},
}

// gopher is a struct that describes GolangKazan GolangConf-2019 challenge participant.
type gopher struct {
	// name is participant full name (first + second).
	name string

	// id is a secondary identifier that is used if name alone can't
	// identify the participant. Optional, can be left empty.
	id string

	// post is a link to a media post about this challenge.
	//
	// We accept any kind of posts, here are some examples:
	//	- VK or Facebook post on a wall
	//	- Tweet
	//	- Instagram post
	//	- Etc, etc. (but read rules below beforehand)
	//
	// Rules:
	//	1. Post must include both #GolangKazan #GolangConf2019 hash tags.
	//	2. Post must survive until we stop the challenge and announce the winner.
	//	3. Account that does the announcement must have at least 10 followers/subscribers.
	post string

	// tester can submit invalid post string, but can't win a prize.
	tester bool

	// nizhnyNovgorod tells whether a participant is a part of Nizhny Novgorod
	// Go community.
	nizhnyNovgorod bool
}

// key returns a hopefully unique participant key.
func (g gopher) key() string {
	if g.id == "" {
		return g.name
	}
	return g.name + "/" + g.id
}
