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
		{"Twilight Sparkle", ""},
		{"Applejack", ""},
		{"Pinkie Pie", ""},
		{"Rainbow Dash", ""},
		{"Rarity", ""},
		{"Fluttershy", ""},
		{"Spike", ""},
		{"Princess Celestia", ""},
		{"Princess Luna", ""},
		{"Princess Cadance", ""},
		{"Applebloom", ""},
		{"Scootaloo", ""},
		{"Sweetie Belle", ""},
		{"Diamond Tiara", ""},
		{"Silver Spoon", ""},
		{"Twist", ""},
		{"Ditzy Doo", ""},
		{"Featherweight", ""},
		{"Pound Cake", ""},
		{"Pumpkin Cake", ""},
		{"Snips", ""},
		{"Snails", ""},
		{"Button Mash", ""},
		{"Star Sparkle", ""},
		{"Nightlight", ""},
		{"Shining Armor", ""},
		{"Limestone", ""},
		{"Marble", ""},
		{"Maud", ""},
		{"Igneous Rock", ""},
		{"Cloudy Quartz", ""},
		{"Rarity's Dad", ""},
		{"Rarity's Mom", ""},
		{"Braeburn", ""},
		{"Big Macintosh", ""},
		{"Granny Smith", ""},
		{"Aunt Orange", ""},
		{"Uncle Orange", ""},
		{"Babs Seed", ""},
		{"Nightmare Moon", ""},
		{"Discord", ""},
		{"Queen Chrysalis", ""},
		{"King Sombra", ""},
		{"Ahuizotl", ""},
		{"The Mane-iac", ""},
		{"Trixie", ""},
		{"Flim", ""},
		{"Flam", ""},
		{"Carrot Top", ""},
		{"Bon Bon", ""},
		{"Lyra Heartstrings", ""},
		{"Berry Punch", ""},
		{"Cheerilee", ""},
		{"Caramel", ""},
		{"Daisy", ""},
		{"Dr. Whooves", ""},
		{"Roseluck", ""},
		{"Colegate", ""},
		{"Tank", ""},
		{"Opalescence", ""},
		{"Angel Bunny", ""},
		{"Philomena", ""},
		{"Owlowiscious", ""},
		{"Winona", ""},
		{"Gummy", ""},
		// Pegasus ponies
		{"Derpy Hooves", ""},
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
		{"fuck", "fucks"},
		{"grease", "greases"},
		{"glue", "glues"},
		{"hammer", "hammers"},
		{"judge", "judges"},
		{"juggle", "juggles"},
		{"joke about", "jokes about"},
		{"love", "loves"},
		{"lick", "licks"},
		{"kill", "kills"},
		{"milk", "milks"},
		{"paddle", "paddles"},
		{"question", "questions"},
		{"rub", "rubs"},
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
		{"bomb", "bombs"},
		{"get baked", "gets baked"},
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
		"a slave",
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
		"sex toys",
		"ballgags and whips",
		"a magical cane",
		"a magical crown",
		"some rope",
		"party supplies",
		"a dildo",
		"some swords and shields",
		"tickets to the gala",
		"a maid outfit",
		"a picnic basket",
		"a giant hat",
		"a bus",
		"a severed pony head",
		"a packet of crayons",
		"a padlock",
		"a jar of honey",
		"a collar",
		"a basket of tomatos",
		"a dinosaur",
		"a stack of pancakes",
		"a gun",
		"a gokart",
		"a butt plug",
		"a chastity cage",
		"a ruler",
		"a big sack of bits",
		"Spike's secret porn collection",
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
		"in The CMC Clubhouse",
		"at The Twin Towers",
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
		"deep within Pinkie Pie's Party Cave",
		"an orgy",
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
		"at Somnambula",
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
	for ; ponyCount >= 0; ponyCount-- {
		p := pickRandomPony(r, ponies, characters)
		images = append(images, p.Photo)
		ponies[ponyCount] = p
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
			noun = pickRandomPony(r, ponies, characters).Name
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
