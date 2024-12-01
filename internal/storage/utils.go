package storage

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func decodePaginationToken(token string) (int, int, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	fmt.Println("decoded", decoded)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid pagination token")
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid pagination token format")
	}

	limit, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid limit in pagination token")
	}

	offset, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid offset in pagination token")
	}

	return limit, offset, nil
}

func encodePaginationToken(limit, offset int) string {
	token := fmt.Sprintf("%d:%d", limit, offset)
	fmt.Println("token", token)
	return base64.StdEncoding.EncodeToString([]byte(token))
}
