package log

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"roofix/config"
	"roofix/pkg/enum"
	"roofix/pkg/util/ptr"
	"roofix/pkg/util/storage"
	"time"
)

func dump(caller, error, stack string) {
	t := time.Now().UTC()
	hash := hash(caller)
	key := fmt.Sprintf("%s/%s/%s/%s.json", enum.DirErrorLogs, t.Format(time.DateOnly), hash, t.Format("15-04-05.000000000"))

	data, err := json.Marshal(map[string]string{
		"timeStamp": t.Format(time.RFC3339),
		"caller":    caller,
		"error":     error,
		"stack":     stack,
	})
	if err != nil {
		fmt.Println("** log.dump marshal error")
		fmt.Println(err.Error())
		fmt.Println("**")
		return
	}

	err = storage.PutObject(context.Background(), config.LogBucket(), key, data, ptr.Str("application/json"))
	if err != nil {
		fmt.Println("** log.dump storage.PutObject error")
		fmt.Println(err.Error())
		fmt.Println("**")
		return
	}
}

func hash(txt string) string {
	hasher := md5.New()
	hasher.Write([]byte(txt))
	return hex.EncodeToString(hasher.Sum(nil))
}
