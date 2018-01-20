package analyser

import "fmt"

type header string

func bucketHeaderBuilder(size string) header {
	bucketHeader := "Bucket\t" + regionHeaderBuilder(size)
	return header(bucketHeader)
}

func regionHeaderBuilder(size string) header {
	regionHeader := fmt.Sprintf("Region\tCount\tTotal (%s)\tCreation\tLast Modified", size)
	return header(regionHeader)
}

func (h header) withStorage() header {
	return header(h.build() + "\tStorage")
}

func (h header) build() string {
	return string(h)
}
