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
func (a *Analyser) Print(writer io.Writer, result *Result) {

	h := bucketHeaderBuilder(a.size)
	formatFunction := a.formatBucket

	if a.byRegion {
		h = regionHeaderBuilder(a.size)
		formatFunction = a.formatRegion
	}

	if a.withStorage {
		h = h.withStorage()
	}

	print(writer, result, h.build(), formatFunction)

}

func (a *Analyser) formatRegion(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%d\t%.f\t%s\t%s\t%s",
		*data.Region, *data.Count, sizeCalc(*data.Size, a.size), data.CreationDate.String(),
		data.LastModified.String(), *data.StorageClass)
}

func (a *Analyser) formatBucket(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%s\t%d\t%.f\t%s\t%s\t%s",
		*data.Bucket, *data.Region, *data.Count, sizeCalc(*data.Size, a.size), data.CreationDate.String(),
		data.LastModified.String(), *data.StorageClass)
}

func sizeCalc(size int64, format string) float64 {
	pow := sizeFormat[format]
	div := math.Pow(float64(1024), float64(pow))
	return float64(size) / div
}

func print(writer io.Writer, result *Result, header string, format formatLine) {
	w := new(tabwriter.Writer)

	w.Init(writer, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, header)

	for _, object := range result.Objects {
		fmt.Fprintln(w, format(*object))
	}
	fmt.Fprintln(w)

	w.Flush()

}
