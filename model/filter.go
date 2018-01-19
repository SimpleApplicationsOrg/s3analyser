package model

import (
	"errors"
	"strings"
)

type FilterMap map[string]string

func (i *FilterMap) String() string {
	return "filter"
}

func (i *FilterMap) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("filter flag already set")
	}
	aux := make(map[string]string)
	for _, dt := range strings.Split(value, ",") {
		parts := strings.Split(dt, "/")
		key := parts[0]

		aux[key] = strings.Join(parts[1:],"/")
	}
	*i = FilterMap(aux)
	return nil
}
