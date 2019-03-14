package pet

type Category struct {
	ID   int64
	Name string
}

type Tag struct {
	ID   int64
	Name string
}

type Pet struct {
	ID        int64
	Category  Category
	Name      string
	PhotoURLs []map[string]interface{}
	Tags      []Tag
	Status    string
}
