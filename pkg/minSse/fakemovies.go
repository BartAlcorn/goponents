package minSse

import (
	"math/rand"
	"time"
)

var fakemovies = []string{
	"The Shawshank Redemption",
	"The Godfather",
	"The Dark Knight",
	"The Godfather Part II",
	"12 Angry Men",
	"Schindler's List",
	"The Lord of the Rings: The Return of the King",
	"Pulp Fiction",
	"The Lord of the Rings: The Fellowship of the Ring",
	"The Good, the Bad and the Ugly",
	"Forrest Gump",
	"Fight Club",
	"The Lord of the Rings: The Two Towers",
	"Inception",
	"Star Wars: Episode V - The Empire Strikes Back",
	"The Matrix",
	"Goodfellas",
	"One Flew Over the Cuckoo's Nest",
	"Se7en",
	"It's a Wonderful Life",
	"Seven Samurai",
	"Interstellar",
	"The Silence of the Lambs",
	"Saving Private Ryan",
	"City of God",
	"Spider-Man: Across the Spider-Verse",
	"Life Is Beautiful",
	"The Green Mile",
	"Star Wars: Episode IV - A New Hope",
	"Terminator 2: Judgment Day",
}

func FakeMovie() string {
	rand.Seed(time.Now().UnixNano())
	max := len(fakemovies)
	n := rand.Intn(max)

	return fakemovies[n]
}
