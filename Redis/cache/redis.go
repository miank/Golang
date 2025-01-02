package cache

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MovieService interface {
	GetMovie(id string) (*Movie, error)
	GetMovies([]*Movie, error)
	CreateMovie(movie *Movie) (*Movie, error)
	UpdateMovie(movie *Movie) (*Movie, error)
	DeleteMovie(id string) error
}
