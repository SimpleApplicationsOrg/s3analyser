package analyser

import (
	"fmt"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
	"io"
	"math"
	"text/tabwriter"
)

var sizeFormat = map[string]int64{"KB": 1, "MB": 2, "GB": 3, "TB": 4}

type formatLine func(data model.ObjectData) string

// Prints the analyze result
func (r *Result) PrintTo(writer io.Writer) {

	h := bucketHeaderBuilder(r.size)
	formatFunction := r.formatBucket

	if r.byRegion {
		h = regionHeaderBuilder(r.size)
		formatFunction = r.formatRegion
	}

	if r.withStorage {
		h = h.withStorage()
	}

	printResult(writer, h.build(), formatFunction, r)

}

func (r *Result) formatRegion(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%d\t%.f\t%s\t%s\t%s",
		data.Region, data.Count, sizeCalc(data.Size, r.size), data.CreationDate.String(),
		data.LastModified.String(), data.StorageClass)
}

func (r *Result) formatBucket(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%s\t%d\t%.f\t%s\t%s\t%s",
		data.Bucket, data.Region, data.Count, sizeCalc(data.Size, r.size), data.CreationDate.String(),
		data.LastModified.String(), data.StorageClass)
}

func sizeCalc(size int64, format string) float64 {
	pow := sizeFormat[format]
	div := math.Pow(float64(1024), float64(pow))
	return float64(size) / div
}

func printResult(writer io.Writer, header string, format formatLine, result *Result) {
	w := new(tabwriter.Writer)

	w.Init(writer, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintln(w, header)

	for _, object := range result.objects {
		_, _ = fmt.Fprintln(w, format(object))
	}
	_, _ = fmt.Fprintln(w)

	_ = w.Flush()

}
