package ponysentence

import (
	"fmt"
	"math/rand"
	"time"
)

type verb struct {
	Singular string
	Plural   string
}

type pony struct {
	Name  string
	Photo string
}

var (
	defaultMaxCharacterCount = 3
)

// NewSentence creates a sentence involving one or more ponies, a verb,
// possibly a thing, and a place where it all happens.
func NewSentence(poneAmt int) (s string) {
	s, _ = generate(poneAmt)
	return s
}

// NewSentenceWithImages creates a sentence involving one or more ponies, a
// verb, possibly a thing, and a place where it all happens. It also returns a
// slice of strings containing URLs to images of the poniesand the place the
// sentence mentions.
func NewSentenceWithImages(poneAmt int) (string, []string) {
	return generate(poneAmt)
}

// TODO: NSFW handling

// NewSentenceNSFW creates a sentence involving NewSentence creates a sentence
// involving one or more ponies, a verb, possibly a thing, and a place where it
// all happens. Has the possibility of generating NSFW content.
func NewSentenceNSFW(poneAmt int) (s string) {
	s, _ = generate(poneAmt)
	return s
}

// NewSentenceWithImagesNSFW creates a sentence involving one or more ponies, a
// verb, possibly a thing, and a place where it all happens. It also returns a
// slice of strings containing URLs to images of the poniesand the place the
// sentence mentions. Has the possibility of generating NSFW content.
func NewSentenceWithImagesNSFW(poneAmt int) (string, []string) {
	return generate(poneAmt)
}

