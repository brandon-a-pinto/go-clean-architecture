package protocol

type IJWTAdapter interface {
	Generate(id string) (string, error)
}
