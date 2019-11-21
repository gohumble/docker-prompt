package docker

import (
	"github.com/docker/docker/api/types"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	suffixes = [5]string{"b", "KB", "MB", "GB", "TB"}
)
var getSuffix = func(base float64) string {
	return suffixes[int(math.Floor(base))]
}

func convertBytes(bytes int64) string {
	base := math.Log(float64(bytes)) / math.Log(1024)
	getSize := round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	return strconv.FormatFloat(getSize, 'f', 2, 64) + " " + getSuffix(base)
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func formatImageDescription(sum types.ImageSummary) string {
	builder := strings.Builder{}
	builder.WriteString(convertBytes(sum.Size))
	builder.WriteString(" | created at ")
	builder.WriteString(strings.Split(time.Unix(sum.Created, 0).String(), " +")[0])
	builder.WriteString(" | ")
	if sum.RepoTags[0] == "<none>:<none>" {
		builder.WriteString("<none>")
	} else {
		builder.WriteString(sum.RepoTags[0])
	}
	return builder.String()
}
func formatImageId(s string) string {
	id := strings.SplitAfter(s, ":")
	return id[1][:12]
}