func generate(ponyCount int) (string, []string) {
	if ponyCount < 1 {
		return "Amount of ponies must be 1 or more.", nil
	}
	ponyCount -= 1

	// Pinkie Pie and Rarity eat a bagel at a party.
	// Choose between 1, 2, 3 characters
	// Some sort of action - take or takes
	// at {some place/event}

	// characters is a map of character names to their images
	characters := []pony{
		{"Twilight Sparkle", "https://i.imgur.com/bdtfMD2.png"},
		{"Applejack", "https://i.imgur.com/tajlyeh.png"},
		{"Pinkie Pie", "https://i.imgur.com/ZdnWps1.png"},
		{"Rainbow Dash", "https://i.imgur.com/43QQXjQ.png"},
		{"Rarity", "https://i.imgur.com/L2MpiRT.png"},
		{"Fluttershy", "https://i.imgur.com/9zS3vwL.png"},
		{"Spike", "https://i.imgur.com/xW74Eo9.png"},
		{"Princess Celestia", "https://i.imgur.com/XvReGIc.png"},
		{"Princess Luna", "https://i.imgur.com/3UTIlc8.png"},
		{"Princess Cadance", "https://i.imgur.com/5XcEnmn.png"},
		{"Applebloom", "https://i.imgur.com/jTh40CS.png"},
		{"Scootaloo", "https://i.imgur.com/B6juw8H.png"},
		{"Sweetie Belle", "https://i.imgur.com/wIneVOn.png"},
		{"Diamond Tiara", "https://i.imgur.com/aAeMJFk.png"},
		{"Silver Spoon", "https://i.imgur.com/7bAqIH7.png"},
		{"Twist", "https://i.imgur.com/qy0psiU.png"},
		{"Ditzy Doo", "https://i.imgur.com/Qr9VagI.png"},
		{"Featherweight", "https://i.imgur.com/kpkIb47.png"},
		{"Pound Cake", "https://i.imgur.com/OaeDCCi.png"},
		{"Pumpkin Cake", "https://i.imgur.com/oebwJVi.png"},
		{"Snips", "https://i.imgur.com/u7JGqe2.png"},
		{"Snails", "https://i.imgur.com/5VcdRiZ.png"},
		{"Button Mash", "https://i.imgur.com/XS7V0QC.png"},
		{"Twilight Velvet", "https://i.imgur.com/lmXlhfu.png"},
		{"Nightlight", "https://i.imgur.com/pk8iqr0.png"},
		{"Shining Armor", "https://i.imgur.com/l1hM3wn.png"},
		{"Limestone", "https://i.imgur.com/7T6npfv.png"},
		{"Marble", "https://i.imgur.com/s8GE4Ga.png"},
		{"Maud", "https://i.imgur.com/b0LtyK5.png"},
		{"Igneous Rock", "https://i.imgur.com/ntmFfMI.png"},
		{"Cloudy Quartz", "https://i.imgur.com/0YSID6S.png"},
		{"Hondo Flanks", "https://i.imgur.com/tTQOMGr.png"},
		{"Cookie Crumbles", "https://i.imgur.com/cn6Rjc7.png"},
		{"Braeburn", "https://i.imgur.com/W2ka8oD.png"},
		{"Big Macintosh", "https://i.imgur.com/zrGsmOD.png"},
		{"Granny Smith", "https://i.imgur.com/1aD6C88.png"},
		{"Aunt Orange", "https://i.imgur.com/XDIdozQ.png"},
		{"Uncle Orange", "https://i.imgur.com/P29zWK2.png"},
		{"Babs Seed", "https://i.imgur.com/7mTRJYL.png"},
		{"Nightmare Moon", "https://i.imgur.com/gTO3IuI.png"},
		{"Discord", "https://i.imgur.com/wPY7002.png"},
		{"Queen Chrysalis", "https://i.imgur.com/1depFMd.png"},
		{"King Sombra", "https://i.imgur.com/pUmRp32.png"},
		{"Ahuizotl", "https://i.imgur.com/Jkt9nEH.png"},
		{"The Mane-iac", "https://i.imgur.com/J5mN6nJ.png"},
		{"Trixie", "https://i.imgur.com/0XYfiJ5.png"},
		{"Flim", "https://i.imgur.com/X2qNVrn.png"},
		{"Flam", "https://i.imgur.com/KC9qj2J.png"},
		{"Carrot Top", "https://i.imgur.com/bUaURcv.png"},
		{"Bon Bon", "https://i.imgur.com/eJaUjTg.png"},
		{"Lyra Heartstrings", "https://i.imgur.com/p3ySzOE.png"},
		{"Berry Punch", "https://i.imgur.com/XYtkvji.png"},
		{"Cheerilee", "https://i.imgur.com/W0OwZfA.png"},
		{"Caramel", "https://i.imgur.com/Gdl4AEa.png"},
		{"Daisy", "https://i.imgur.com/qD9ZRiU.png"},
		{"Dr. Hooves", "https://i.imgur.com/aO9FvpF.png"},
		{"Roseluck", "https://i.imgur.com/fuEhbGH.png"},
		{"Colegate", "https://i.imgur.com/X5PebHh.png"},
		{"Tank", "https://i.imgur.com/kh02lmH.png"},
		{"Opalescence", "https://i.imgur.com/6cvt6Yg.png"},
		{"Angel Bunny", "https://i.imgur.com/dGcpxd3.png"},
		{"Philomena", "https://i.imgur.com/uDaGuqZ.png"},
		{"Owlowiscious", "https://i.imgur.com/SojErWd.png"},
		{"Winona", "https://i.imgur.com/ubaPpe9.png"},
		{"Gummy", "https://i.imgur.com/IWlc8qc.png"},
		// Pegasus ponies
		{"Derpy Hooves", "https://i.imgur.com/rWhIuTX.png"},
	}

	// Verbs which are to be combined with a noun
	// Noun could be either a thing or a pony
	// Note: If you were to separate sfw and nsfw verbs, have two slices,
	// generate a number using the sum of both lengths, then pick from one
	// depending on whether the number is higher than the first's length.
	nounverbs := []verb{
		{"brush", "brushes"},
		{"arrest", "arrests"},
		{"approve", "approves"},
		{"attack", "attacks"},
		{"admire", "admires"},
		{"alert", "alerts"},
		{"avoid", "avoids"},
		{"annoy", "annoys"},
		{"eat", "eats"},
		{"cover", "covers"},
		{"bruise", "bruises"},
		{"borrow", "borrows"},
		{"bang", "bangs"},
		{"bathe", "bathes"},
		{"battle", "battles"},
		{"bless", "blesses"},
		{"boil", "boils"},
		{"correct", "corrects"},
		{"cure", "cures"},
		{"command", "commands"},
		{"contemplate", "contemplates"},
		{"destroy", "destroys"},
		{"desert", "deserts"},
		{"break up with", "breaks up with"},
		{"dare", "dares"},
		{"disarm", "disarms"},
		{"drown", "drowns"},
		{"find", "finds"},
		{"face", "faces"},
		{"grease", "greases"},
		{"glue", "glues"},
		{"hammer", "hammers"},
		{"judge", "judges"},
		{"juggle", "juggles"},
		{"joke about", "jokes about"},
		{"love", "loves"},
		{"lick", "licks"},
		{"kill", "kills"},
		{"question", "questions"},
		{"remove", "removes"},
		{"stop", "stops"},
		{"soothe", "soothes"},
		{"soak", "soaks"},
		{"support", "supports"},
		{"take", "takes"},
		{"taste", "tastes"},
		{"touch", "touches"},
		{"trip", "trips"},
		{"tease", "teases"},
		{"use", "uses"},
		{"visit", "visits"},
		{"walk around", "walks around"},
		{"wrestle", "wrestles"},
		{"whip", "whips"},
		{"wash", "washes"},
		{"watch", "watches"},
		{"warn", "warns"},
		{"wrap", "wraps"},
	}

	// verbs are a slice of verbs that don't require a noun
	verbs := []verb{
		{"walk around", "walks around"},
		{"beg", "begs"},
		{"attend a show", "attends a show"},
		{"breathe", "breathes"},
		{"cough", "coughs"},
		{"confess", "confesses"},
		{"contemplate life", "contemplates life"},
		{"dress", "dresses"},
		{"crawl", "crawls"},
		{"dance", "dances"},
	}

	things := []string{
		"an apple",
		"a snake",
		"a lizard",
		"a colt",
		"a filly",
		"the sun",
		"some jelly",
		"a quill",
		"a knife",
		"a spider",
		"a bowl of oatmeal",
		"a bag of milk",
		"a bottle of milk",
		"a sofa",
		"a flowerpot",
		"a potato",
		"some spaghetti",
		"a head of lettuce",
		"a lollipop",
		"someone's grandmother",
		"a baby",
		"a train",
		"some grass",
		"a bowl of candy",
		"a hairbrush",
		"a bear",
		"a bunny rabbit",
		"a parasprite",
		"a pear",
		"a bottle of wine",
		"a cactus",
		"a set of pony booties",
		"a scarf",
		"a bike",
		"a bed",
		"a toad",
		"a puzzle",
		"a tree",
		"a Hector",
		"a plushie",
		"some plushies",
		"some clothes",
		"a body pillow",
		"a cake",
		"a bass cannon",
		"the dreaded horse mask",
		"a diary",
		"a treasure chest",
		"farming supplies",
		"a magical cane",
		"a magical crown",
		"some rope",
		"party supplies",
		"some swords and shields",
		"tickets to the gala",
		"a picnic basket",
		"a giant hat",
		"a bus",
		"a packet of crayons",
		"a padlock",
		"a jar of honey",
		"a basket of tomatos",
		"a dinosaur",
		"a stack of pancakes",
		"a gun",
		"a gokart",
		"a ruler",
		"a big sack of bits",
		"a magical pencil",
	}

	// TODO: Show a picture of each place as well
	places := []string{
		"in a prison",
		"at a party",
		"at work",
		"in Sugarcube Corner",
		"at Twilight's Castle",
		"at the School of Friendship",
		"in Cloudsdale",
		"in Rainbow's Parent's House",
		"in the CMC Clubhouse",
		"at Fluttershy's Cottage",
		"at Sweet Apple Acres",
		"in Rarity's Boutique",
		"in the Changling Caverns",
		"in Las Pegasus",
		"within the Wonderbolt's locker room",
		"in Cheerilee's classroom",
		"in the Dragon Lands",
		"in Seaquestria",
		"at Mt. Aris",
		"at the Town of Griffonstone",
		"within the Everfree Forest",
		"in Daring Do's shack",
		"at the Pie Rock Farm",
		"in Appleloosa",
		"in Manehatten",
		"in Ghastly Gorge",
		"in Maud's Cavern",
		"at the Crystal Empire",
		"at Rainbow Falls",
		"in Baltimare",
		"in Our Town",
		"in Somnambula",
		"at Hayseed Swamp",
		"at the train station",
		"at Twilight's Library",
		"in the mountains",
		"in the desert",
		"in Yak-Yakistan",
	}

	// Generate a sentence
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Images to display on screen
	images := []string{}

	// Grab them
	ponies := make([]pony, ponyCount+1, ponyCount+1)
	for i := 0; i <= ponyCount; i++ {
		p := pickRandomPony(r, ponies, characters)
		images = append(images, p.Photo)
		ponies[i] = p
	}

	// Transform the slice into a string
	poniesStr := ""
	switch len(ponies) {
	case 1:
		poniesStr = ponies[0].Name
	case 2:
		poniesStr = fmt.Sprintf("%s and %s", ponies[0].Name, ponies[1].Name)
	default:
		i := 0
		for ; i < len(ponies)-2; i++ {
			poniesStr += fmt.Sprintf("%s, ", ponies[i].Name)
		}
		poniesStr += fmt.Sprintf("%s and %s", ponies[i].Name, ponies[i+1].Name)
	}

	// Use a nounverb (50%), or a regular verb (50%)?
	useNounVerb := r.Intn(2) != 0

	verb := ""
	if useNounVerb {
		// If nounverb, use a pony (20%) or an object (80%)?
		noun := ""
		if r.Intn(10) >= 8 {
			randomPony := pickRandomPony(r, ponies, characters)
			noun = randomPony.Name
			images = append(images, randomPony.Photo)
		} else {
			noun = things[r.Intn(len(things))]
		}

		// Pick a random nounverb
		nounverb := nounverbs[r.Intn(len(nounverbs))]

		// Check if the verb should be plural or not
		if len(ponies) == 1 {
			verb = nounverb.Plural
		} else {
			verb = nounverb.Singular
		}

		// Attach the noun to the end
		verb = fmt.Sprintf("%s %s", verb, noun)
	} else {
		// Just pick a normal verb

		// Check if the verb should be plural or not
		verbObject := verbs[r.Intn(len(verbs))]
		if len(ponies) == 1 {
			verb = verbObject.Plural
		} else {
			verb = verbObject.Singular
		}
	}

	// Pick a random place
	place := places[r.Intn(len(places))]

	// Build the sentence!
	sentence := fmt.Sprintf("%s %s %s.", poniesStr, verb, place)

	return sentence, images
}

// pickRandomPony picks a random pony that's not already in a given list of ponies
func pickRandomPony(r *rand.Rand, chosenPonies, allPonies []pony) pony {
PoneLoop:
	for {
		// Grab a random pony
		p := allPonies[r.Intn(len(allPonies))]

		// See if they're already in the chosen slice
		for _, existingPony := range chosenPonies {
			if existingPony.Name == p.Name {
				// They've already been chosen, pick another pony
				continue PoneLoop
			}
		}

		// Guess they haven't been chosen yet. Return!
		return p
	}
}
