package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	Title string `json:"title"`	
	Post string `json:"post"`
}

type repo struct {
	Items []Item
}


func New() *repo {
	return &repo{
		Items: []Item{},
	}
}

func (r *repo) Add(item Item) {
	r.Items = append(r.Items, item)
}


func (r *repo) GetAll() []Item {
	return r.Items
}