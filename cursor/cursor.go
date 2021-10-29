package cursor

type Cursor interface {
	Get(string) int
	Save() error
	Update(string, int)
}
