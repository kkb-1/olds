package jwt

import "context"

const FieldJWTKey = "userId"

func GetAuthToken(accessSecret string, accessExpire int64, id string) (Token, error) {
	opt := TokenOptions{
		AccessSecret: accessSecret,
		AccessExpire: accessExpire,
		Fields: map[string]interface{}{
			FieldJWTKey: id,
		},
	}

	return GenerateToken(opt)
}

func GetAuthValue(ctx context.Context) string {
	var value string
	var ok bool
	if value, ok = ctx.Value(FieldJWTKey).(string); !ok {
		value = ""
	}
	return value
}
