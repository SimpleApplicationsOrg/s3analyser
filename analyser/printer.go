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

func (sat *sat) Print(result *Result) {

	h := headerFactory().bucket(sat.size)
	formatFunction := sat.formatBucket

	if sat.byRegion {
		h = h.region(sat.size)
		formatFunction = sat.formatRegion
	}

	if sat.withStorage {
		h = h.withStorage()
	}

	print(result, h.string(), formatFunction)

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

type header string

func headerFactory() header {
	return header("")
}

func (h header) string() string {
	return string(h)
}

func (h header) bucket(size string) header {
	bucketHeader := fmt.Sprintf("Name\tRegion\tCount\tTotal (%s)\tCreation\tLast Modified", size)
	return header(bucketHeader)
}

func (h header) region(size string) header {
	regionHeader := fmt.Sprintf("Region\tCount\tTotal (%s)\tCreation\tLast Modified", size)
	return header(regionHeader)
}

func (h header) withStorage() header {
	return header(h.string() + "\tStorage")
}

func (sat *sat) formatRegion(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%d\t%.f\t%s\t%s\t%s",
		*data.Region, *data.Count, sizeCalcul(*data.Size, sat.size), data.CreationDate.String(),
		data.LastModified.String(), *data.StorageClass)
}

func (sat *sat) formatBucket(data model.ObjectData) string {
	return fmt.Sprintf("%s\t%s\t%d\t%.f\t%s\t%s\t%s",
		*data.Bucket, *data.Region, *data.Count, sizeCalcul(*data.Size, sat.size), data.CreationDate.String(),
		data.LastModified.String(), *data.StorageClass)
}

func sizeCalcul(size int64, format string) float64 {
	pow := sizeFormat[format]
	div := math.Pow(float64(1024), float64(pow))
	return float64(size) / div
}
