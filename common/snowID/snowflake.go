package snowID

import "github.com/bwmarrin/snowflake"

func GetID(node int) (string, error) {
	newNode, err := snowflake.NewNode(int64(node))
	if err != nil {
		return "", err
	}

	id := newNode.Generate()
	return id.String(), err
}
