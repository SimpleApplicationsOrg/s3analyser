package analyser

import (
	"fmt"
	"github.com/SimpleApplicationsOrg/s3analyser/model"
	"math"
	"os"
	"text/tabwriter"
)

var sizeFormat = map[string]int64{"KB": 1, "MB": 2, "GB": 3, "TB": 4}

type formatLine func(data model.ObjectData) string

// Prints the analyze result
func (sat *sat) Print(result *Result) {

	h := bucketHeaderBuilder(sat.size)
	formatFunction := sat.formatBucket

	if sat.byRegion {
		h = regionHeaderBuilder(sat.size)
		formatFunction = sat.formatRegion
	}

	if sat.withStorage {
		h = h.withStorage()
	}

	print(result, h.build(), formatFunction)

}

func (sat *sat) formatRegion(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%d\t%.f\t%s\t%s\t%s",
		*data.Region, *data.Count, sizeCalc(*data.Size, sat.size), data.CreationDate.String(),
		data.LastModified.String(), *data.StorageClass)
}

func (sat *sat) formatBucket(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%s\t%d\t%.f\t%s\t%s\t%s",
		*data.Bucket, *data.Region, *data.Count, sizeCalc(*data.Size, sat.size), data.CreationDate.String(),
		data.LastModified.String(), *data.StorageClass)
}

func sizeCalc(size int64, format string) float64 {
	pow := sizeFormat[format]
	div := math.Pow(float64(1024), float64(pow))
	return float64(size) / div
}

func print(result *Result, header string, format formatLine) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, header)

	for _, object := range result.Objects {
		fmt.Fprintln(w, format(*object))
	}
	fmt.Fprintln(w)

	w.Flush()

}
